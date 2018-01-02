package controller

import (
	"fmt"
	"github.com/dgraph-io/badger"
	"github.com/gorilla/websocket"
	"github.com/jackshapow/shapow/api/model"
	"github.com/labstack/echo"
	"net/http"
)

// Handler is the handler for route controller actions
type Handler struct {
	// DB *pg.DB
	DB           *badger.DB
	NodeSettings *model.Node
	// KV *badger.KV
}

var (
	upgrader = websocket.Upgrader{}
)

func (h *Handler) Websocket(c echo.Context) error {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {

		// Write
		//err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err = ws.WriteJSON(map[string]int{"Message": 12345}); err != nil {
			fmt.Println(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}
