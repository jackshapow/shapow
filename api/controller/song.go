package controller

import (
	//"encoding/json"
	"fmt"
	// "github.com/dgrijalva/jwt-go"
	//	"github.com/gocraft/dbr"
	"github.com/jackshapow/shapow/api/model"
	"github.com/jackshapow/shapow/api/util"
	"github.com/labstack/echo"
	"net/http"
	//"reflect"
	// "time"
	"io"
	"os"
	//"strconv"
	//"strings"
	"path/filepath"
	//"time"
)

func (h *Handler) SongInteraction(c echo.Context) error {
	type Payload struct{ Song string }
	var pLoad Payload
	c.Bind(&pLoad)

	interaction, _ := model.SetInteraction(*h.DB, pLoad.Song, c.Param("kind")) //, ) // "play" or "favorite"

	return c.JSON(http.StatusOK, interaction)
}

func (h *Handler) SongUpload(c echo.Context) error {
	// Read form fields
	//name := c.FormValue("name")
	//email := c.FormValue("email")

	//------------
	// Read files
	//------------

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["file"]

	for _, file := range files {
		// artificial delay time.Sleep(1000 * time.Millisecond)
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		fmt.Println("OK", h.NodeSettings.MediaPath)
		// Destination
		full_path := filepath.Join(h.NodeSettings.MediaPath(), file.Filename)
		os.MkdirAll(filepath.Dir(full_path), os.ModePerm)
		dst, err := os.Create(full_path)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		// fmt.Println("----------")
		// fmt.Println(reflect.TypeOf(dst))
		// fmt.Println("----------")

		// Process
		// files.Process(pathname)
		file := model.File{Path: full_path}
		file.Import(*h.DB, h.NodeSettings)

	}

	//	return c.HTML(http.StatusOK, fmt.Sprintf("<p>Uploaded successfully %d files with fields name=%s and email=%s.</p>", len(files), name, email))

	return c.JSON(http.StatusOK, "")
}

func (h *Handler) SongInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, stubSongInfo())
}

func (h *Handler) SongPlay(c echo.Context) error {
	//idInt64, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	file := model.File{Id: c.Param("id")}
	err := file.FindById(*h.DB)
	if err != nil {
		fmt.Println("oh shit couldnt find it")
	}
	filename := util.Basepath() + "/" + file.Path
	filename = file.Path
	fmt.Println("Playing: " + filename)

	// filename := "/Users/jack/go/src/github.com/jackshapow/shapow/api/music/Vampire Weekend - Modern Vampires Of The City [2013] 320/Modern Vampires Of The City @ 320/03 Step.mp3"
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
