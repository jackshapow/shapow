package model

import (
	"fmt"
	"github.com/gocraft/dbr"
	// "golang.org/x/crypto/bcrypt"
	// "reflect"
)

// User is the user datasource skeleton
type Song struct {
	PrimaryID
	Title dbr.NullString `json:"title"`
	Path  dbr.NullString `json:"path"`
}

func (song *Song) FindById(db dbr.Session) bool {
	fmt.Println("derp")
	var s Song
	err := db.Select("id", "title", "path").From("users").Where("id = ?", song.PrimaryID).LoadStruct(&s)
	if err != nil {
		return false
	}

	return false
}
