package data

type LiveInfo struct {
	UploadStruct
	ViewerId  string                   `json:"viewerId"`
	NumChunks int                      `json:"numChunks"`
	Time      int                      `json:"time"`
	State     string                   `json:"state"`
	AbsSize   int                      `json:"absSize"`
	EventMap  map[string][]ReplayEvent `json:"EventMap,omitempty"`
	StreamMap map[int]StreamStruct     `json:"StreamMap,omitempty"`
}

