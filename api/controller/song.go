package controller

import (
	//"encoding/json"
	//"fmt"
	// "github.com/dgrijalva/jwt-go"
	// "github.com/gocraft/dbr"
	// "github.com/jackshapow/shapow/api/model"
	"github.com/labstack/echo"
	"net/http"
	// "reflect"
	// "time"
	//"os"
)

// var data string

func (h *Handler) SongInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, stubSongInfo())
}

func (h *Handler) SongPlay(c echo.Context) error {
	filename := "/Users/jack/go/src/github.com/jackshapow/shapow/api/music/Vampire Weekend - Modern Vampires Of The City [2013] 320/Modern Vampires Of The City @ 320/03 Step.mp3"
	//filename := "/Users/jack/go/src/github.com/jackshapow/shapow/api/music/test.mp3"
	// fmt.Println("xxxxxxxxxxxxx")
	// fmt.Println(c.Param("id"))
	// fmt.Println("xxxxxxxxxxxxx")
	// c.Response().Header().Set("Content-Type", "audio/mp3")

	// option 1
	return c.Attachment(filename, "file.mp3")

	// option 2
	// f, err := os.Open(filename)
	// if err != nil {
	// 	return err
	// }
	// return c.Stream(http.StatusOK, "audio/mp3", f)

}

func stubSongInfo() string {
	return `{
    "lyrics": "",
    "album_info": false,
    "artist_info": false,
    "youtube": false
}`
}
