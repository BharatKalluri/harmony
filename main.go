package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bharatkalluri/harmony/config"
	"github.com/bharatkalluri/harmony/models"
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

func initiateSync(shellType string) {
	shell := GetShellTypeFromStr(shellType)
	shellHistoryGist, _ := models.NewShellHistoryGist()
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

func main() {
	appConfig, err := config.ReadAppConfig()
	args := os.Args[1:]
	if len(args) == 0 {
		if err == config.MissingConfigFileError {
			fmt.Println("Hey there! I could not find a config file over at ~/.config/harmony/config")
			fmt.Println("Use `harmony config` to set up the configuration file")
			fmt.Println("For more detailed explanation, visit https://github.com/BharatKalluri/harmony/blob/master/README.md")
			os.Exit(0)
			return
		}
		initiateSync(appConfig.ShellType)
		return
	}
	if args[0] == "configure" {
		config.WriteAppConfig()
	}
}
