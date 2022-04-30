package main

import (
	"log"

	db "main/src/database"
	routes "main/src/server"

	"github.com/joho/godotenv"
)

func main () {
	// abstract this to a server init func
	err := godotenv.Load(".env")
	if err != nil {
        log.Println("Error loading .env file")
    }
	db.Init()
	routes.InitializeRoutes()
}