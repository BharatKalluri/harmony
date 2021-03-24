package models

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BashShell struct{}

func (b BashShell) DecodeHistoryItem(encodedHistoryString string) (HistoryItem, error) {
	// TODO: this is a very sensitive function and will blow up with a panic if something does not go well. Fix.
	// TODO: does not even support comments!
	splitOnNewLine := strings.Split(encodedHistoryString, "\n")
	cmdInHistory := splitOnNewLine[len(splitOnNewLine)-2]
	timeStampStr := splitOnNewLine[0]
	timeStamp, err := strconv.Atoi(strings.TrimSpace(timeStampStr))
	if err != nil {
		fmt.Println("Looks like the timestamp is not present in the bash history for item \"", cmdInHistory, "\". Exiting!")
		fmt.Println("Refer: https://github.com/BharatKalluri/harmony/blob/master/README.md#note-to-bash-users")
		os.Exit(1)
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
	// TODO: Need to figure this out, read about HISTTIMEFORMAT
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
