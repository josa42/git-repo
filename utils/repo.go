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
func (repo *Repo) URL(urlType string) string {
	url := ""

	switch repo.hoster {
	case "github":

		switch urlType {
		case "issues":
			return "https://github.com/" + repo.user + "/" + repo.name + "/issues"
		case "prs":
			return "https://github.com/" + repo.user + "/" + repo.name + "/pulls"
		case "commits":
			return "https://github.com/" + repo.user + "/" + repo.name + "/commits"
		case "home":
		default:
			return "https://github.com/" + repo.user + "/" + repo.name
		}
	}

	return url
}
