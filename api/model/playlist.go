package model

import (
	"fmt"
	// "golang.org/x/crypto/bcrypt"
	"errors"
	//"github.com/davecgh/go-spew/spew"
	"github.com/dgraph-io/badger"
	"github.com/golang/protobuf/proto"
	"reflect"
	//"strings"
	"encoding/json"
	"math/rand"
	"path/filepath"
	//"strconv"
	"github.com/imdario/mergo"
)

func (p Playlist) CoverUrl(artwork_path string) string {
	if p.Cover == "" {
		return "public/img/covers/unknown.jpg"
	} else {
		return filepath.Join(artwork_path, p.Cover)
	}
}

func (p Playlist) MarshalJSON() ([]byte, error) {
	type JsonPlaylist Playlist

	aux := struct {
		JsonPlaylist
		ParentId string `json:"artist_id"`
		//Files    []*File `json:"songs"`
		Songs []string `json:"songs"`
	}{
		JsonPlaylist: JsonPlaylist(p),
		ParentId:     p.ParentId,
	}

	if len(aux.Songs) == 0 { // if we dont do this then vue dies when it gets songs:null and not songs:[]
		aux.Songs = []string{}
	}

	for _, file := range aux.Files {
		aux.Songs = append(aux.Songs, file.Id)
	}

	return json.Marshal(aux)
}

// Keep this around for later use?
// func (u *MyUser) UnmarshalJSON(data []byte) error {
// 	type Alias MyUser
// 	aux := &struct {
// 		LastSeen int64 `json:"lastSeen"`
// 		*Alias
// 	}{
// 		Alias: (*Alias)(u),
// 	}
// 	if err := json.Unmarshal(data, &aux); err != nil {
// 		return err
// 	}
// 	u.LastSeen = time.Unix(aux.LastSeen, 0)
// 	return nil
// }

// func (p Playlist) MarshalJSON() ([]byte, error) {
// 	type JsonPlaylist Playlist
// 	aux := struct {
// 		//Songs int64 `json:"lastSeen"`
// 		Cat string `json:"cats"`
// 		*JsonPlaylist
// 	}{
// 		JsonPlaylist: (*JsonPlaylist)(p),
// 	}
// 	if err := json.Marshal(&aux); err != nil {
// 		return err
// 	}
// 	p.Cat = "herpderp"
// 	//u.LastSeen = time.Unix(aux.LastSeen, 0)
// 	return nil
// }

func PlaylistDelete(db badger.DB, id string) error {

	err := db.Update(func(txn *badger.Txn) error {
		// grab playlist
		q := []byte("p:" + string(PlaylistType_value["User"]) + id)

		err := txn.Delete(q)

		if err != nil {
			fmt.Println("Playlist not found")
			return errors.New("     Playlist not found")
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error deleting in badger")
	}

	return nil
}

func UpdatePlaylistSongs(db badger.DB, id string, song_ids []string) error {
	// fmt.Println("id:", id)
	// fmt.Println("song_ids:", song_ids)
	var playlistFiles = make([]*File, len(song_ids))

	err := db.Update(func(txn *badger.Txn) error {
		// grab playlist
		q := []byte("p:" + string(PlaylistType_value["User"]) + id)

		item, err := txn.Get(q)
		// fmt.Println("ref", reflect.TypeOf(PlaylistType_value["User"]))
		// fmt.Println("1>", string(PlaylistType_User), ":", id)
		if err != nil {
			//fmt.Println(err)
			//spew.Dump(reflect.TypeOf(p))
			fmt.Println("playlist not found")
			return errors.New("     Playlist not found")
		}

		data, err := item.Value()

		if err != nil {
			fmt.Println("couldnt decode value")
		}

		playlist := &Playlist{}
		//		err = proto.Unmarshal(data, playlist)

		if err := proto.Unmarshal(data, playlist); err != nil {
			fmt.Println("fuckzzzz")
			return errors.New("     Unmarshalling error")
		}

		for i, song_id := range song_ids {
			playlistFiles[i] = &File{Id: song_id, Track: uint32(i)} //i
		}
		//playlist.Files
		playlist.Files = playlistFiles
		fmt.Println("DO IT SON", playlist)
		fmt.Println("DO IT SON", reflect.TypeOf(playlist))
		data, err = proto.Marshal(playlist)

		if err != nil {
			fmt.Println("Error marshalling proto")
		}

		err = txn.Set(q, data)

		if err != nil {
			fmt.Println("Error setting new value")
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error saving to badger")
	}

	return nil
}

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
			//newFile.Title = newFile.Meta["Title"]
			//newPlaylist.Cover = filepath.Join(node.ArtworkPath(), newPlaylist.Cover)

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

func (playlist *Playlist) Update(db badger.DB) error {
	var newPlaylist = playlist

	err := db.Update(func(txn *badger.Txn) error {
		// grab playlist
		q := []byte("p:" + string(PlaylistType_value["User"]) + newPlaylist.Id)

		item, err := txn.Get(q)

		if err != nil {
			fmt.Println("Playlist not found")
			return errors.New("     Playlist not found")
		}

		data, err := item.Value()

		if err != nil {
			return errors.New("     Could not decode playlist")
		}

		playlist := &Playlist{}

		if err := proto.Unmarshal(data, playlist); err != nil {
			return errors.New("     Unmarshalling error")
		}

		if err := mergo.Merge(newPlaylist, playlist); err != nil {
			fmt.Println("Error merging values", err)
		}

		data, err = proto.Marshal(newPlaylist)

		if err != nil {
			fmt.Println("Error marshalling proto")
			return err
		}

		err = txn.Set(q, data)

		if err != nil {
			fmt.Println("Error setting new value")
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error saving to badger")
		return err
	}

	return nil

}

func (playlist *Playlist) Create(db badger.DB) error {

	if playlist.Cover == "" {
		playlist.Cover = "public/img/covers/unknown.jpg"
	} else {
		playlist.Cover = filepath.Join("/artwork/", playlist.Cover)
	}

	if playlist.Id == "" {
		//playlist.Id = "11032090402961465666"
		playlist.Id = fmt.Sprint(rand.Uint32())
		fmt.Println("IDDDDDDD:", playlist.Id)
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
