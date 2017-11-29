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

func ResetDB(db badger.DB) error {

	db.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		prefix := []byte("f:")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			fmt.Println("DELETE F ", k)
			txn.Delete(k)

		}

		nit := txn.NewIterator(opts)
		prefix = []byte("p:")
		for nit.Seek(prefix); nit.ValidForPrefix(prefix); nit.Next() {
			item := nit.Item()
			k := item.Key()
			fmt.Println("DELETE P ", k)
			txn.Delete(k)

		}

		return nil
	})

	// Load existing media
	RescanFolder(db)

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
