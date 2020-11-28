package router

import (
	"../db"
	"github.com/gin-gonic/gin"
)

func SearchCrop(ctx *gin.Context) {
	name := ctx.Query("name")
	info := db.FindCropInfo(name)
	ctx.JSON(200, gin.H{"result": info})
}

func RankingCrop(ctx *gin.Context) {
	count := db.GetCropCount()
	ctx.JSON(200, gin.H{"result": count})
}
