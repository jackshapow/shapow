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
	"log"
)

// RegisterRoutes registers all API routes for the app
func RegisterRoutes(e *echo.Echo) {
	// badger
	// opt := badger.DefaultOptions
	// dir := "database/badger"
	// opt.Dir = dir
	// opt.ValueDir = dir
	// kv, _ := badger.NewKV(&opt)
	// fmt.Println("......")
	// fmt.Println(reflect.TypeOf(kv))
	// fmt.Println("......")

	// initialize badger
	opt := badger.DefaultOptions
	opt.Dir = "database/badger"
	opt.ValueDir = "database/badger"
	db, err := badger.Open(opt)
	if err != nil {
		log.Fatal(err)
	}

	// Reset DB?
	model.ResetDB(*db)

	//db.Close()
	//	defer db.Close()

	// sqlite
	// db, err := database.GetConnection()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // env := &Env{db: db}

	h := &controller.Handler{
		DB: db,
		//KV: kv,
	}

	e.POST("/", h.Home)
	e.Use(middleware.Static("media"))

	e.POST("/me", h.UserLogin)
	e.DELETE("/me", h.UserLogout)

	e.GET("/datatest", h.UserDataTest)

	//e.GET("/data", h.UserData)

	r := e.Group("/data")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", h.UserData)

	e.POST("/songs", h.SongUpload)
	e.GET("/songs/:id/info", h.SongInfo)
	e.GET("/songs/:id/play", h.SongPlay)

	e.POST("/user", h.UserCreate)

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
