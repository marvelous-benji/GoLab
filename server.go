package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marvelous-benji/gin_proj/controller"
	"github.com/marvelous-benji/gin_proj/entity"
	"github.com/marvelous-benji/gin_proj/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func GetAllVideos(ctx *gin.Context) {
	var videos []entity.Video
	videos = videoController.Findall()
	ctx.JSON(200, gin.H{
		"status": "success",
		"videos": videos,
	})

}

func CreateVideo(ctx *gin.Context) {
	video := videoController.Save(ctx)
	ctx.JSON(201, gin.H{
		"status": "success",
		"msg":    "video created successfully",
		"video":  video,
	})
}

func main() {
	server := gin.Default()

	server.GET("/videos", GetAllVideos)
	server.POST("/videos", CreateVideo)

	server.Run()
}
