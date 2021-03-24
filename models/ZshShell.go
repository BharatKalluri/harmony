package models

import (
	"bufio"
	"bytes"
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
	splitOnSemiColon := strings.Split(strings.TrimSpace(encodedHistoryString), "0;")
	cmdInHistory := splitOnSemiColon[len(splitOnSemiColon)-1]
	splitOnColon := strings.Split(encodedHistoryString, ":")
	timeStampStr := strings.TrimSpace(splitOnColon[1])
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

	// remove escaped new lines and all \ chars, will flatten all history entries to one line
	scanner := bufio.NewScanner(
		bytes.NewReader(
			bytes.ReplaceAll(
				bytes.ReplaceAll(shellHistory, []byte("\\\n"), []byte("")),
				[]byte("\\"),
				[]byte(""),
			),
		),
	)
	var historyItemStrArr []string
	for scanner.Scan() {
		line := scanner.Text()
		historyItemStrArr = append(historyItemStrArr, line)
	}

	var historyItemsArr []HistoryItem
	for _, el := range historyItemStrArr {
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
