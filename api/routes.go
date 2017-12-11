package main

import (
	"github.com/jackshapow/shapow/api/controller"
	"github.com/jackshapow/shapow/api/model"
	//"github.com/jackshapow/shapow/api/database"
	// "github.com/stryveapp/stryve-api/controller"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/badger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"net/http/httputil"
	// "reflect"
	//"log"
)

// RegisterRoutes registers all API routes for the app
func RegisterRoutes(e *echo.Echo, db *badger.DB, node_settings *model.Node) {
	h := &controller.Handler{
		DB:           db,
		NodeSettings: node_settings,
	}

	r := e.Group("")
	r.Use(middleware.JWT([]byte("secret")))

	e.POST("/", h.Home)
	e.Use(middleware.Static(NodeSettings.MediaPath))

	e.POST("/me", h.UserLogin)
	e.DELETE("/me", h.UserLogout)
	r.PUT("/me", h.UserUpdate)

	e.GET("/datatest", h.UserDataTest)

	r.GET("/data", h.UserData)

	e.POST("/songs", h.SongUpload)

	e.POST("/interaction/:kind", h.SongInteraction)

	e.POST("/playlist", h.PlaylistCreate)
	e.DELETE("/playlist/:id", h.PlaylistDelete)
	e.PUT("/playlist/:id", h.PlaylistUpdate)
	e.PUT("/playlist/:id/sync", h.PlaylistSync)

	e.GET("/songs/:id/info", h.SongInfo)
	e.GET("/songs/:id/play", h.SongPlay)

	e.POST("/user", h.UserCreate)
	e.DELETE("/user/:id", h.UserDelete)
	r.PUT("/user/:id", h.UserUpdate)

	e.POST("/zme", func(c echo.Context) error {
		decoder := json.NewDecoder(c.Request().Body)
		fmt.Println(decoder)
		fmt.Println("------------------")
		var v interface{}
		json.NewDecoder(c.Request().Body).Decode(&v)
		fmt.Println(v)
		fmt.Println("------------------")
		return c.String(http.StatusOK, "ok then!")
	})

	e.GET("/dumpdata", func(c echo.Context) error {

		requestDump, err := httputil.DumpRequest(c.Request(), true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))

		return c.String(http.StatusOK, "ok then!")
	})

}
