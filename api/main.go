package main

import (
	"fmt"
	"github.com/jackshapow/shapow/api/database"
	"github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
	// "github.com/dgrijalva/jwt-go"
	"github.com/gocraft/dbr"
	"github.com/jackshapow/shapow/api/model"
	// "time"
)

func init() {
	db, err := database.GetConnection()
	if err != nil {
		panic(err)
	}

	u := model.User{Name: dbr.NewNullString("Bobby Jenkins"), Email: dbr.NewNullString("bjenkins@gmail.comz"), Password: dbr.NewNullString("passwordhere")}
	_, err = db.InsertInto("users").Columns("name", "email", "password").Record(&u).Exec()

	if err != nil {
		//panic(err)
	}

}

func main() {
	e := echo.New()
	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
	}))

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Println(string(reqBody))
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost", "http://localhost:3000", "http://localhost:8080", "http://localhost:3001", "http://koel.app"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	//e.Use(middleware.Recover())

	RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":3001"))
}
