# Harmony

> A tool to sync ZSH shell history across systems and across different shells.

## Installation

If you have go language installed on your system

```shell script
go install github.com/bharatkalluri/harmony@latest
```

Use `harmony configure` to set up a harmony configuration. config lives at `~/.config/harmony/config`

> Harmony uses [secret gists](https://help.github.com/en/enterprise/2.13/user/articles/about-gists)
as a data store. Instructions for how to get a GitHub token can be found
> [here](https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line),
> Make sure you select the "gist" checkbox on the permissions page.


### Run

To sync your history, just run

```shell script
harmony
```

## Note to bash users

Currently, bash is not supported. Although the code exists in the repo, bash support is explicitly disabled as it is not
yet ready for production.

Bash support is on the roadmap and will be worked on soon!
