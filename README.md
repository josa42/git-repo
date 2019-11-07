# git-repo

[![Build Status](https://github.com/josa42/git-repo/workflows/Test/badge.svg)](https://github.com/josa42/git-repo/actions?query=workflow%3ATest)
[![Go Report Card](https://goreportcard.com/badge/github.com/josa42/git-repo)](https://goreportcard.com/report/github.com/josa42/git-repo)
[![License](https://img.shields.io/github/license/josa42/git-repo.svg)](https://github.com/josa42/git-repo/blob/master/LICENSE)

## Installation

**Homebrew (macOS)**

```
brew tap josa42/homebrew-git-tools
brew install git-repo
```

**Other**

```
go get github.com/josa42/git-repo
```

## Usage

```
Usage:
  git-repo home
  git-repo commits
  git-repo issues
  git-repo prs
  git-repo pr
  git-repo compare <older-revision> [<newer-revision>]

Options:
  -h --help          Show this screen.
  --version          Show version.
```

## License

[MIT © Josa Gesell](LICENSE)
