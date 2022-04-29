package routes

import (
	"log"
	db "main/src/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReqNewPlayer struct {
	Name string
	AccessCode int
}

func InitializeRoutes() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hi there")
	})

	r.GET("/players", func(ctx *gin.Context) {
		players := db.GetAllPlayers()
		ctx.JSON(200, players)
	})

	r.POST("/players/add", func(ctx *gin.Context) {
		var newPlayer ReqNewPlayer
		
		if err := ctx.BindJSON(&newPlayer); err != nil {
			log.Fatal(err)
		}

		db.AddNewPlayer(newPlayer.Name)
		ctx.String(http.StatusOK, "New Player " + newPlayer.Name + " added!")
	})
	
	r.Run()
}