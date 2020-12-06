package router

import (
	"../db"
	"../util"
	"github.com/gin-gonic/gin"
)

func SearchCrop(ctx *gin.Context) {
	name := ctx.Query("name")
	info, err := db.FindCropInfo(name)
	if err != nil {
		ctx.JSON(400, gin.H{"result": err})
		return
	}
	ctx.JSON(200, util.ResultJSON("Success", info))
}

func RankingCrop(ctx *gin.Context) {
	count, err := db.GetCropCount()
	if err != nil {
		ctx.JSON(400, gin.H{"result": err})
		return
	}
	ctx.JSON(200, util.ResultJSON("Success", count))
}
