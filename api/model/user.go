package model

import (
	"fmt"
	"github.com/gocraft/dbr"
	// "golang.org/x/crypto/bcrypt"
	// "reflect"
)

// User is the user datasource skeleton
type User struct {
	PrimaryID
	Name     dbr.NullString `json:"name"`
	Email    dbr.NullString `json:"email"`
	Password dbr.NullString `json:"password"`
}

func (user *User) Create(db dbr.Session) bool {
	fmt.Println(user)
	// user.Name = dbr.NewNullString("zzzzzzzzzz")
	// user.Email = dbr.NewNullString("zzzz@zzzzz.com")
	// user.Password = dbr.NewNullString("zzzpass")

	// u := User{Name: dbr.NewNullString("Bobby Jenkins"), Email: dbr.NewNullString("bjenkindds@gmail.comz"), Password: dbr.NewNullString("passwordhere")}

	_, err := db.InsertInto("users").Columns("name", "email", "password").Record(user).Exec()

	if err != nil {
		return false
	}

	// //_, err = db.InsertInto("users").Columns("name", "email", "password").Record(&u).Exec()
	// fmt.Println("----------------------")
	// fmt.Println(user)
	// //fmt.Println(u)
	// fmt.Println("----------------------")
	// fmt.Println(a)
	// fmt.Println(err)
	// //a.Build(dbr.InsertBuilder, buf)
	// fmt.Println(reflect.TypeOf(a))

	// fmt.Println("----------------------")
	// //b := a.InsertStmt

	// if err != nil {
	// 	fmt.Println(err)
	// 	return false
	// }

	return true
}

func (user *User) Authenticate(db dbr.Session) bool {
	var u User
	err := db.Select("id", "email", "password").From("users").Where("email = ?", user.Email).LoadStruct(&u)
	if err != nil {
		return false
	}

	//.cleartext password derp
	if user.Password == u.Password {
		return true
	} else {
		return false
	}

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
