package utils

import (
	"regexp"

	git "github.com/josa42/go-gitutils"
)

// Repo :
type Repo struct {
	hoster string
	user   string
	name   string
}

// GetRepo :
func GetRepo() Repo {

	repo := Repo{}

	remotes := git.Remotes()
	remote := remotes["origin"]

	if remote.Name != "" {
		re, _ := regexp.Compile(`^git@github.com:([^/]+)/([^/]+).git$`)
		result := re.FindStringSubmatch(remote.Fetch)
		if result != nil {
			repo.hoster = "github"
			repo.user = result[1]
			repo.name = result[2]
		}
	}

	return repo
}

// URL :
func (repo *Repo) URL() string {
	switch repo.hoster {
	case "github":
		return "https://github.com/" + repo.user + "/" + repo.name
	}

	return ""
}
