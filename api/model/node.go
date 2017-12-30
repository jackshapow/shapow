package model

import (
	"fmt"
	// "golang.org/x/crypto/bcrypt"
	"errors"
	"github.com/dgraph-io/badger"
	"github.com/golang/protobuf/proto"
	//"golang.org/x/crypto/ed25519"
	//  "reflect"
	//"strings"
	"path/filepath"
)

func (node *Node) MediaPath() string {
	return filepath.Join(node.RootPath, "Media")
}

func (node *Node) ArtworkPath() string {
	return filepath.Join(node.RootPath, "Artwork")
}

func (node *Node) Initialize(db badger.DB) error {
	// Initialize a node

	err := db.Update(func(txn *badger.Txn) error {
		item, _ := txn.Get([]byte("node:settings"))
		if item != nil {
			// Load node settings
			data, err := item.Value()

			err = proto.Unmarshal(data, node)
			if err != nil {
				return errors.New("Cannot load node settings.")
			}

			fmt.Println("Loading existing node...", node)
		} else if item == nil {
			// Initialize new node
			// node.MediaPath = "mediaz"

			fmt.Println("Setting up new node...", node)
			data, err := proto.Marshal(node)

			if err != nil {
				return err
			}

			fmt.Println("ERRRRRRRRR", err)
			err = txn.Set([]byte("node:settings"), data)
		}

		return nil
	})

	if err != nil {
		fmt.Println("-Error saving badger-------")
		fmt.Println(err)
		return err
	}

	return nil
}
