package models

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/bharatkalluri/harmony/config"
	"github.com/google/go-github/v30/github"
	"golang.org/x/oauth2"
	"strconv"
	"strings"
)

// TODO: This could be cleaner, rethink

type ShellHistoryGist struct {
	Context context.Context
	Client  *github.Client
}

func NewShellHistoryGist() ShellHistoryGist {
	appConfig := config.ReadAppConfig()
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: appConfig.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return ShellHistoryGist{
		Context: ctx,
		Client:  client,
	}
}

func (s ShellHistoryGist) encodeHistoryItemToString(item HistoryItem) string {
	return fmt.Sprintf("%d:%s;", item.TimeStamp, item.Command)
}

func (s ShellHistoryGist) decodeHistoryItemFromString(item string) (HistoryItem, error) {
	splitOnColon := strings.Split(item, ":")
	timeStampStr := strings.TrimSpace(splitOnColon[0])
	timeStamp, err := strconv.Atoi(timeStampStr)
	if err != nil {
		return HistoryItem{}, err
	}
	command := strings.TrimSpace(splitOnColon[1])
	return HistoryItem{
		Command:   command,
		TimeStamp: timeStamp,
	}, nil
}

func (s ShellHistoryGist) GetShellHistoryGistObject(content string) github.Gist {
	shellHistoryFile := github.GistFile{
		Filename: github.String("shell_history"),
		Content:  github.String(content),
	}
	gistFileMap := make(map[github.GistFilename]github.GistFile)
	gistFileMap["shell_history"] = shellHistoryFile
	shellHistoryGist := github.Gist{
		Description: github.String("Shell history stored by harmony"),
		Public:      github.Bool(false),
		Files:       gistFileMap,
	}
	return shellHistoryGist
}

func (s ShellHistoryGist) ConvertShellHistoryToStringForGist(history ShellHistory) string {
	var shellHistoryStr string
	for _, el := range history.History {
		shellHistoryStr = shellHistoryStr + fmt.Sprintf("%s\n", s.encodeHistoryItemToString(el))
	}
	return shellHistoryStr
}

func (s ShellHistoryGist) GetShellHistoryFromStringForGist(shellHistoryStr string) (ShellHistory, error) {
	var shellHistoryArr []HistoryItem
	splitShellHistoryStr := strings.Split(shellHistoryStr, ";")
	for _, el := range splitShellHistoryStr {
		newLineStrippedEl := strings.ReplaceAll(el, "\n", "")
		if len(newLineStrippedEl) > 0 {
			decodedHistoryItem, err := s.decodeHistoryItemFromString(el)
			if err != nil {
				return ShellHistory{}, err
			}
			shellHistoryArr = append(shellHistoryArr, decodedHistoryItem)
		}
	}
	return ShellHistory{History: shellHistoryArr}, nil
}

func (s ShellHistoryGist) CreateShellHistoryGist(shellHistory ShellHistory) error {
	shellHistoryStr := s.ConvertShellHistoryToStringForGist(shellHistory)
	shellHistoryB64 := base64.StdEncoding.EncodeToString([]byte(shellHistoryStr))
	shellHistoryGist := s.GetShellHistoryGistObject(shellHistoryB64)
	_, _, err := s.Client.Gists.Create(s.Context, &shellHistoryGist)
	if err != nil {
		return err
	}
	return nil
}

func (s ShellHistoryGist) UpdateShellHistoryGist(shellHistory ShellHistory) error {
	updatedHistoryContents := s.ConvertShellHistoryToStringForGist(shellHistory)
	shellHistoryB64 := base64.StdEncoding.EncodeToString([]byte(updatedHistoryContents))
	shellHistoryGistDetails := s.ShellHistoryGistDetails()
	shellHistoryGist := s.GetShellHistoryGistObject(shellHistoryB64)
	_, _, err := s.Client.Gists.Edit(s.Context, *shellHistoryGistDetails.ID, &shellHistoryGist)
	return err
}

func (s ShellHistoryGist) ShellHistoryGistDetails() *github.Gist {
	gistList, _, _ := s.Client.Gists.List(s.Context, "", nil)
	for _, gist := range gistList {
		if (gist.Files[github.GistFilename("shell_history")] != github.GistFile{}) {
			gistFile, _, _ := s.Client.Gists.Get(s.Context, *gist.ID)
			return gistFile
		}
	}
	return nil
}

func (s ShellHistoryGist) GetShellHistoryFromGist() (ShellHistory, error) {
	shellHistoryGistDetails := s.ShellHistoryGistDetails()
	if shellHistoryGistDetails == nil {
		return ShellHistory{History: nil}, nil
	}
	encodedShellHistory := shellHistoryGistDetails.Files["shell_history"].Content
	if encodedShellHistory == nil {
		panic("Shell history file online is empty!")
	}
	decodedShellHistory, err := base64.StdEncoding.DecodeString(*encodedShellHistory)
	if err != nil {
		return ShellHistory{}, err
	}
	shellHistory, err := s.GetShellHistoryFromStringForGist(string(decodedShellHistory))
	if err != nil {
		return ShellHistory{}, err
	}
	return shellHistory, nil
}
