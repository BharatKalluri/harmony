package models

import (
	"fmt"
	"github.com/bharatkalluri/harmony/config"
	"io/ioutil"
	"sort"
)

type ShellHistory struct {
	History []HistoryItem
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

func WriteLocalShellHistory(shellHistory ShellHistory, shell Shell) error {
	shellHistoryStr := shellHistory.ConvertToString(shell)
	appConfig := config.ReadAppConfig()
	err := ioutil.WriteFile(appConfig.ShellHistoryPath, []byte(shellHistoryStr), 0644)
	return err
}

func addMap(a map[int]string, b map[int]string) {
	for k, v := range b {
		a[k] = v
	}
}

func MergeShellHistories(shellHistoryOne ShellHistory, shellHistoryTwo ShellHistory) ShellHistory {
	shellHistoryOneMap := make(map[int]string)
	if shellHistoryOne.History != nil {
		for _, el := range shellHistoryOne.History {
			shellHistoryOneMap[el.TimeStamp] = el.Command
		}
	}
	shellHistoryTwoMap := make(map[int]string)
	if shellHistoryTwo.History != nil {
		for _, el := range shellHistoryTwo.History {
			shellHistoryTwoMap[el.TimeStamp] = el.Command
		}
	}

	addMap(shellHistoryOneMap, shellHistoryTwoMap)
	var mergedShellHistory []HistoryItem
	for k, v := range shellHistoryOneMap {
		mergedShellHistory = append(mergedShellHistory, HistoryItem{
			Command:   v,
			TimeStamp: k,
		})
	}
	sort.Slice(mergedShellHistory[:], func(i, j int) bool {
		return mergedShellHistory[i].TimeStamp < mergedShellHistory[j].TimeStamp
	})
	return ShellHistory{History: mergedShellHistory}
}
