package config

import (
	"context"
	"log"
	"os"
	"time"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type (
	ArangoDB struct {
		DBLive driver.Database
	}
)

var (
	// Database Username
	DBUSERNAME = os.Getenv("DB_USERNAME")
	// Database Password
	DBPASSWORD = os.Getenv("DB_PASSWORD")
	// URL connection to database
	DBURL = os.Getenv("DB_URL")
	// Databse name
	DBName = os.Getenv("DB_NAME")
	// Port on running server
	PORT = os.Getenv("PORT_ETETIKA")
)

func NewArangoDBDatabase() *ArangoDB {
	ctx, cancel := NewArangoDBContext()
	defer cancel()

	// create a connection to DB
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{DBURL}, //DB Url
	})
	if err != nil {
		panic(err)
	}
	// create a new connection to DB client
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(DBUSERNAME, DBPASSWORD),
	})
	if err != nil {
		panic(err)
	}
	// connect client and DB
	db, err := client.Database(ctx, DBName)
	if err != nil {
		log.Printf("Error connecting to database, cause: %+v\n", err)
		panic(err)
	}

	return &ArangoDB{
		DBLive: db,
	}
}

func NewArangoDBContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
