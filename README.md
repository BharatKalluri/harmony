# Harmony

> A tool to sync ZSH shell history across systems and across different shells.

## Installation

If you have go language installed on your system

```shell script
go get -u github.com/bharatkalluri/harmony
```

Use `harmony configure` to setup a harmony configuration

> Instructions for how to get a github token can be found
> [here](https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line),
> Make sure you select the "gist" checkbox on the permissions page.

Here is how a sample config file would look

```shell script
$ cat ~/.config/harmony/config
GITHUB_TOKEN=<your github token>
SHELL_HISTORY_PATH=/Users/bharatkalluri/.zsh_history
SHELL_TYPE=zsh
```

### Run

Harmony uses [secret gists](https://help.github.com/en/enterprise/2.13/user/articles/about-gists)
as a data store. Every time the command is run, it either updates the gist or updates the local shell history by looking
at which one was updated the latest.

To sync your history, just run

```shell script
harmony
```

and it should auto magically make sure all your shell history is in sync!

## Note to bash users

Currently, bash is not supported. Although the code exists in the repo, bash support is explicitly disabled as it is not
yet ready for production.

Bash support is on the roadmap and will be worked on soon!
