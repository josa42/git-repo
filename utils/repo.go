package utils

import (
	"regexp"

	git "github.com/josa42/go-gitutils"
)

// Repo :
type Repo struct {
	hoster string
	owner  string
	name   string
}

var hosterExps = []string{
	`^git@(github).com:([^/]+)/([^/]+).git$`,
	`^https://(github).com/([^/]+)/([^/]+).git`,
	`^git@(bitbucket).org:([^/]+)/([^/]+).git$`,
	`^https://([a-zA-Z0-9]+)@(bitbucket).org/([^/]+)/([^/]+).git`,
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
				repo.owner = result[2]
				repo.name = result[3]

				return repo
			}
		}
	}

	return repo
}

// URL :
func (repo *Repo) URL(urlType string) string {

	switch repo.hoster {
	case "github":
		switch urlType {
		case "issues":
			return "https://github.com/" + repo.owner + "/" + repo.name + "/issues"
		case "prs":
			return "https://github.com/" + repo.owner + "/" + repo.name + "/pulls"
		case "commits":
			return "https://github.com/" + repo.owner + "/" + repo.name + "/commits"
		case "home":
			return "https://github.com/" + repo.owner + "/" + repo.name
		}

	case "bitbucket":
		switch urlType {
		case "issues":
			return "https://bitbucket.org/" + repo.owner + "/" + repo.name + "/issues?status=new&status=open"
		case "prs":
			return "https://bitbucket.org/" + repo.owner + "/" + repo.name + "/pull-requests/"
		case "commits":
			return "https://bitbucket.org/" + repo.owner + "/" + repo.name + "/commits/all"
		case "home":
			return "https://bitbucket.org/" + repo.owner + "/" + repo.name
		}
	}

	return ""
}
