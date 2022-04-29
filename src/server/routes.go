package routes

import (
	"fmt"
	"log"
	db "main/src/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReqNewPlayer struct {
	Player db.Player
	AccessCode int
}

type ReqNewGame struct {
	Game db.Game
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

		db.CreatePlayer(newPlayer.Player)
		ctx.String(http.StatusOK, "New Player " + newPlayer.Player.Alias + " added!")
	})


	r.GET("/games", func(ctx *gin.Context) {
		games := db.GetAllGames()
		ctx.JSON(200, games)
	})


	r.POST("/games/add", func(ctx *gin.Context) {
		var newGame ReqNewGame

		if err := ctx.BindJSON(&newGame); err != nil {
			log.Fatal(err)
		}

		db.CreateGame(newGame.Game)
	})


	r.Run()
}