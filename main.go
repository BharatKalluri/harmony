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
		panic("SHELL_TYPE needs to be either bash or zsh for now")
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
		fmt.Printf("Failed to retrieve local history: %s\n", err)
	}

	fmt.Printf("Found %d entries online\n", len(onlineHistory.History))
	fmt.Printf("Found %d entries locally\n", len(localHistory.History))
	mergedShellHistory := models.MergeShellHistories(onlineHistory, localHistory)
	fmt.Printf("Merged and Uploading %d entries, found %d common\n", len(mergedShellHistory.History), (len(onlineHistory.History)+len(localHistory.History))-len(mergedShellHistory.History))
	err = models.WriteLocalShellHistory(mergedShellHistory, shell)
	if err != nil {
		panic("Writing history to local has failed!")
	}
	err = shellHistoryGist.UpsertShellHistoryGist(mergedShellHistory)
	if err != nil {
		panic("Writing history to gist has failed!")
	}
	fmt.Println("Sync complete!")
}
