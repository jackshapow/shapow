package model

import (
	"fmt"
	//"github.com/gocraft/dbr"
	// "golang.org/x/crypto/bcrypt"
	// "reflect"
	"github.com/dgraph-io/badger"
)

// User is the user datasource skeleton
type Song struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

func (song *Song) FindById(db badger.DB) bool {
	fmt.Println("Looking up song " + song.Id)
	//var s Song
	// err := db.Select("id", "title", "path").From("songs").Where("id = ?", song.Id).LoadStruct(&song)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return false
	// }

	return true
}

// func (user *User) Authenticate(db dbr.Session) bool {
//   var u User
//   err := db.Select("id", "email", "password").From("users").Where("email = ?", user.Email).LoadStruct(&u)
//   if err != nil {
//     return false
//   }

//   //.cleartext password derp
//   if user.Password == u.Password {
//     return true
//   } else {
//     return false
//   }

//   return false
// }
