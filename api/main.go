package main

import (
	"fmt"
	// "github.com/jackshapow/shapow/api/database"
	"github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
	// "github.com/dgrijalva/jwt-go"
	// "github.com/gocraft/dbr"
	"github.com/dgraph-io/badger"
	"github.com/golang/protobuf/proto"
	"github.com/jackshapow/shapow/api/model"
	// "time"
	"encoding/binary"
	"os"
	"path/filepath"
	"strconv"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/skratchdot/open-golang/open"
)

func init() {
}

var (
	NodeSettings model.Node
)

func main() {
	// systray
	onExit := func() {
		fmt.Println("Finished onExit")
	}
	fmt.Println("here we go...")
	systray.Run(onReady, onExit)
	fmt.Println("xhere we go..")
}

func onReady() {
	fmt.Println("herrpppperrrr")
	systray.SetIcon(icon.Data)
	//systray.SetTitle("Awesome App")
	systray.SetTooltip("Lantern")
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	// We can manipulate the systray in other goroutines
	go func() {
		systray.SetIcon(icon.Data)
		//		systray.SetTitle("Awesome App")
		//systray.SetTooltip("Pretty awesome棒棒嗒")
		// mChange := systray.AddMenuItem("Change Me", "Change Me")
		// mChecked := systray.AddMenuItem("Unchecked", "Check Me")
		// mEnabled := systray.AddMenuItem("Enabled", "Enabled")
		//systray.AddMenuItem("Ignored", "Ignored")
		mUrl := systray.AddMenuItem("Open Shapow", "my home")
		//mQuit := systray.AddMenuItem("退出", "Quit the whole app")
		// systray.AddSeparator()
		// mToggle := systray.AddMenuItem("Toggle", "Toggle the Quit button")
		// shown := true
		for {
			select {
			// case <-mChange.ClickedCh:
			// 	mChange.SetTitle("I've Changed")
			// case <-mChecked.ClickedCh:
			// 	if mChecked.Checked() {
			// 		mChecked.Uncheck()
			// 		mChecked.SetTitle("Unchecked")
			// 	} else {
			// 		mChecked.Check()
			// 		mChecked.SetTitle("Checked")
			// 	}
			// case <-mEnabled.ClickedCh:
			// 	mEnabled.SetTitle("Disabled")
			// 	mEnabled.Disable()
			case <-mUrl.ClickedCh:
				open.Run("http://localhost:8080")
				// case <-mToggle.ClickedCh:
				// 	if shown {
				// 		mQuitOrig.Hide()
				// 		// mEnabled.Hide()
				// 		shown = false
				// 	} else {
				// 		mQuitOrig.Show()
				// 		// mEnabled.Show()
				// 		shown = true
				// 	}
				// case <-mQuit.ClickedCh:
				// 	systray.Quit()
				// 	fmt.Println("Quit2 now...")
				// 	return
			}
		}
	}()

	// Initialize badger
	opt := badger.DefaultOptions
	opt.Dir = "database/badger"
	opt.ValueDir = "database/badger"
	db, err := badger.Open(opt)

	if err != nil {
		fmt.Println(err)
	}

	// Initialize media store
	var node_settings = new(model.Node)
	err = db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("node:settings"))
		if err != nil {
			// Initial new settings
			node_settings.MediaPath = "media"
			nodeProto, _ := proto.Marshal(node_settings)
			err := txn.Set([]byte("node:settings"), nodeProto)
			if err != nil {
				return err
			}
		} else {
			data, err := item.Value()
			err = proto.Unmarshal(data, node_settings)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("ERROR LOADING NODE", err)
	}

	fmt.Println("GIDDYUP", node_settings)
	newpath := filepath.Join(".", node_settings.MediaPath)
	os.MkdirAll(newpath, os.ModePerm)

	// Reset DB?
	model.ResetDB(*db)

	e := echo.New()
	//	e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
	}))

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		//fmt.Println(string(reqBody))
		if binary.Size(reqBody) > 1000 {
			fmt.Println("Body request is", strconv.Itoa(binary.Size(reqBody)), "bytes")
		} else {
			fmt.Println("Request Body:", string(reqBody))
		}
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost", "http://localhost:3000", "http://localhost:8080", "http://localhost:3001", "http://koel.app"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	//e.Use(middleware.Recover())

	RegisterRoutes(e, db, node_settings)

	//test()
	e.Logger.Fatal(e.Start(":3001"))

}
