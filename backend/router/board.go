package router

import (
	"strconv"

	"../config"
	"../db"
	"../model"
	"../util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

func GetBoardList(ctx *gin.Context) {
	boardList, err := db.GetBoardList()
	if err != nil {
		ctx.JSON(400, gin.H{"result": err.Error()})
	}

	ctx.JSON(200, util.ResultJSON("Success", boardList))
}

func GetBoard(ctx *gin.Context) {
	pageString := ctx.Query("page")
	pageRangeString := ctx.Query("range")
	name := ctx.Query("name")

	count, err := db.GetPostCount(name)
	if err != nil {
		ctx.JSON(400, gin.H{"result": err.Error()})
		return
	}

	pageRange, err := strconv.ParseInt(pageRangeString, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"result": err.Error()})
		return
	}

	page, err := strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"result": err.Error()})
		return
	}

	start := (page - 1) * pageRange
	list, err := db.GetPostList(name, start, pageRange)
	if err != nil {
		ctx.JSON(400, gin.H{"result": err.Error()})
		return
	}

	totalPage := count / pageRange
	if count%pageRange != 0 {
		totalPage += 1
	}
	boardPage := model.BoardPage{totalPage, list}
	ctx.JSON(200, util.ResultJSON("Success", boardPage))

}

func AddPost(ctx *gin.Context) {
	var post model.Post
	err := ctx.BindJSON(&post)
	if err != nil {
		ctx.JSON(400, gin.H{"result": err.Error()})
		return
	}

	id, err := db.AddPost(post)
	if err != nil {
		ctx.JSON(400, gin.H{"result": err.Error()})
		return
	}
	ctx.JSON(200, util.ResultJSON("Success", id))
}

func UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")
	var post model.Post
	err := ctx.BindJSON(&post)
	if err != nil {
		ctx.JSON(400, gin.H{"result": err.Error()})
		return
	}

	err = db.UpdatePost(post, id)
	if err != nil {
		ctx.JSON(400, gin.H{"result": err.Error()})
		return
	}
	ctx.JSON(200, util.ResultJSON("updated", ""))
}

func LikedPost(ctx *gin.Context) {
	id := ctx.Param("id")
	err := db.LikedPost(id)
	if err != nil {
		ctx.JSON(400, gin.H{"result": err.Error()})
		return
	}
	ctx.JSON(200, util.ResultJSON("updated", ""))
}

func GetPost(ctx *gin.Context) {
	id := ctx.Param("id")

	post, err := db.GetPost(id)
	if err != nil {
		ctx.JSON(400, gin.H{"result": err.Error()})
		return
	}

	ctx.JSON(200, util.ResultJSON("Success", post))
}

func AddComment(ctx *gin.Context) {
	var comment model.Comment
	err := ctx.BindJSON(&comment)
	if err != nil {
		ctx.JSON(400, gin.H{"result": err.Error()})
		return
	}

	id, err := db.AddComment(comment)
	if err != nil {
		ctx.JSON(400, gin.H{"result": err.Error()})
	}
	ctx.JSON(200, util.ResultJSON("Success", id))
}

func UpdateComment(ctx *gin.Context) {
	var comment model.Comment
	err := ctx.BindJSON(&comment)
	if err != nil {
		ctx.JSON(400, gin.H{"result": err.Error()})
		return
	}

	err = db.UpdateComment(comment)
	if err != nil {
		ctx.JSON(400, gin.H{"result": err.Error()})
		return
	}
	ctx.JSON(200, util.ResultJSON("updated", ""))
}

func UploadImage(c *gin.Context) {
	sess := c.MustGet("sess").(*session.Session)
	uploader := s3manager.NewUploader(sess)

	MyBucket := config.AWS.BucketName
	MyRegion := config.AWS.Region

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{
			"result": err.Error(),
		})
		return
	}
	filename := header.Filename

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(MyBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		c.JSON(400, gin.H{
			"result": err.Error(),
		})
		return
	}
	filepath := "https://" + MyBucket + "." + "s3-" + MyRegion + ".amazonaws.com/" + filename

	c.JSON(200, util.ResultJSON("Success", filepath))
}
