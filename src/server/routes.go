package routes

import (
	"fmt"
	"log"
	db "main/src/database"
	static "main/src/static"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ReqNewPlayer struct {
	Player db.Player
	AccessCode string 
}

type ReqNewGame struct {
	Game db.Game
	AccessCode string 
}

type ReqGameResult struct {
	Game db.GameResult
	AccessCode string
}

func InitializeRoutes() {
	r := gin.Default()
	r.Use(cors.Default())


	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, static.Welcome)
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

	r.GET("/players/:id/games", func(ctx *gin.Context) {
		id := ctx.Param("id")

		games := db.GetPlayerMatchHistory(id)

		ctx.JSON(200, games)
	})

	r.POST("/players/add", func(ctx *gin.Context) {
		var newPlayer ReqNewPlayer
		
		if err := ctx.BindJSON(&newPlayer); err != nil {
			log.Fatal(err)
		}

		if newPlayer.AccessCode == os.Getenv("ACCESS_CODE") {
			db.CreatePlayer(newPlayer.Player)
			ctx.String(http.StatusOK, "New Player " + newPlayer.Player.Alias + " added!")
		} else {
			ctx.Status(http.StatusUnauthorized)
		}
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

		if newGame.AccessCode == os.Getenv("ADMIN_CODE") {
			db.CreateGame(newGame.Game)
			ctx.String(http.StatusOK, "New game created!")
		} else {
			ctx.Status(http.StatusUnauthorized)
		}
	})


	r.PATCH("/games/result", func(ctx *gin.Context) {
		var gameResult ReqGameResult

		if err := ctx.BindJSON(&gameResult); err != nil {
			log.Fatal(err)
		}

		if gameResult.AccessCode == os.Getenv("ACCESS_CODE") {
			db.FinalizeGameResult(gameResult.Game)
			ctx.String(http.StatusOK, "Game results reported!")
		} else {
			ctx.Status(http.StatusUnauthorized)
		}
	})


	r.Run()
}