package model

import (
	"fmt"
	// "golang.org/x/crypto/bcrypt"
	"crypto/rand"
	"errors"
	"github.com/btcsuite/btcutil/base58"
	"github.com/davecgh/go-spew/spew"
	"github.com/dgraph-io/badger"
	"github.com/golang/protobuf/proto"
	"golang.org/x/crypto/ed25519"
	//	"reflect"
	"strings"
)

func (user *User) SetKeypair() bool {
	if user.Id == "" {
		pubkey, privkey, _ := ed25519.GenerateKey(rand.Reader)
		user.Id = base58.CheckEncode(pubkey, 0)
		user.PubKey = pubkey
		user.PrivKey = privkey
	}
	return true
}

func (user *User) Create(db badger.DB) bool {

	user.SetKeypair()

	data, err := proto.Marshal(user)

	if err != nil {
		return false
	}

	err = db.Update(func(txn *badger.Txn) error {
		// check for existing account
		_, err := txn.Get([]byte("ue:" + strings.ToLower(user.Email)))
		if err == nil {
			return errors.New("That email already exists.")
		}

		err = txn.Set([]byte("u:"+user.Id), data)
		err = txn.Set([]byte("ue:"+strings.ToLower(user.Email)), []byte(user.Id))
		return err
	})

	if err != nil {
		fmt.Println("-Error saving badger-------")
		fmt.Println(err)
		return false
	}

	fmt.Println("---------------------------")
	spew.Dump(user)
	fmt.Println("---------------------------")

	return true
}

func (user *User) Authenticate(db badger.DB) bool {
	password := user.Password

	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("ue:" + strings.ToLower(user.Email)))
		if err != nil {
			return errors.New("That email address doesn't exist.")
		}

		id, err := item.Value()
		val, err := txn.Get(append([]byte("u:"), id...))

		if err != nil {
			return errors.New("Cannot load account.")
		}

		data, err := val.Value()

		err = proto.Unmarshal(data, user)
		if err != nil {
			return errors.New("Cannot load account.")
		}

		return nil
	})

	if err != nil {
		return false //,errors.New("Cannot load account.")
	}

	if password == user.Password {
		return true
	} else {
		return false
	}

	// return false
	return false
}

// func (user *User) Save(db *pg.DB) {
// 	err := db.Insert(&user)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func generateEmailVerificationToken(db orm.DB, length int) string {
// 	var token string
// 	for {
// 		token = util.GenerateRandomStrig(60, true, true, true, false)

// 		count, _ := db.Model(User{}).
// 			Where("verification_token = ?", token).
// 			Count()

// 		if count < 1 {
// 			break
// 		}
// 	}

// 	return token
// }

// func generatePasswordHash(password string) string {
// 	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// 	return string(hash)
// }
