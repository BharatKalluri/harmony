package models

import (
	"fmt"
	"strconv"
	"strings"
)

type ZSHShell struct{}

func (z ZSHShell) DecodeHistoryItem(encodedHistoryString string) (HistoryItem, error) {
	if !strings.Contains(encodedHistoryString, ";") {
		fmt.Println("Invalid history line: ", encodedHistoryString)
		return HistoryItem{}, nil
	}
	splitOnSemiColon := strings.Split(encodedHistoryString, ";")
	cmdInHistory := splitOnSemiColon[len(splitOnSemiColon)-1]
	splitOnColon := strings.Split(encodedHistoryString, ":")
	timeStampStr := splitOnColon[0]
	timeStamp, err := strconv.Atoi(strings.TrimSpace(timeStampStr))
	if err != nil {
		return HistoryItem{}, err
	}
	return HistoryItem{
		TimeStamp: timeStamp,
		Command:   cmdInHistory,
	}, nil
}

func (z ZSHShell) EncodeHistoryItem(historyItem HistoryItem) string {
	return fmt.Sprintf(": %d:0;%s", historyItem.TimeStamp, historyItem.Command)
}

func (z ZSHShell) GetShellHistoryFromBytes(shellHistory []byte) (ShellHistory, error) {
	historyItemsStrArr := strings.Split(string(shellHistory), ": ")
	var historyItemsArr []HistoryItem
	for _, el := range historyItemsStrArr {
		if len(el) > 1 {
			historyItem, err := z.DecodeHistoryItem(el)
			if err != nil {
				return ShellHistory{}, err
			}
			historyItemsArr = append(historyItemsArr, historyItem)
		}
	}
	return ShellHistory{History: historyItemsArr}, nil
}
