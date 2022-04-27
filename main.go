package main

import (
	"log"

	db "github.com/matteomiceli/skunkleague/src/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main () {
	// abstract this to a server init func
	err := godotenv.Load(".env")
	if err != nil {
        log.Fatal("Error loading .env file")
    }

	db.Init()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"hi": "hi"})
	})
	
	r.Run()
}