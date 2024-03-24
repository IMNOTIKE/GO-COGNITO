package service

import (
	"cognito/books/config"
	"cognito/books/orm"

	"github.com/go-pg/pg/v10"
)

var db *pg.DB

func Connect(config config.AppConfig) {
	db = orm.Connect(config)
}

func Close() {
	db.Close()
}
