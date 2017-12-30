package main

import (
	"github.com/jackshapow/shapow/api/controller"
	"github.com/jackshapow/shapow/api/model"
	//"github.com/jackshapow/shapow/api/database"
	// "github.com/stryveapp/stryve-api/controller"
	//"encoding/json"
	//"fmt"
	//"github.com/GeertJohan/go.rice"
	"github.com/dgraph-io/badger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	//"net/http/httputil"
	// "reflect"
	//"log"
)

// RegisterRoutes registers all API routes for the app
func RegisterRoutes(e *echo.Echo, db *badger.DB, node_settings *model.Node) {
	h := &controller.Handler{
		DB:           db,
		NodeSettings: node_settings,
	}

	// // Serving static/UI assets
	e.GET("/", echo.WrapHandler(http.FileServer(assets)))
	e.GET("/public/*", echo.WrapHandler(http.FileServer(assets)))

	// // Serving media assets
	e.Static("/media", node_settings.MediaPath())
	e.Static("/artwork", node_settings.ArtworkPath())
	// e.Use(middleware.Static(node_settings.MediaPath))

	// e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
	// 	Root:   (node_settings.MediaPath + "/"),
	// 	Browse: true,
	// }))

	//e.Use(middleware.Static(filepath.Join(*DataRoot, "Media")))

	// Login no authentication
	//e.Static("/api/media", node_settings.MediaPath)

	e.POST("/api/me", h.UserLogin)
	e.GET("/songs/:id/info", h.SongInfo)
	e.GET("/songs/:id/play", h.SongPlay)
	e.GET("/api/songs/:id/info", h.SongInfo)
	e.GET("/api/songs/:id/play", h.SongPlay)
	e.POST("/api/songs", h.SongUpload)

	// Require authentication
	r := e.Group("/api")
	r.Use(middleware.JWT([]byte("secret")))
	r.PUT("/me", h.UserUpdate)
	r.GET("/data", h.UserData)
	r.PUT("/user/:id", h.UserUpdate)
	r.DELETE("/me", h.UserLogout)
	r.GET("/datatest", h.UserDataTest)
	r.POST("/interaction/:kind", h.SongInteraction)
	r.POST("/playlist", h.PlaylistCreate)
	r.DELETE("/playlist/:id", h.PlaylistDelete)
	r.PUT("/playlist/:id", h.PlaylistUpdate)
	r.PUT("/playlist/:id/sync", h.PlaylistSync)
	r.POST("/user", h.UserCreate)
	r.DELETE("/user/:id", h.UserDelete)

	// e.POST("/metest", func(c echo.Context) error {
	// 	decoder := json.NewDecoder(c.Request().Body)
	// 	fmt.Println(decoder)
	// 	fmt.Println("------------------")
	// 	var v interface{}
	// 	json.NewDecoder(c.Request().Body).Decode(&v)
	// 	fmt.Println(v)
	// 	fmt.Println("------------------")
	// 	return c.String(http.StatusOK, "ok then!")
	// })

	// e.GET("/dumpdata", func(c echo.Context) error {

	// 	requestDump, err := httputil.DumpRequest(c.Request(), true)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(string(requestDump))

	// 	return c.String(http.StatusOK, "ok then!")
	// })

}
