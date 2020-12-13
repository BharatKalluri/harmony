# Harmony
 > A tool to sync shell history across systemw. Supports bash and zsh!

## Installation

If you have go language installed on your system

```shell script
go get github.com/BharatKalluri/harmony
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

If you use bash as your primary shell, run this so that timestamps are also logged in bash history
Deleting shell history file is advised as the history file will anyways not have timestamps. 
Currently harmony breaks if the history file does not have timestamps (A bug is [filed](https://github.com/BharatKalluri/harmony/issues/2), I will fix it when I find time) 
```shell script
echo "HISTTIMEFORMAT=\"%s\"" >> ~/.bashrc
```

### Run

Harmony uses [secret gists](https://help.github.com/en/enterprise/2.13/user/articles/about-gists) 
as a data store. Every time the command is run, it either updates the gist or updates the local shell history
by looking at which one was updated the latest.

To sync your history, just run
```shell script
harmony
```

and it should auto magically make sure all your shell history is in sync!
