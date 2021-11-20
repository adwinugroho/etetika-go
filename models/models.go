package models

import (
	"github.com/adwinugroho/etetika-go/config"
	"github.com/arangodb/go-driver"
)

type (
	DB struct {
		DBLive driver.Database
	}
)

func NewDAO(conn *config.ArangoDB) *DB {
	return &DB{
		DBLive: conn.DBLive,
	}
}
