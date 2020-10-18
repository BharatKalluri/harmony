package models

type HistoryItem struct {
	Command   string `json:"command"`
	TimeStamp int    `json:"time_stamp"`
}
