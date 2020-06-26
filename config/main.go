package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type AppConfig struct {
	GithubToken      string
	ShellHistoryPath string
	ShellType        string
}

func ReadAppConfig() AppConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath("$HOME/.config/harmony")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Println("Hey there! I could not find a config file over at ~/.config/harmony/config")
		fmt.Println("It just needs to have three variables, A GITHUB_TOKEN, SHELL_HISTORY_PATH and SHELL_TYPE")
		fmt.Println("For more detailed explanation, visit https://github.com/BharatKalluri/harmony/blob/master/README.md")
		os.Exit(0)
	}
	githubToken := viper.GetString("GITHUB_TOKEN")
	if githubToken == "" {
		panic("This program cannot function without GITHUB_TOKENn in ~/.config/harmony/config !")
	}
	shellHistoryPath := viper.GetString("SHELL_HISTORY_PATH")
	if shellHistoryPath == "" {
		panic("This program cannot function without SHELL_HISTORY_PATH in ~/.config/harmony/config !")
	}
	shellType := viper.GetString("SHELL_TYPE")
	if shellType == "" {
		panic("Need shell type to be set to either bash or zsh!")
	}
	return AppConfig{GithubToken: githubToken, ShellHistoryPath: shellHistoryPath, ShellType: shellType}
}
