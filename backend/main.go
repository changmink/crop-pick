package main

import (
	"./config"
	"./router"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

var AccessKeyID string
var SecretAccessKey string
var MyRegion string

func init() {
	config.LoadConfig("config.json")
}

func ConnectAws() *session.Session {
	AccessKeyID = config.AWS.AccessKeyId
	SecretAccessKey = config.AWS.Secret
	MyRegion = config.AWS.Region
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}

func main() {
	sess := ConnectAws()

	engine := gin.Default()

	engine.Use(func(c *gin.Context) {
		c.Set("sess", sess)
		c.Next()
	})

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
		v1.POST("/images", router.UploadImage)
	}

	engine.Run(":8080")
}
