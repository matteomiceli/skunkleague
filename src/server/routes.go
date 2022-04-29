package routes

import (
	"fmt"
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

	r.GET("/players/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		player := db.GetPlayerById(id)
		fmt.Println(player)

		ctx.JSON(200, player)
	})


	r.POST("/players/add", func(ctx *gin.Context) {
		var newPlayer ReqNewPlayer
		
		if err := ctx.BindJSON(&newPlayer); err != nil {
			log.Fatal(err)
		}

		db.AddNewPlayer(newPlayer.Name)
		ctx.String(http.StatusOK, "New Player " + newPlayer.Name + " added!")
	})


	r.GET("/games", func(ctx *gin.Context) {
		games := db.GetAllGames()
		ctx.JSON(200, games)
	})


	
	r.Run()
}