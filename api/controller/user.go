package controller

import (
	//"encoding/json"
	// "fmt"
	"github.com/dgrijalva/jwt-go"
	// "github.com/gocraft/dbr"
	"github.com/jackshapow/shapow/api/model"
	"github.com/labstack/echo"
	"net/http"
	// "reflect"
	"time"
)

// var data string

func (h *Handler) UserData(c echo.Context) error {
	return c.String(http.StatusOK, stubData())
}

func (h *Handler) UserCreate(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	if u.Create(*h.DB) {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jonz Snow"
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

	// // json_map := make(map[string]interface{})
	// // err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	// // var v interface{}
	// // json.NewDecoder(c.Request().Body).Decode(&v)

	// fmt.Println("----------------")
	// fmt.Println(u.Email)
	// fmt.Println(u.Password)
	// // fmt.Println(string(password))
	// //fmt.Printf("%v", c.FormParams())
	// fmt.Println("----------------")

	// aa := model.User{Name: dbr.NewNullString("Bobby Jenkins"), Email: dbr.NewNullString("bjenkins@gmail.com"), Password: dbr.NewNullString("passwordhere")}
	// fmt.Println("1......")
	// fmt.Println(reflect.TypeOf(h.DB))
	// fmt.Println("2......")
	// fmt.Println(aa)
	// fmt.Println("3......")

	// _, err := h.DB.InsertInto("users").Columns("name", "email", "password").Record(&u).Exec()

	// //_, err := h.DB.InsertInto("users").Columns("name").Record(&u).Exec()
	// //fmt.Println(a)
	// fmt.Println("4......")
	// fmt.Println(err)
	// fmt.Println("5......")

	// if username == "jack@test.com" && password == "pass" {
	if u.Authenticate(*h.DB) {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jonz Snow"
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

func stubData() string {
	return `{
    "albums": [
        {
            "id": 3,
            "artist_id": 2,
            "name": "Kaleidoscope EP",
            "cover": "http://koel.app/public/img/covers/599cb7202544d1.46641926.jpeg",
            "created_at": "2017-08-22 22:58:40",
            "is_compilation": true
        },
        {
            "id": 2,
            "artist_id": 2,
            "name": "Modern Vampires Of The City",
            "cover": "http://koel.app/public/img/covers/599cb5e12f5d89.97869087.jpeg",
            "created_at": "2017-08-22 22:53:21",
            "is_compilation": true
        },
        {
            "id": 1,
            "artist_id": 1,
            "name": "Unknown Album",
            "cover": "http://koel.app/public/img/covers/unknown-album.png",
            "created_at": "2017-08-17 21:18:38",
            "is_compilation": false
        }
    ],
    "artists": [
        {
            "id": 4,
            "name": "Coldplay",
            "image": null
        },
        {
            "id": 5,
            "name": "Coldplay & Big Sean",
            "image": null
        },
        {
            "id": 6,
            "name": "Coldplay & The Chainsmokers",
            "image": null
        },
        {
            "id": 1,
            "name": "Unknown Artist",
            "image": null
        },
        {
            "id": 3,
            "name": "Vampire Weekend",
            "image": null
        },
        {
            "id": 2,
            "name": "Various Artists",
            "image": null
        }
    ],
    "songs": [
        {
            "id": "1806358967a29fe64b6822fafdd72e5f",
            "album_id": 3,
            "artist_id": 4,
            "title": "A L I E N S",
            "length": 282.5,
            "track": 3,
            "created_at": "2017-08-22 22:58:40"
        },
        {
            "id": "19bfa1969dab011d47b447e12b17f63a",
            "album_id": 2,
            "artist_id": 3,
            "title": "Ya Hey",
            "length": 312.69,
            "track": 10,
            "created_at": "2017-08-22 22:53:21"
        },
        {
            "id": "1a0746a20d7e06393bc91d584616dd15",
            "album_id": 3,
            "artist_id": 4,
            "title": "All I Can Think About Is You",
            "length": 274.66,
            "track": 1,
            "created_at": "2017-08-22 22:58:40"
        },
        {
            "id": "2d6fc57ed956048beb7e78ff8c5ddcb2",
            "album_id": 2,
            "artist_id": 3,
            "title": "Diane Young",
            "length": 160.1,
            "track": 4,
            "created_at": "2017-08-22 22:53:21"
        },
        {
            "id": "34f380415095420cdce93f7779e3de8d",
            "album_id": 2,
            "artist_id": 3,
            "title": "Everlasting Arms",
            "length": 183.35,
            "track": 7,
            "created_at": "2017-08-22 22:53:21"
        },
        {
            "id": "38e5980d86c6af478995801d93eef032",
            "album_id": 2,
            "artist_id": 3,
            "title": "Young Lion",
            "length": 105.4,
            "track": 12,
            "created_at": "2017-08-22 22:53:21"
        },
        {
            "id": "3fa0fb6dc0d9391e42b1e97cb87e5c0c",
            "album_id": 2,
            "artist_id": 3,
            "title": "Hannah Hunt",
            "length": 238,
            "track": 6,
            "created_at": "2017-08-22 22:53:21"
        },
        {
            "id": "4440a36cd128cd332e8a8c737e29dd3e",
            "album_id": 2,
            "artist_id": 3,
            "title": "Worship You",
            "length": 201.27,
            "track": 9,
            "created_at": "2017-08-22 22:53:21"
        },
        {
            "id": "4f2083b7460591e956b635376c66c372",
            "album_id": 2,
            "artist_id": 3,
            "title": "Obvious Bicycle",
            "length": 251.3,
            "track": 1,
            "created_at": "2017-08-22 22:53:21"
        },
        {
            "id": "691d25fba7dbbbefaa4c680a60b3854f",
            "album_id": 3,
            "artist_id": 5,
            "title": "Miracles (Someone Special)",
            "length": 276.97,
            "track": 2,
            "created_at": "2017-08-22 22:58:40"
        },
        {
            "id": "700b16100d01df3ea872355f72174623",
            "album_id": 2,
            "artist_id": 3,
            "title": "Finger Back",
            "length": 206,
            "track": 8,
            "created_at": "2017-08-22 22:53:21"
        },
        {
            "id": "7eb70058d522ac3293c18608419d7aa9",
            "album_id": 3,
            "artist_id": 6,
            "title": "Something Just Like This (Tokyo Remix)",
            "length": 273.79,
            "track": 4,
            "created_at": "2017-08-22 22:58:40"
        },
        {
            "id": "b3bf16db26886b4d571d5c76a1d259d7",
            "album_id": 2,
            "artist_id": 3,
            "title": "Hudson",
            "length": 254.85,
            "track": 11,
            "created_at": "2017-08-22 22:53:21"
        },
        {
            "id": "b65ae046a159321cc1cd2c254c71fbe5",
            "album_id": 3,
            "artist_id": 4,
            "title": "Hypnotised (EP Mix)",
            "length": 391.45,
            "track": 5,
            "created_at": "2017-08-22 22:58:41"
        },
        {
            "id": "d65df9004ad5303d4c8848eb732567f7",
            "album_id": 2,
            "artist_id": 3,
            "title": "Step",
            "length": 251.66,
            "track": 3,
            "created_at": "2017-08-22 22:53:21"
        },
        {
            "id": "e9f74609f248f1d6a67b840628a62e6a",
            "album_id": 2,
            "artist_id": 3,
            "title": "Don't Lie",
            "length": 213.45,
            "track": 5,
            "created_at": "2017-08-22 22:53:21"
        },
        {
            "id": "fbb13f9b59e6712e3bd5226335182c53",
            "album_id": 2,
            "artist_id": 3,
            "title": "Unbelievers",
            "length": 202.71,
            "track": 2,
            "created_at": "2017-08-22 22:53:21"
        }
    ],
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
        "name": "herp derp",
        "email": "herp@derp.com",
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
