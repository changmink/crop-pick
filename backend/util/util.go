package util

import (
	"log"

	"github.com/gin-gonic/gin"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ResultJSON(msg string, result interface{}) gin.H {
	return gin.H{
		"message": msg,
		"result":  result,
	}
}
