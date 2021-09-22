package data

type UploadStruct struct {
	App          string `json:"app",form:"app"`
	Version      string `json:"version",form:"version"`
	CL           string `json:"cl",form:"cl"`
	FriendlyName string `json:"friendlyName",form:"friendlyName"`
}

type StreamStruct struct {
	NumChunks int `form:"numChunks"`
	Time      int `form:"time"`
	MTime1    int `form:"mTime1"`
	MTime2    int `form:"mTime2"`
	AbsSize   int `form:"absSize"`
}

type ReplayEvent struct {
	Id            string `from:"id"`
	Group         string `from:"group"`
	Meta          string `from:"meta"`
	Time1         int    `from:"time1"`
	Time2         int    `from:"time2"`
	IncrementSize bool   `from:"incrementSize"`
}

type SelectFilter struct {
	UploadStruct
	Meta   string `json:"Meta"`
	User   string `json:"User"`
	Recent string `json:"Recent"`
}
