package model

import (
	"fmt"
	// "golang.org/x/crypto/bcrypt"
	"errors"
	//"github.com/davecgh/go-spew/spew"
	"github.com/dgraph-io/badger"
	"github.com/golang/protobuf/proto"
	//"reflect"
	//"strings"
)

func AllPlaylists(db badger.DB, playlistType PlaylistType) ([]Playlist, error) {
	playlistSlice := make([]Playlist, 0)

	db.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		prefix := []byte("p:" + string(playlistType))
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			v, err := item.Value()

			//fmt.Println(".....................", v)
			newPlaylist := &Playlist{}
			err = proto.Unmarshal(v, newPlaylist)
			if err != nil {
				fmt.Println("unmarshaling error: ", err)
			}

			// cleanup: set values for view
			// ]newFile.AlbumId = "2"
			// newFile.ArtistId = "110320904029614"
			// newFile.Title = newFile.Meta["Title"]

			if playlistType == newPlaylist.Type {
				playlistSlice = append(playlistSlice, *newPlaylist)
			}
			//k := item.Key()
			//fmt.Println("PLAYLIST ID ", string(k))
			//spew.Dump(newPlaylist)
			// err = txn.Delete(k)
			if err != nil {
				fmt.Println("     ERROR:", err)
			}
			//fmt.Println("Delete playlist: ", k)

			if err != nil {
				return err
			}
			// fmt.Printf("key=%s, value=%s\n", k, v)
			// fmt.Printf("...", newFile.GetId(), "...")

		}
		return nil
	})

	// file_slices := make([]File, 0, 2)
	// file_slices = append(file_slices, File{Id: 12345, Path: "my song gere"})
	// file_slices = append(file_slices, File{Id: 29033, Path: "my other gere"})
	// return file_slices, nil
	return playlistSlice, nil
}

func (playlist *Playlist) Create(db badger.DB) error {
	if playlist.Cover == "" {
		playlist.Cover = "public/img/covers/unknown.jpg"
	}

	data, err := proto.Marshal(playlist)

	if err != nil {
		return err
	}

	err = db.Update(func(txn *badger.Txn) error {
		// check for existing account
		_, err := txn.Get([]byte("p:" + string(playlist.Type) + playlist.Id))
		if err == nil {
			//fmt.Println(err)
			//spew.Dump(reflect.TypeOf(p))
			return errors.New("     Playlist already exists.")
		}

		ee := txn.Set([]byte("p:"+string(playlist.Type)+playlist.Id), data)

		if ee != nil {
			fmt.Println("     ERROR: Error storing playlist value")
			fmt.Println(err)
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println("     ERROR: ", err)
		//fmt.Println(err)
		return err
	}

	return nil
}
