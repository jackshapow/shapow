package main

import (
	"github.com/jackshapow/shapow/api/controller"
	"github.com/jackshapow/shapow/api/database"
	// "github.com/stryveapp/stryve-api/controller"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/badger"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httputil"
	"reflect"
)

// RegisterRoutes registers all API routes for the app
func RegisterRoutes(e *echo.Echo) {
	// badger
	opt := badger.DefaultOptions
	dir := "database/badger"
	opt.Dir = dir
	opt.ValueDir = dir
	kv, _ := badger.NewKV(&opt)
	fmt.Println("......")
	fmt.Println(reflect.TypeOf(kv))
	fmt.Println("......")

	// sqlite
	db, err := database.GetConnection()
	if err != nil {
		fmt.Println(err)
	}
	// env := &Env{db: db}

	h := &controller.Handler{
		DB: db,
		KV: kv,
	}

	e.POST("/", h.Home)

	e.POST("/me", h.UserLogin)
	e.DELETE("/me", h.UserLogout)
	e.GET("/data", h.UserData)

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
