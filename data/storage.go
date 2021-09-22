package data

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"os"
)

var LiveMap = make(map[string]*LiveInfo)

func GetLive(viewerId string) *LiveInfo {
	live := getLive(viewerId)

	if live == nil {
		fmt.Println("不存在的播放uuid" + viewerId)
		return nil
	}

	return live
}

func StartLive(info *LiveInfo) {
	ul := uuid.NewV4()
	info.ViewerId = ul.String()
	info.State = "Live"

	//初始化map
	info.EventMap = make(map[string][]ReplayEvent)
	info.StreamMap = make(map[int]StreamStruct)

	LiveMap[info.ViewerId] = info
	fmt.Println("uuid生成" + info.ViewerId)
}

func StopLive(viewerId string, ss StreamStruct) {
	live := getLive(viewerId)

	if live == nil {
		fmt.Println("不存在的播放uuid" + viewerId)
		return
	}

	live.State = "Over"
	live.NumChunks = ss.NumChunks
	live.Time = ss.Time
	live.AbsSize = ss.AbsSize
}

func AddReplayEvent(viewerId string, group string, event ReplayEvent) {
	live := getLive(viewerId)
	if live == nil {
		return
	}

	ary := live.EventMap[group]
	if ary == nil {
		live.EventMap[group] = make([]ReplayEvent, 0)
		ary = live.EventMap[group]
	}

	ary = append(ary, event)
}

func GetReplayEvent(viewerId string, group string) []ReplayEvent {
	live := getLive(viewerId)
	if live == nil {
		return nil
	}

	return live.EventMap[group]
}

func AddStream(viewerId string, index int, ss StreamStruct) {
	live := getLive(viewerId)

	if live == nil {
		return
	}

	live.Time = ss.Time
	live.NumChunks = ss.NumChunks
	live.AbsSize = ss.AbsSize

	live.StreamMap[index] = ss
}

func GetStream(viewerId string, index int) (StreamStruct, bool) {
	live := getLive(viewerId)
	if live == nil {
		return StreamStruct{}, false
	}

	value, ok := live.StreamMap[index]
	return value, ok
}

func SelectReplay(filter *SelectFilter) []*LiveInfo {
	var replays = make([]*LiveInfo, 0)

	for _, v := range LiveMap {
		//if v.App == filter.App {
			replays = append(replays, v)
		//}
	}

	return replays
}

func getLive(viewerId string) *LiveInfo {
	//live, ok := LiveMap[viewerId]
	//if !ok {
	//	fmt.Println("不存在的播放uuid" + viewerId)
	//	return nil
	//}
	//
	//return live
	return LiveMap[viewerId]
}

var filePath = "d:/temp"

func WriteFileByHttp(c *gin.Context, viewerId, fileName string) {
	_, err := os.Stat(filePath + "/" + viewerId)
	if err != nil || !os.IsExist(err) {
		os.MkdirAll(filePath+"/"+viewerId, os.ModePerm)
	}

	file, err := os.Create(filePath + "/" + viewerId + "/" + fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	file.ReadFrom(c.Request.Body)
}

func ReadFileToHttp(c *gin.Context, viewerId, file string) {
	c.File(filePath + "/" + viewerId + "/" + file)
}
