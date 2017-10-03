package controller

import (
	//"github.com/jackshapow/shapow/api/database"
	"github.com/dgraph-io/badger"
	"github.com/gocraft/dbr"
	_ "github.com/mattn/go-sqlite3"
)

// Handler is the handler for route controller actions
type Handler struct {
	// DB *pg.DB
	DB *dbr.Session
	KV *badger.KV
}
