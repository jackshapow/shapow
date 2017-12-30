package model

import (
	"github.com/dgraph-io/badger"
	// "log"
	"fmt"
	"time"
)

var DB *badger.DB

func init() {

}

func ResetDB(db badger.DB, node_settings *Node) error {
	fmt.Println("Resetting database...")

	db.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions

		it := txn.NewIterator(opts)
		prefix := []byte("f:")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			txn.Delete(k)

		}

		nit := txn.NewIterator(opts)
		prefix = []byte("p:")
		for nit.Seek(prefix); nit.ValidForPrefix(prefix); nit.Next() {
			item := nit.Item()
			k := item.Key()
			txn.Delete(k)

		}

		zit := txn.NewIterator(opts)
		prefix = []byte("fi:")
		for zit.Seek(prefix); zit.ValidForPrefix(prefix); zit.Next() {
			item := zit.Item()
			k := item.Key()
			txn.Delete(k)
		}

		uit := txn.NewIterator(opts)
		prefix = []byte("u:")
		for uit.Seek(prefix); uit.ValidForPrefix(prefix); uit.Next() {
			item := uit.Item()
			k := item.Key()
			txn.Delete(k)
		}

		zuit := txn.NewIterator(opts)
		prefix = []byte("ue:")
		for zuit.Seek(prefix); zuit.ValidForPrefix(prefix); zuit.Next() {
			item := zuit.Item()
			k := item.Key()
			txn.Delete(k)
		}

		return nil
	})

	u := User{Name: "Demo User", Email: "demo@demo.com", Password: "demo"}
	u.Create(db)
	u2 := User{Name: "Jerry Smith", Email: "jerry@smith.com", Password: "demo"}
	u2.Create(db)
	u3 := User{Name: "Barry Jones", Email: "barry@jones.com", Password: "demo"}
	u3.Create(db)
	// if err != nil {
	// 	fmt.Println("FUCK FUCK FUCK", err)
	// }

	// Load existing media
	RescanFolder(db, node_settings)

	return nil
}

// PrimaryID is an ID field common to most models
type PrimaryID struct {
	ID uint `json:"id"`
}

// CommonDates is a set of common dates fields
type CommonDates struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt *time.Time `json:"deleted_at"`
}
