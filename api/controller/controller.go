package controller

import (
	"github.com/dgraph-io/badger"
	"github.com/jackshapow/shapow/api/model"
	// "github.com/gocraft/dbr"
	// _ "github.com/mattn/go-sqlite3"
)

// Handler is the handler for route controller actions
type Handler struct {
	// DB *pg.DB
	DB           *badger.DB
	NodeSettings *model.Node
	// KV *badger.KV
}
