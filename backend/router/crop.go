package router

import (
	"../db"
	"github.com/gin-gonic/gin"
)

func Search(ctx *gin.Context) {
	name := ctx.Query("name")
	info := db.FindInfo(name)
	ctx.JSON(200, gin.H{"data": info})
}

func Ranking(ctx *gin.Context) {

}
