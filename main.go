package main

import (
	"fmt"
	"log"

	"github.com/adwinugroho/etetika-go/config"
	"github.com/adwinugroho/etetika-go/models"
	"github.com/adwinugroho/etetika-go/routes"
	"github.com/adwinugroho/etetika-go/services"
	"github.com/labstack/echo/v4"
)

func main() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
	e := echo.New()
	// tes cpnnection DB
	getConnect := config.NewArangoDBDatabase()
	fmt.Println(getConnect)
	// get models
	getModels := models.NewDAO(getConnect)
	// get services
	getServices := services.NewServices(getModels)
	// get static for assets
	e.Static("/", "views")

	routes.Init(e, getServices)

	//log.Printf("listen server on port %+v\n", config.PORT)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.PORT)))

}
