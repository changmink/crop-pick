package main

import (
	"./config"
	"./router"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfig("config.json")
}

func main() {
	engine := gin.Default()
	api := engine.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.GET("/search", router.SearchCrop)
		v1.GET("/ranking", router.RankingCrop)

		v1.GET("/posts", router.GetBoard)
		v1.POST("/posts", router.AddPost)
		v1.PUT("/posts/:id", router.UpdatePost)
		v1.PUT("/posts/:id/liked", router.LikedPost)
		v1.GET("/posts/:id", router.GetPost)
		v1.POST("/comment", router.AddComment)
		v1.PUT("/comment", router.UpdateComment)
	}

	engine.Run(":8080")
}
