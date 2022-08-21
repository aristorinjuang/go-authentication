package main

import (
	"log"

	"github.com/aristorinjuang/go-authentication/internal/config"
	httpDelivery "github.com/aristorinjuang/go-authentication/internal/delivery/http"
	"github.com/aristorinjuang/go-authentication/internal/infrastructure/database"
)

func main() {
	c, err := config.New()
	if err != nil {
		log.Panic(err)
	}

	db, err := database.ConnectMySQL(&c.Database)
	if err != nil {
		log.Panic(err)
	}

	log.Panic(httpDelivery.Run(c, database.NewMySQL(db)))
}
