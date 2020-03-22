package models

import (
	"fmt"
	"strconv"
	"strings"
)

type HistoryItem struct {
	Command   string
	TimeStamp int
}

func (h HistoryItem) Encode() string {
	return fmt.Sprintf(": %d:0;%s", h.TimeStamp, h.Command)
}

func DecodeHistoryItem(encodedHistoryString string) (HistoryItem, error) {
	splitOnSemiColon := strings.Split(encodedHistoryString, ";")
	cmdInHistory := splitOnSemiColon[len(splitOnSemiColon)-1]
	splitOnColon := strings.Split(encodedHistoryString, ":")
	timeStampStr := splitOnColon[1]
	timeStamp, err := strconv.Atoi(strings.TrimSpace(timeStampStr))
	if err != nil {
		return HistoryItem{}, err
	}
	return HistoryItem{
		TimeStamp: timeStamp,
		Command:   cmdInHistory,
	}, nil
}
