package config

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

type AppConfig struct {
	GithubToken      string
	ShellHistoryPath string
	ShellType        string
}

var MissingConfigFileError = errors.New("config not found")

func getConfigBasePath() string {
	userHomeDir, _ := os.UserHomeDir()
	joinedFolderPath := path.Join(userHomeDir, ".config", "harmony")
	return joinedFolderPath
}

func getInputFromUser(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic("Could not read from stdio")
	}
	return strings.Trim(text, "\n")
}

func WriteAppConfig() {
	configBasePath := getConfigBasePath()
	configPath := path.Join(configBasePath, "/config")
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(getConfigBasePath(), 0700)
		if err != nil {
			fmt.Println("Something went wrong while creating the config directory at ", configBasePath)
			return
		}
		githubToken := getInputFromUser("Github token: ")
		shellType := getInputFromUser("Shell type (zsh or bash): ")
		shellHistoryPath := getInputFromUser("Shell history file path: ")
		configContents := fmt.Sprintf("GITHUB_TOKEN=%s\nSHELL_HISTORY_PATH=%s\nSHELL_TYPE=%s", githubToken, shellHistoryPath, shellType)
		err = ioutil.WriteFile(configPath, []byte(configContents), 0644)
		if err != nil {
			fmt.Println("Was not able to write to config at ", configBasePath)
			return
		}
	} else {
		fmt.Println("Config file already exists at " + configPath + "!")
		return
	}
}

func ReadAppConfig() (AppConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(getConfigBasePath())
	err := viper.ReadInConfig()
	if err != nil {
		return AppConfig{}, MissingConfigFileError
	}
	githubToken := viper.GetString("GITHUB_TOKEN")
	if githubToken == "" {
		panic("This program cannot function without GITHUB_TOKENn in ~/.config/harmony/config !")
	}
	shellHistoryPath := viper.GetString("SHELL_HISTORY_PATH")
	if shellHistoryPath == "" {
		panic("This program cannot function without SHELL_HISTORY_PATH in ~/.config/harmony/config !")
	} else {
		if strings.HasPrefix(shellHistoryPath, "~/") {
			userHomeDir, _ := os.UserHomeDir()
			shellHistoryPath = path.Join(userHomeDir, strings.TrimPrefix(shellHistoryPath, "~/"))
		}
	}
	shellType := viper.GetString("SHELL_TYPE")
	if shellType == "" {
		panic("Need shell type to be set to either bash or zsh!")
	}
	return AppConfig{GithubToken: githubToken, ShellHistoryPath: shellHistoryPath, ShellType: shellType}, nil
}
