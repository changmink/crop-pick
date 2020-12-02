package util

import (
	"github.com/gin-gonic/gin"
)

func ResultJSON(msg string, result interface{}) gin.H {
	return gin.H{
		"message": msg,
		"result":  result,
	}
}
