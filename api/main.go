package main

import (
	"fmt"
	// "github.com/jackshapow/shapow/api/database"
	"github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
	// "github.com/dgrijalva/jwt-go"
	// "github.com/gocraft/dbr"
	"github.com/dgraph-io/badger"
	// "github.com/golang/protobuf/proto"
	"github.com/jackshapow/shapow/api/model"
	// "time"
	// "encoding/binary"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/mitchellh/go-homedir"
	"github.com/skratchdot/open-golang/open"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path/filepath"
	// "reflect"
	// "github.com/shurcooL/vfsgen"
	// "net/http"
	"runtime"
	// "strconv"
)

func init() {
}

var (
	NodeSettings model.Node
	AppName      = "Bonfire"
	DataRoot     = kingpin.Flag("data", "User data directory").ExistingDir()
	ResetDB      = kingpin.Flag("reset", "Reset accounts and settings").Bool()
	Port         = kingpin.Flag("port", "Port to run webserver").String()
	//logger *log.Logger
)

func bootstrap() {
	fmt.Println("Initializing", AppName, "on", runtime.GOOS, "...")

	// Parse command line flags
	kingpin.Parse()

	if *Port == "" {
		if runtime.GOOS == "linux" {
			NodeSettings.Port = "80"
		} else {
			NodeSettings.Port = "31337"
		}
	}

	// Set file folders
	if *DataRoot != "" {
		// Use existing data directory
		// *DataRoot, _ = filepath.Abs(*DataRoot)
		NodeSettings.RootPath, _ = filepath.Abs(*DataRoot)
	} else {
		// Set default data directory
		root, _ := homedir.Dir()
		// *DataRoot = filepath.Join(root, AppName)
		NodeSettings.RootPath = filepath.Join(root, AppName)
		os.MkdirAll(NodeSettings.RootPath, os.ModePerm)
		os.MkdirAll(NodeSettings.MediaPath(), os.ModePerm)
	}

	fmt.Println("Data directory", NodeSettings.RootPath)

	// Set executable path
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	fmt.Println("Executing from", exPath)

	// Set FFMPEG location
	ffmpeg_path := filepath.Join(exPath, "binaries", runtime.GOOS, "ffmpeg")

	if _, err := os.Stat(ffmpeg_path); err == nil {
		fmt.Println("FFMPEG found ", ffmpeg_path)
		NodeSettings.FfmpegPath = ffmpeg_path
	}

	ffmpeg_path = filepath.Join(exPath, "ffmpeg")

	if _, err := os.Stat(ffmpeg_path); err == nil {
		fmt.Println("FFMPEG found ", ffmpeg_path)
		NodeSettings.FfmpegPath = ffmpeg_path
	}

}

func main() {
	bootstrap()

	// systray
	onExit := func() {
		fmt.Println(AppName, "quit.")
	}

	systray.Run(onReady, onExit)
}

func onReady() {

	// We can manipulate the systray in other goroutines
	go func() {
		systray.SetIcon(icon.Data)
		systray.SetTooltip(AppName)
		mUrl := systray.AddMenuItem(("Open " + AppName), "My home")
		mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
		systray.SetIcon(icon.Data)

		go func() {
			<-mQuitOrig.ClickedCh
			//fmt.Println("Quitting...")
			systray.Quit()
			//fmt.Println("Finished quitting")
		}()

		for {
			select {
			case <-mUrl.ClickedCh:
				open.Run("http://localhost:31337")
			}
		}
	}()

	// Initialize media store
	//var node_settings = new(model.Node)
	//var node_settings = model.InitializeNode()

	// err = db.Update(func(txn *badger.Txn) error {
	// 	item, err := txn.Get([]byte("node:settings"))
	// 	if err != nil {
	// 		// Initial new settings
	// 		node_settings.MediaPath = "media"
	// 		nodeProto, _ := proto.Marshal(node_settings)
	// 		err := txn.Set([]byte("node:settings"), nodeProto)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	} else {
	// 		data, err := item.Value()
	// 		err = proto.Unmarshal(data, node_settings)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}

	// 	return nil
	// })

	// if err != nil {
	// 	fmt.Println("ERROR LOADING NODE", err)
	// }

	// Initialize badger
	opt := badger.DefaultOptions
	opt.Dir = filepath.Join(NodeSettings.RootPath, "Database")
	opt.ValueDir = filepath.Join(NodeSettings.RootPath, "Database")
	db, err := badger.Open(opt)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Cannot open database, already in use. Exiting.")
		os.Exit(1)
	}

	// Reset DB?
	if *ResetDB == true {
		model.ResetDB(*db, &NodeSettings)
	}

	e := echo.New()
	//e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
	}))

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		//fmt.Println(string(reqBody))
		// if binary.Size(reqBody) > 1000 {
		// 	fmt.Println("Body request is", strconv.Itoa(binary.Size(reqBody)), "bytes")
		// } else {
		// 	fmt.Println("Request Body:", string(reqBody))
		// }
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost", "http://localhost:3000", "http://localhost:8080", "http://localhost:3001", "http://localhost:31337", "http://koel.app", "http://mixtape:31337", "http://mixtape"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))

	//e.Use(middleware.Recover())

	RegisterRoutes(e, db, &NodeSettings)

	if DevMode == false && runtime.GOOS != "linux" {
		open.Run("http://localhost:" + NodeSettings.Port)
	}

	e.HideBanner = true

	fmt.Println("Starting server http://localhost:" + NodeSettings.Port)

	e.Logger.Fatal(e.Start(":" + NodeSettings.Port))

}
