package models

import (
	"context"
	"log"

	"github.com/adwinugroho/etetika-go/config"
	"github.com/arangodb/go-driver"
)

type (
	DB struct {
		DBLive driver.Database
	}
)

type SessionUser struct {
	Token string `json:"token"`
}

func NewDAO(conn *config.ArangoDB) *DB {
	return &DB{
		DBLive: conn.DBLive,
	}
}

func (db *DB) GetUserByEmail(email string) (*User, error) {
	var ctx = context.Background()
	var result = User{}
	var query = `FOR x IN etetika_user FILTER x.email == %s RETURN x`
	cursor, err := db.DBLive.Query(ctx, query, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	for {
		_, err := cursor.ReadDocument(ctx, &result)
		if err != nil && driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Printf("Error reading document: %+v \n", err)
		}
	}
	return &result, nil
}
