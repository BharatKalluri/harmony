package main

import (
	"fmt"
	"github.com/bharatkalluri/harmony/config"
	"github.com/bharatkalluri/harmony/models"
	"strings"
)

func GetShellTypeFromStr(shellTypeStr string) models.Shell {
	if strings.Compare(shellTypeStr, "bash") == 0 {
		return models.BashShell{}
	} else if strings.Compare(shellTypeStr, "zsh") == 0 {
		return models.ZSHShell{}
	} else {
		panic("SHELL_TYPE needs to be either bash or zsh")
	}
}

func main() {
	appConfig := config.ReadAppConfig()

	shell := GetShellTypeFromStr(appConfig.ShellType)

	shellHistoryGist := models.NewShellHistoryGist()
	onlineHistory, err := shellHistoryGist.GetShellHistoryFromGist()
	if err != nil {
		panic("Failed to retrieve history from gist")
	}
	localHistory, err := models.GetShellHistory(shell)
	if err != nil {
		panic("Failed to retrieve local history")
	}

	onlineHistoryLastUpdateOn := onlineHistory.GetLatestUpdatedTime()
	localHistoryLastUpdateOn := localHistory.GetLatestUpdatedTime()

	fmt.Println("Starting shell history sync...")

	if onlineHistoryLastUpdateOn == 0 {
		// No history exists online, push
		err := shellHistoryGist.CreateShellHistoryGist(localHistory)
		if err != nil {
			panic("Failed to push history!")
		}
		fmt.Println("Successfully pushed all history")
	} else if localHistoryLastUpdateOn == 0 {
		// No history exists local, pull
		err := models.WriteShellHistory(onlineHistory, shell)
		if err != nil {
			panic("Pull from online history has failed!")
		}
		fmt.Println("Successfully pulled all history")
	} else if onlineHistoryLastUpdateOn > localHistoryLastUpdateOn {
		// Online history is ahead of the local history, pull
		missingHistory := onlineHistory.GetShellHistoryAfter(onlineHistory, localHistoryLastUpdateOn)
		totalHistory := append(missingHistory, localHistory.History...)
		// Write total history to local file
		err := models.WriteShellHistory(models.ShellHistory{History: totalHistory}, shell)
		if err != nil {
			panic("Failed to pull history changes")
		}
		fmt.Println("Successfully pulled history")
	} else if localHistoryLastUpdateOn > onlineHistoryLastUpdateOn {
		// Local history is ahead of online history, push
		missingHistory := localHistory.GetShellHistoryAfter(localHistory, onlineHistoryLastUpdateOn)
		totalHistory := append(missingHistory, onlineHistory.History...)
		err = shellHistoryGist.UpdateShellHistoryGist(models.ShellHistory{History: totalHistory})
		if err != nil {
			panic("Failed to push history changes")
		}
		fmt.Println("Successfully pushed history")
	} else if localHistoryLastUpdateOn == onlineHistoryLastUpdateOn {
		fmt.Println("All in sync!")
	}

}
