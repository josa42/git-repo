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

var hosterExps = []string{
	`^git@(github).com:([^/]+)/([^/]+).git$`,
	`^https://(github).com/([^/]+)/([^/]+).git`,
}

// GetRepo :
func GetRepo() Repo {

	repo := Repo{}

	remotes := git.Remotes()
	remote := remotes["origin"]

	if remote.Name != "" {

		for _, exp := range hosterExps {
			re, _ := regexp.Compile(exp)
			result := re.FindStringSubmatch(remote.Fetch)
			if result != nil {
				repo.hoster = result[1]
				repo.user = result[2]
				repo.name = result[3]

				return repo
			}
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
