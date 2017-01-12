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
func (repo *Repo) URL(urlType string, arguments map[string]interface{}) string {

	switch repo.hoster {
	case "github":
		base := "https://github.com/" + repo.owner + "/" + repo.name
		switch urlType {
		case "issues":
			return base + "/issues"
		case "prs":
			return base + "/pulls"
		case "commits":
			return base + "/commits"
		case "compare":
			revA, revB := compareRevisions(arguments)
			return base + "/compare/" + revA + "..." + revB
		case "home":
			return base
		}

	case "bitbucket":
		base := "https://bitbucket.org/" + repo.owner + "/" + repo.name
		switch urlType {
		case "issues":
			return base + "/issues?status=new&status=open"
		case "prs":
			return base + "/pull-requests/"
		case "commits":
			return base + "/commits/all"
		case "compare":
			revA, revB := compareRevisions(arguments)
			return base + "/branches/compare/" + revB + ".." + revA + "#diff"
		case "home":
			return base
		}
	}

	return ""
}

func compareRevisions(arguments map[string]interface{}) (string, string) {

	revA := arguments["<older-revision>"]
	revB := arguments["<newer-revision>"]

	if revB == nil {
		revB, _ = git.Exec("rev-parse", "--verify", "HEAD")
	}

	return revA.(string), revB.(string)
}
