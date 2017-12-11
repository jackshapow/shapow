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

	if err := u.Create(*h.DB); err != nil {
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
		claims["id"] = u.Id
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

func (h *Handler) UserUpdate(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		userJwt := c.Get("user").(*jwt.Token)
		claims := userJwt.Claims.(jwt.MapClaims)
		id = claims["id"].(string)
	}

	u := model.User{Id: id}

	if err := c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Sorry couldn't bind user value.")
	}

	fmt.Println("Update user...", u)
	err := u.Update(*h.DB)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Sorry couldn't update user.")
	}

	return c.JSON(http.StatusOK, make(map[string]string, 0))
}

func (h *Handler) UserDelete(c echo.Context) error {

	err := model.UserDelete(*h.DB, c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Sorry couldn't delete user.")
	}

	return c.JSON(http.StatusOK, make(map[string]string, 0))
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
	userJwt := c.Get("user").(*jwt.Token)
	claims := userJwt.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	//id := "123"
	fmt.Println("CLAIMS", claims)
	fmt.Println("ID", id)

	user := model.User{Id: id}
	err := user.FindById(*h.DB)
	if err != nil {
		fmt.Println("Couldn't load current user", err)
	}

	fmt.Println("USER", user)

	return c.String(http.StatusOK, stubData(h, user))
}

func stubData(h *Handler, user model.User) string {
	// Songs/files
	all_files, err := model.AllFiles(*h.DB)
	if err != nil {
		fmt.Println("Error fetching all files")
	}
	songArray, _ := json.Marshal(all_files)

	// Albums
	all_albums, err := model.AllPlaylists(*h.DB, model.PlaylistType_Album)
	if err != nil {
		fmt.Println("Error fetching all albums")
	}
	albumArray, _ := json.Marshal(all_albums)

	// Artists
	all_artists, err := model.AllPlaylists(*h.DB, model.PlaylistType_Artist)
	if err != nil {
		fmt.Println("Error fetching all artists")
	}
	artistArray, _ := json.Marshal(all_artists)

	// Playlists
	all_playlists, err := model.AllPlaylists(*h.DB, model.PlaylistType_User)
	if err != nil {
		fmt.Println("Error fetching all playlists")
	}

	playlistArray, _ := json.Marshal(all_playlists)

	// Interactions
	all_interactions, err := model.AllInteractions(*h.DB, all_files)
	if err != nil {
		fmt.Println("Error fetching all interactions")
	}

	interactionArray, _ := json.Marshal(all_interactions)

	// Users
	all_users, err := model.AllUsers(*h.DB)
	if err != nil {
		fmt.Println("Error fetching all users")
	}

	userArray, _ := json.Marshal(all_users)

	return `{
    "albums": ` + string(albumArray) + `,
    "artists": ` + string(artistArray) + `,
    "songs": ` + string(songArray) + `,
    "settings": {
        "media_path": "media"
    },
    "playlists": ` + string(playlistArray) + `,
    "interactions": ` + string(interactionArray) + `,
    "users": ` + string(userArray) + `,
    "currentUser": {
        "id": "` + user.Id + `",
        "name": "` + user.Name + `",
        "email": "` + user.Email + `",
        "is_admin": true,
        "preferences": []
    },
    "useLastfm": false,
    "useYouTube": false,
    "useiTunes": true,
    "allowDownload": false,
    "supportsTranscoding": false,
    "cdnUrl": "http://localhost:3000/",
    "currentVersion": "v0.1",
    "latestVersion": "v0.1"
}`
}
