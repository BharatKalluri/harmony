package models

import (
	"fmt"
	"strconv"
	"strings"
)

type BashShell struct{}

func (b BashShell) DecodeHistoryItem(encodedHistoryString string) (HistoryItem, error) {
	splitOnNewLine := strings.Split(encodedHistoryString, "\n")
	cmdInHistory := splitOnNewLine[len(splitOnNewLine)-2]
	timeStampStr := splitOnNewLine[0]
	timeStamp, err := strconv.Atoi(strings.TrimSpace(timeStampStr))
	if err != nil {
		return HistoryItem{}, err
	}
	return HistoryItem{
		TimeStamp: timeStamp,
		Command:   cmdInHistory,
	}, nil
}

func (b BashShell) EncodeHistoryItem(historyItem HistoryItem) string {
	return fmt.Sprintf("#%d\n%s", historyItem.TimeStamp, historyItem.Command)
}

func (b BashShell) GetShellHistoryFromBytes(shellHistory []byte) (ShellHistory, error) {
	historyItemsStrArr := strings.Split(string(shellHistory), "#")
	var historyItemsArr []HistoryItem
	for _, el := range historyItemsStrArr {
		if len(el) > 1 {
			historyItem, err := b.DecodeHistoryItem(el)
			if err != nil {
				return ShellHistory{}, err
			}
			historyItemsArr = append(historyItemsArr, historyItem)
		}
	}
	return ShellHistory{History: historyItemsArr}, nil
}
