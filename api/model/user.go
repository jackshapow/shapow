package model

import (
	"fmt"
	// "golang.org/x/crypto/bcrypt"
	"crypto/rand"
	"errors"
	"github.com/btcsuite/btcutil/base58"
	//"github.com/davecgh/go-spew/spew"
	"github.com/dgraph-io/badger"
	"github.com/golang/protobuf/proto"
	"golang.org/x/crypto/ed25519"
	//	"reflect"
	"github.com/imdario/mergo"
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

func (user *User) FindById(db badger.DB) error {
	fmt.Println("Looking up user " + user.Id)

	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("u:" + user.Id))

		if err != nil {
			return errors.New("That user id doesnt exist")
		}

		data, err := item.Value()

		err = proto.Unmarshal(data, user)
		if err != nil {
			return errors.New("Cannot unpack protobuf.")
		}

		return nil
	})

	if err != nil {
		fmt.Println("Can't get user", err)
		return err
	}

	return nil
}

func (user *User) Create(db badger.DB) error {

	user.SetKeypair()

	data, err := proto.Marshal(user)

	if err != nil {
		return err
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
		return err
	}

	// fmt.Println("---------------------------")
	// spew.Dump(user)
	// fmt.Println("---------------------------")

	return nil
}

func AllUsers(db badger.DB) ([]User, error) {
	userSlice := make([]User, 0)

	db.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		prefix := []byte("u:")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			v, err := item.Value()

			newUser := &User{}
			err = proto.Unmarshal(v, newUser)
			if err != nil {
				fmt.Println("unmarshaling error: ", err)
			}

			userSlice = append(userSlice, *newUser)

			if err != nil {
				return err
			}
			// fmt.Printf("key=%s, value=%s\n", k, v)

		}
		return nil
	})

	return userSlice, nil
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

func (user *User) Update(db badger.DB) error {
	var newUser = user

	err := db.Update(func(txn *badger.Txn) error {
		q := []byte("u:" + newUser.Id)

		item, err := txn.Get(q)

		if err != nil {
			fmt.Println("User not found")
			return errors.New("     User not found")
		}

		data, err := item.Value()

		if err != nil {
			return errors.New("     Could not decode user")
		}

		user := &User{}

		if err := proto.Unmarshal(data, user); err != nil {
			return errors.New("     Unmarshalling error")
		}

		if err := mergo.Merge(newUser, user); err != nil {
			fmt.Println("Error merging values", err)
		}

		data, err = proto.Marshal(newUser)

		if err != nil {
			fmt.Println("Error marshalling proto")
			return err
		}

		err = txn.Set(q, data)

		if err != nil {
			fmt.Println("Error setting new value")
			return err
		}

		err = txn.Set([]byte("ue:"+strings.ToLower(newUser.Email)), []byte(newUser.Id))

		return nil
	})

	if err != nil {
		fmt.Println("Error saving to badger")
		return err
	}

	return nil

}

func UserDelete(db badger.DB, id string) error {

	err := db.Update(func(txn *badger.Txn) error {
		// grab playlist
		q := []byte("u:" + id)

		err := txn.Delete(q)

		if err != nil {
			fmt.Println("User not found")
			return errors.New("     User not found")
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error deleting in badger")
	}

	return nil
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
