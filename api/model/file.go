package model

import (
	"encoding/binary"
	//"encoding/json"
	"fmt"
	// "golang.org/x/crypto/bcrypt"
	// "crypto/rand"
	// "errors"
	// "github.com/btcsuite/btcutil/base58"
	//"github.com/davecgh/go-spew/spew"
	"github.com/dgraph-io/badger"
	"github.com/golang/protobuf/proto"
	// "golang.org/x/crypto/ed25519"
	"reflect"
	// "strings"
	"bytes"
	"errors"
	"github.com/OneOfOne/xxhash"
	"github.com/dhowden/tag"
	//"github.com/hajimehoshi/go-mp3"
	"gopkg.in/h2non/filetype.v1"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Interaction struct {
	SongId    string `json:"song_id"`
	Liked     bool   `json:"liked"`
	PlayCount uint32 `json:"play_count"`
}

func RescanFolder(db badger.DB) error {

	searchDir := "./media"

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})

	if e != nil {
		panic(e)
	}

	for _, file := range fileList {
		fmt.Println("RESCAN: ", file)
		newFile := File{Path: file}
		newFile.Import(db)
	}

	return nil

}

func AllInteractions(db badger.DB, files []File) ([]Interaction, error) {
	var interactions []Interaction

	for _, file := range files {
		interaction := Interaction{file.Id, false, 0}

		err := db.View(func(txn *badger.Txn) error {
			// PLAY COUNT
			item, err := txn.Get([]byte("fi:play:" + interaction.SongId))
			var iBytes = make([]byte, 4)

			if err == nil {
				iBytes, err = item.Value()
				binary.LittleEndian.PutUint32(iBytes, interaction.PlayCount)
			}

			// LIKE
			_, err = txn.Get([]byte("fi:like:" + interaction.SongId))
			if err != nil {
				interaction.Liked = false
			} else {
				interaction.Liked = true
			}

			return nil
		})

		//fmt.Println("::::", interaction)

		if err == nil {
			interactions = append(interactions, interaction)
		}

	}

	return interactions, nil
}

func AllFiles(db badger.DB) ([]File, error) {
	fileSlice := make([]File, 0)

	db.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		prefix := []byte("f:")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			v, err := item.Value()

			newFile := &File{}
			err = proto.Unmarshal(v, newFile)
			if err != nil {
				fmt.Println("unmarshaling error: ", err)
			}

			//fmt.Println(newFile)
			// Only show audio files
			if strings.HasPrefix(newFile.Mime, "audio/") {
				fileSlice = append(fileSlice, *newFile)
			}

			if err != nil {
				return err
			}
			// fmt.Printf("key=%s, value=%s\n", k, v)

		}
		return nil
	})

	return fileSlice, nil
}

func (file *File) FindById(db badger.DB) error {
	fmt.Println("Looking up file " + file.Id)
	//fmt.Println("Looking up file " + strconv.FormatUint(file.Id, 10))

	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("f:" + file.Id))

		//val, err := txn.Get(append([]byte("f:"), file.Id...))

		if err != nil {
			return errors.New("That id doesnt exist")
		}

		data, err := item.Value()

		err = proto.Unmarshal(data, file)
		if err != nil {
			return errors.New("Cannot unpack protobuf.")
		}

		// fmt.Println("XXXXXXXXXX")
		// spew.Dump(file)
		// fmt.Println("XXXXXXXXXX")

		return nil
	})

	//var s Song
	// err := db.Select("id", "title", "path").From("songs").Where("id = ?", song.Id).LoadStruct(&song)
	if err != nil {
		fmt.Println("ohshit")
		fmt.Println(err)
		return err
	}

	return nil
}

func GenerateId(data []byte) string {
	h := xxhash.New64()
	r := bytes.NewReader(data)
	io.Copy(h, r)
	return strconv.FormatUint(h.Sum64(), 10)
}

func (file *File) SetId() bool {
	if file.Id == "" {
		h := xxhash.New64()
		r, err := os.Open(file.Path)
		if err != nil {
			fmt.Println("     Cant open file")
		}
		defer r.Close()
		//r := strings.NewReader(F)
		io.Copy(h, r)
		file.Id = strconv.FormatUint(h.Sum64(), 10)
		//fmt.Println("xxhash.Backend:", h.Sum(b)
		fmt.Println("     XXHASH:", file.Id)
	}
	return true
}

func (file *File) SetName() bool {
	if file.Title == "" {
		file.Title = file.Path

		if file.Meta["Title"] != "" {
			file.Title = file.Meta["Title"]
		}

	}
	return true
}

func (file *File) SetAlbumArtist(db badger.DB) bool {
	err := db.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false

		// Create artist
		artist_hash_name := file.Meta["AlbumArtist"]
		if len(artist_hash_name) == 0 {
			artist_hash_name = "Unknown"
		}

		fmt.Println("     SET ARTIST ", artist_hash_name)

		ah := xxhash.New64()
		io.Copy(ah, strings.NewReader(artist_hash_name))
		file.ArtistId = strconv.FormatUint(ah.Sum64(), 10)

		artist := Playlist{Id: file.ArtistId, Name: artist_hash_name, Type: PlaylistType_Artist}
		artist.Create(db)
		fmt.Println("     CREATE PLAYLIST: ", artist)

		// Create album
		var album_name string
		album_hash_name := artist_hash_name + file.Meta["Album"]
		if len(file.Meta["Album"]) == 0 {
			album_name = "Unknown"
		} else {
			album_name = file.Meta["Album"]
		}
		fmt.Println("     SET ALBUM ", album_name)

		aah := xxhash.New64()
		io.Copy(aah, strings.NewReader(album_hash_name))
		file.AlbumId = strconv.FormatUint(ah.Sum64(), 10)

		album := Playlist{Id: file.AlbumId, Name: album_name, ParentId: artist.Id, Type: PlaylistType_Album, ImageId: file.Meta["image_id"], Cover: file.Meta["cover"]}
		err := album.Create(db)
		if err != nil {
			fmt.Println("     COULD NOT CREATE PLAYLIST: ", album_name)
			return nil
		}
		fmt.Println("     CREATE PLAYLIST: ", album_name)

		return nil
	})

	if err != nil {
		fmt.Println("fuck fuck fuck")
	}
	return true
}

func (file *File) SetMime() error {
	buf, _ := ioutil.ReadFile(file.Path)

	kind, unknown := filetype.Match(buf)
	if unknown != nil {
		return unknown
	}

	if kind.Extension == "unknown" {
		return errors.New("Unknown File Extension")
	}

	//fmt.Println("Nice", ":", kind.Extension, ":")

	file.Mime = kind.MIME.Value
	return nil
}

func (file *File) ParseID3(db badger.DB) error {
	f, err := os.Open(file.Path)
	m, err := tag.ReadFrom(f)

	if err != nil {
		fmt.Println("     ParseID3 Error:", err)
		return nil
	}

	file.Meta = make(map[string]string)
	fmt.Println("     Parse Meta Data")
	//file.Meta["Format"] = m.Format()
	file.Meta["FileType"] = string(m.FileType())
	file.Meta["Title"] = m.Title() // The title of the track (see Metadata interface for more details).
	file.Meta["Album"] = m.Album()
	file.Meta["Artist"] = m.Artist()
	file.Meta["AlbumArtist"] = m.AlbumArtist()
	file.Meta["Composer"] = m.Composer()
	file.Meta["Genre"] = m.Genre()
	file.Meta["Year"] = string(m.Year())
	//file.Meta["Disc"] = m.Disc()
	k, v := m.Disc()
	//fmt.Println(reflect.TypeOf(k))
	//fmt.Println(reflect.TypeOf(v))
	//fmt.Println(m.Disc())
	//file.Meta["Picture"] = m.Picture()
	fmt.Println("     DISC: ", k, ":", v)
	fmt.Println("     PHOTO: ", m.Picture())
	fmt.Println("     ", reflect.TypeOf(m.Picture()))
	if m.Picture() != nil {
		id := GenerateId(m.Picture().Data)
		picture_path := filepath.Join(".", "media", "artwork", id+"."+strings.ToLower(m.Picture().Ext))
		os.MkdirAll(filepath.Dir(picture_path), os.ModePerm)

		err := ioutil.WriteFile(picture_path, m.Picture().Data, 0644)

		if err != nil {
			fmt.Println("     Error writing cover image")
		}
		//spew.Dump(m.Picture())
		newFile := File{Path: picture_path, Title: m.Picture().Type}
		newFile.Import(db)

		file.Meta["image_id"] = newFile.Id
		file.Meta["cover"] = strings.Replace(newFile.Path, "media/", "api/media/", 1)
		//public/img/covers/unknown-album.png
		//file.
	}
	file.Meta["Lyrics"] = m.Lyrics()
	track, _ := m.Track()
	file.Track = uint32(track)

	return nil
}

func (file *File) SetDuration() error {
	var (
		//cmdOut []byte
		err error
	)
	cmdName := "./binaries/osx/ffmpeg"
	cmdArgs := []string{"-i", file.Path} // Windows should be 2>NUL

	cmd := exec.Command(cmdName, cmdArgs...)
	output, err := cmd.CombinedOutput()

	// Have to ignore the return code because not specifying -output in FFMPEG throws error code
	if err != nil {
		//fmt.Println("FFMPEG cannot parse duration")
		//fmt.Println(fmt.Sprint(err) + ": " + string(output))
	}

	re1, err := regexp.Compile(`Duration: (.*?),\W`) // Prepare our regex
	result_slice := re1.FindAllStringSubmatch(string(output), -1)[0][1]

	hours, _ := strconv.Atoi(result_slice[0:2])
	minutes, _ := strconv.Atoi(result_slice[3:5])
	seconds, _ := strconv.Atoi(result_slice[6:8])
	subsec, _ := strconv.Atoi(result_slice[9:11])

	total_seconds := (hours * 60 * 60) + (minutes * 60) + seconds
	if subsec > 50 {
		total_seconds = total_seconds + 1
	}

	if total_seconds > 0 {
		file.Length = uint32(total_seconds)
	}

	return nil
}

func SetInteraction(db badger.DB, id string, kind string) (Interaction, error) {
	interaction := Interaction{id, false, 0}
	err := db.Update(func(txn *badger.Txn) error {
		// PLAY COUNT
		item, err := txn.Get([]byte("fi:play:" + id))
		var iBytes = make([]byte, 4)

		if err != nil { // Set default
			interaction.PlayCount = 1
			binary.LittleEndian.PutUint32(iBytes, interaction.PlayCount)
		} else { // Increment
			iBytes, err = item.Value()
			if err != nil {
			}
			interaction.PlayCount = binary.LittleEndian.Uint32(iBytes) + 1
			binary.LittleEndian.PutUint32(iBytes, interaction.PlayCount)
		}

		if kind == "play" {
			err = txn.Set([]byte("fi:play:"+id), iBytes)

			if err != nil {
				fmt.Println("Saving play failed.")
				return err
			} else {
				fmt.Println("Saving play success!")
			}
		}

		// LIKE
		_, err = txn.Get([]byte("fi:like:" + id))
		if err != nil {
			interaction.Liked = false
		} else {
			interaction.Liked = true
		}

		if kind == "like" {
			if interaction.Liked == false {
				// Like it
				zbytes := make([]byte, 0) //{0}
				err = txn.Set([]byte("fi:like:"+id), zbytes)
				if err == nil {
					interaction.Liked = true
				}
			} else {
				// Unlike it
				err = txn.Delete([]byte("fi:like:" + id))
				if err == nil {
					interaction.Liked = false
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error saving badger transaction", err)
	}

	//val, err := json.Marshal(interaction)

	//fmt.Println("NEW VALUE:", string(val))

	return interaction, err
}

func (file *File) Import(db badger.DB) error {
	// grab meta data
	// import into list / search index
	err := file.SetMime()
	if err != nil {
		fmt.Println("     File type not recognized.")
		return err
	}
	file.SetId()
	file.ParseID3(db)
	file.SetName()
	file.SetAlbumArtist(db)
	if strings.Contains(file.Path, ".mp3") {
		file.SetDuration()
	}

	fmt.Println("      ArtistID", file.ArtistId)
	fmt.Println("       AlbumID", file.AlbumId)

	data, err := proto.Marshal(file)

	if err != nil {
		return err
	}

	err = db.Update(func(txn *badger.Txn) error {
		err = txn.Set([]byte("f:"+file.Id), data)
		//strings.Join([]string{"key:", "value", ", key2:", strconv.Itoa(100)}, "")

		// err = txn.Set([]byte("ue:"+strings.ToLower(user.Email)), []byte(user.Id))
		return err
	})

	if err != nil {
		fmt.Println("ERROR: ", err)
		return err
	}

	// fmt.Println("---------------------------")
	// spew.Dump(file)
	// fmt.Println("---------------------------")

	return nil
}
