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
	}

	engine.Run(":8080")
}
