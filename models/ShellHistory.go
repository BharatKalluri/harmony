package models

import (
	"fmt"
	"strings"
)

type ShellHistory struct {
	History []HistoryItem
}

func (s ShellHistory) GetShellHistoryAfter(completeShellHistory ShellHistory, timestamp int) []HistoryItem {
	var shellHistoryAfterTimestamp []HistoryItem
	for _, el := range completeShellHistory.History {
		if el.TimeStamp > timestamp {
			shellHistoryAfterTimestamp = append(shellHistoryAfterTimestamp, el)
		}
	}
	return shellHistoryAfterTimestamp
}

func (s ShellHistory) ConvertToString() string {
	var shellHistoryStr string
	for _, el := range s.History {
		shellHistoryStr = shellHistoryStr + fmt.Sprintf("%s\n", el.Encode())
	}
	return shellHistoryStr
}

func (s ShellHistory) GetShellHistoryFromBytes(shellHistory []byte) ShellHistory {
	historyItemsStrArr := strings.Split(string(shellHistory), "\n")
	var historyItemsArr []HistoryItem
	for _, el := range historyItemsStrArr {
		if len(el) > 1 {
			historyItem, err := DecodeHistoryItem(el)
			if err != nil {
				return ShellHistory{}
			}
			historyItemsArr = append(historyItemsArr, historyItem)
		}
	}
	return ShellHistory{History: historyItemsArr}
}

func (s ShellHistory) GetLatestUpdatedTime() int {
	var max int
	for i, el := range s.History {
		if i == 0 || el.TimeStamp > max {
			max = el.TimeStamp
		}
	}
	return max
}
