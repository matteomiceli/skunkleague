package routes

import (
	db "main/src/database"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		db.GetAllPlayers()
		
		ctx.JSON(200, map[string]string{"hi": "hi"})
	})
	
	r.Run()
}