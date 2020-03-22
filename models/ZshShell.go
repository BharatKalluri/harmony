package models

import (
	"io/ioutil"
)

type ZSHShell struct {
	HistoryFilePath string
}

func (z ZSHShell) GetShellHistoryInBytes() ([]byte, error) {
	data, err := ioutil.ReadFile(z.HistoryFilePath)
	return data, err
}

func (z ZSHShell) GetShellHistory() (ShellHistory, error) {
	data, err := z.GetShellHistoryInBytes()
	if err != nil {
		return ShellHistory{}, err
	}
	return ShellHistory{}.GetShellHistoryFromBytes(data), nil
}

func (z ZSHShell) WriteShellHistory(shellHistory ShellHistory) error {
	shellHistoryStr := shellHistory.ConvertToString()
	err := ioutil.WriteFile(z.HistoryFilePath, []byte(shellHistoryStr), 0644)
	return err
}
