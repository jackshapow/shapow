package controller

import (
	//"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	// "github.com/gocraft/dbr"
	"github.com/jackshapow/shapow/api/model"
	"github.com/labstack/echo"
	"net/http"
	//"reflect"
	// "github.com/golang/protobuf/proto"
	// "github.com/jackshapow/shapow/api/model/proto"
	//"io/ioutil"
	"encoding/json"
	"time"
)

func (h *Handler) UserCreate(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	fmt.Println("Creating user...")

	if u.Create(*h.DB) {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = u.Name
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, map[string]string{
			"token": t,
		})
	}

	return echo.NewHTTPError(http.StatusBadRequest, "Sorry couldn't create account.")
}

func (h *Handler) UserLogout(c echo.Context) error {
	return c.String(http.StatusOK, "Bye")
}

func (h *Handler) UserLogin(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	if u.Authenticate(*h.DB) {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = u.Name
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	//	return echo.ErrUnauthorized
	return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func (h *Handler) UserDataTest(c echo.Context) error {

	res, err := model.AllFiles(*h.DB)
	if err != nil {
		fmt.Println("Error fetching all files")
	}

	r, _ := json.Marshal(res)
	//res := []string{"foo", "bar"}

	return c.String(http.StatusOK, string(r))
}

// func stubData(name string) string {
//     re
// }

func (h *Handler) UserData(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	return c.String(http.StatusOK, stubData(h, name))
}

func stubData(h *Handler, name string) string {
	all_files, err := model.AllFiles(*h.DB)
	if err != nil {
		fmt.Println("Error fetching all files")
	}
	songArray, _ := json.Marshal(all_files)

	all_albums, err := model.AllPlaylists(*h.DB, model.PlaylistType_Album)
	if err != nil {
		fmt.Println("Error fetching all albums")
	}

	albumArray, _ := json.Marshal(all_albums)

	all_artists, err := model.AllPlaylists(*h.DB, model.PlaylistType_Artist)
	if err != nil {
		fmt.Println("Error fetching all artists")
	}

	artistArray, _ := json.Marshal(all_artists)

	// fmt.Println(".........")
	// fmt.Println(reflect.TypeOf(songArray))
	// fmt.Println(".........")
	// fmt.Println(songArray)
	// fmt.Println(".........")
	// fmt.Println(string(songArray))
	// fmt.Println(".........")

	// "artists": [
	//     {
	//         "id": "11032090402961465788",
	//         "name": "STATIC Chance the Rapper Artist",
	//         "image": "http://koel.app/public/img/covers/599cb7202544d1.46641926.jpeg"
	//     }
	// ],

	// {
	//     "id": "3352024042766435742",
	//     "artist_id": "11032090402961465788",
	//     "name": "STATIC: Coloring Book Album",
	//     "cover": "http://koel.app/public/img/covers/599cb7202544d1.46641926.jpeg",
	//     "created_at": "2017-08-22 22:58:40",
	//     "is_compilation": true
	// }

	return `{
    "albums": ` + string(albumArray) + `,
    "artists": ` + string(artistArray) + `,
    "songs": ` + string(songArray) + `,
    "settings": {
        "media_path": "/home/vagrant/Code/koel/music"
    },
    "playlists": [],
    "interactions": [
        {
            "song_id": "1a0746a20d7e06393bc91d584616dd15",
            "liked": false,
            "play_count": 1
        }
    ],
    "users": [
        {
            "id": 1,
            "name": "herp derp",
            "email": "herp@derp.com",
            "is_admin": true,
            "preferences": []
        }
    ],
    "currentUser": {
        "id": 1,
        "name": "` + name + `",
        "email": "",
        "is_admin": true,
        "preferences": []
    },
    "useLastfm": false,
    "useYouTube": false,
    "useiTunes": true,
    "allowDownload": true,
    "supportsTranscoding": false,
    "cdnUrl": "http://localhost:3000/",
    "currentVersion": "v3.6.2",
    "latestVersion": "v3.6.2"
}`
}

// `[
//         {
//             "id": "1806358967a29fe64b6822fafdd72e5f",
//             "album_id": 3,
//             "artist_id": 4,
//             "title": "A L I E N S",
//             "length": 282.5,
//             "track": 3,
//             "created_at": "2017-08-22 22:58:40"
//         },
//         {
//             "id": "19bfa1969dab011d47b447e12b17f63a",
//             "album_id": 2,
//             "artist_id": 3,
//             "title": "Ya Hey",
//             "length": 312.69,
//             "track": 10,
//             "created_at": "2017-08-22 22:53:21"
//         },
//         {
//             "id": "1a0746a20d7e06393bc91d584616dd15",
//             "album_id": 3,
//             "artist_id": 4,
//             "title": "All I Can Think About Is You",
//             "length": 274.66,
//             "track": 1,
//             "created_at": "2017-08-22 22:58:40"
//         }
//     ]`
