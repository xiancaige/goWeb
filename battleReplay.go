package main

import (
	"github.com/gin-gonic/gin"
	"goWeb/handle"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	//---上传---
	router.POST("/replay/:SessionName", handle.UploadStart)
	router.POST("/replay/:SessionName/users", handle.PostUsers)
	router.POST("/replay/:SessionName/file/replay.header", handle.PostHeader)
	router.POST("/replay/:SessionName/file/:Stream", handle.PostStream)
	router.POST("/replay/:SessionName/event", handle.PostEvent)
	router.POST("/replay/:SessionName/stopUploading", handle.PostStop)

	//---下载---
	router.POST("/replay/:SessionName/viewer/:viewerName", handle.Heartbeat)
	router.POST("/replay/:SessionName/startDownloading", handle.DownloadStart)
	router.GET("/replay/:SessionName/file/replay.header", handle.GetHeader)
	router.GET("/replay/:SessionName/event", handle.GetEvent)
	router.GET("/replay/:SessionName/file/:Stream", handle.GetStream)

	//---搜索---
	router.GET("/replay", handle.GetReplay)

	router.Run(":8011")
}
