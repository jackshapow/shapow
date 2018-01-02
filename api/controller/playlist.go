package controller

import (
	//"encoding/json"
	//"encoding/json"
	"fmt"
	"github.com/jackshapow/shapow/api/model"
	"github.com/labstack/echo"
	"net/http"
	//"reflect"
	//"time"
)

func (h *Handler) PlaylistDelete(c echo.Context) error {

	err := model.PlaylistDelete(*h.DB, c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Sorry couldn't delete playlist.")
	}

	return c.JSON(http.StatusOK, make(map[string]string, 0))
}

func (h *Handler) PlaylistUpdate(c echo.Context) error {
	p := model.Playlist{Id: c.Param("id")} //new(model.Playlist)

	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Sorry couldn't bind playlist value.")
	}

	fmt.Println("Update playlist...", p)
	err := p.Update(*h.DB)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Sorry couldn't update playlist.")
	}

	return c.JSON(http.StatusOK, make(map[string]string, 0))
}

func (h *Handler) PlaylistSync(c echo.Context) error {
	type Payload struct{ Songs []string }
	var pLoad Payload
	c.Bind(&pLoad)

	//  newCount, _ := model.SetInteraction(*h.DB, pLoad.Song, c.Param("kind")) //, ) // "play" or "favorite"

	fmt.Println("Updating playlist...")
	err := model.UpdatePlaylistSongs(*h.DB, c.Param("id"), pLoad.Songs)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Sorry couldn't update playlist.")
	}

	//  pLoad.Songs

	// fmt.Println("pload:", )
	// fmt.Println("pload:", reflect.TypeOf(pLoad.Songs))
	//playCount := map[string]uint32{"play_count": newCount}

	//fmt.Println("-------", newCount)
	return c.JSON(http.StatusOK, make(map[string]string, 0))
}

func (h *Handler) PlaylistCreate(c echo.Context) error {
	p := model.Playlist{Type: model.PlaylistType_UserType} //new(model.Playlist)

	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Sorry couldn't create playlist.")
	}

	fmt.Println("Creating playlist...")
	err := p.Create(*h.DB)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Sorry couldn't create playlist.")
	}

	fmt.Println("CREATED:", p)

	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Sorry couldn't create playlist.")
	}

	//fmt.Println("ZZZZZZZZZ", p)
	//{"name":"herper","id":1452,"songs":[]}

	return c.JSON(http.StatusOK, p)

	//  return c.JSON(http.StatusCreated, map[string]string{
	// 	"token": "jokin",
	// })

}
