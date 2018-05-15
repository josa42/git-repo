package main

import (
	docopt "github.com/docopt/docopt-go"
	"github.com/josa42/git-repo/utils"
	stringutils "github.com/josa42/go-stringutils"
	"github.com/skratchdot/open-golang/open"
)

func openRepo(repo utils.Repo, urlType string, arguments map[string]interface{}) {
	url := repo.URL(urlType, arguments)
	if url != "" {
		open.Run(url)
	}
}

func main() {
	usage := stringutils.TrimLeadingTabs(`
		Usage:
		  git-repo home
		  git-repo commits
		  git-repo issues
		  git-repo prs
		  git-repo pr
		  git-repo compare <older-revision> [<newer-revision>]
		  git-repo ci [--appveyor][--bitbucket][--circle][--gitlab][--jenkins][--travis]

		Options:
		  -h --help          Show this screen.
		  --version          Show version.
  `)

	arguments, _ := docopt.Parse(usage, nil, true, "Git Release 0.5.6", false)

	repo := utils.GetRepo()

	urlType := ""
	urlTypes := []string{"home", "issues", "prs", "pr", "commits", "compare", "ci"}

	for _, key := range urlTypes {
		if arguments[key] == true {
			urlType = key
		}
	}

	if urlType != "" {
		openRepo(repo, urlType, arguments)
	}
}
