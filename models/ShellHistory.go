package models

import (
	"fmt"
	"github.com/bharatkalluri/harmony/config"
	"io/ioutil"
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

func (s ShellHistory) ConvertToString(shell Shell) string {
	var shellHistoryStr string
	for _, el := range s.History {
		shellHistoryStr = shellHistoryStr + fmt.Sprintf("%s\n", shell.EncodeHistoryItem(el))
	}
	return shellHistoryStr
}

func GetShellHistoryInBytes() ([]byte, error) {
	appConfig := config.ReadAppConfig()
	data, err := ioutil.ReadFile(appConfig.ShellHistoryPath)
	return data, err
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

func GetShellHistory(shell Shell) (ShellHistory, error) {
	data, err := GetShellHistoryInBytes()
	if err != nil {
		return ShellHistory{}, err
	}
	shellHistory, err := shell.GetShellHistoryFromBytes(data)
	if err != nil {
		return ShellHistory{}, err
	}
	return shellHistory, nil
}

func WriteShellHistory(shellHistory ShellHistory, shell Shell) error {
	shellHistoryStr := shellHistory.ConvertToString(shell)
	appConfig := config.ReadAppConfig()
	err := ioutil.WriteFile(appConfig.ShellHistoryPath, []byte(shellHistoryStr), 0644)
	return err
}
