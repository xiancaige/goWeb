package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goWeb/data"
	"net/http"
	"strconv"
	"strings"
)

func PostUsers(c *gin.Context) {
	defer c.Abort()
	fmt.Println("PostUsers  ------")

	c.JSON(http.StatusNoContent, gin.H{"sessionId": "1001"})
}

func UploadStart(c *gin.Context) {
	defer c.Abort()
	fmt.Println("PostStart  -----")

	var live data.LiveInfo
	if c.ShouldBindQuery(&live) == nil {
		data.StartLive(&live)
		c.JSON(http.StatusOK, gin.H{"sessionId": live.ViewerId})
	}
}

func PostHeader(c *gin.Context) {
	defer c.Abort()

	viewerId := c.Param("SessionName")
	file := "replay.header"

	data.WriteFileByHttp(c, viewerId, file)

	c.String(http.StatusNoContent, "")
}

func PostStream(c *gin.Context) {
	defer c.Abort()

	viewerId := c.Param("SessionName")
	file := c.Param("Stream")

	var ss data.StreamStruct
	if c.ShouldBindQuery(&ss) == nil {
		data.AddStream(viewerId, ss.NumChunks-1, ss)
	}

	data.WriteFileByHttp(c, viewerId, file)
	c.String(http.StatusNoContent, "")
}

func PostEvent(c *gin.Context) {
	defer c.Abort()

	group := c.Param("group")
	viewerId := c.Param("SessionName")

	var event data.ReplayEvent
	if c.ShouldBindQuery(&event) == nil {
		data.AddReplayEvent(viewerId, group, event)
	}

	c.String(http.StatusOK, "")
}

func PostStop(c *gin.Context) {
	defer c.Abort()

	fmt.Println("stopUploading  ------")
	viewerId := c.Param("SessionName")

	var ss data.StreamStruct
	if c.ShouldBindQuery(&ss) == nil {
		data.StopLive(viewerId, ss)
		c.String(http.StatusNoContent, "")
	}
}

func Heartbeat(c *gin.Context) {
	defer c.Abort()

	viewerName := c.Param("viewerName")
	fmt.Println("Heartbeat" + viewerName)
	c.String(http.StatusNoContent, "")
}

func DownloadStart(c *gin.Context) {
	defer c.Abort()

	fmt.Println("DownloadStart  ------")
	viewerId := c.Param("SessionName")

	c.JSON(http.StatusOK, data.GetLive(viewerId))
}

func GetHeader(c *gin.Context) {
	defer c.Abort()

	viewerId := c.Param("SessionName")

	file := "replay.header"

	data.ReadFileToHttp(c, viewerId, file)
}

func GetStream(c *gin.Context) {
	defer c.Abort()

	stream := c.Param("Stream")
	viewerId := c.Param("SessionName")

	live := data.GetLive(viewerId)
	if live == nil {
		return
	}

	ary := strings.Split(stream, ".")

	chunkIndex, _ := strconv.Atoi(ary[1])

	file := stream

	streamInfo, ok := data.GetStream(viewerId, chunkIndex)

	if !ok {
		return
	}

	c.Header("NumChunks", strconv.Itoa(live.NumChunks))
	c.Header("State", live.State)
	c.Header("Time", strconv.Itoa(live.Time))
	c.Header("MTime1", strconv.Itoa(streamInfo.MTime1))
	c.Header("MTime2", strconv.Itoa(streamInfo.MTime2))

	data.ReadFileToHttp(c, viewerId, file)

}

func GetEvent(c *gin.Context) {
	defer c.Abort()

	group := c.Param("group")
	viewerId := c.Param("SessionName")

	list := data.GetReplayEvent(viewerId, group)

	c.JSON(http.StatusOK, gin.H{"events": list})
	c.Abort()
}

func GetReplay(c *gin.Context) {
	defer c.Abort()

	var filter data.SelectFilter
	if c.ShouldBindQuery(&filter) == nil {

		var list = data.SelectReplay(&filter)
		c.JSON(http.StatusNoContent, list)
	}
}
