package main

import (
	docopt "github.com/docopt/docopt-go"
	"github.com/josa42/git-repo/utils"
	stringutils "github.com/josa42/go-stringutils"
	"github.com/skratchdot/open-golang/open"
)

func openRepo(repo utils.Repo, urlType string) {
	url := repo.URL()
	if url != "" {
		open.Run(url)
	}
}

func main() {
	usage := stringutils.TrimLeadingTabs(`
		Usage:
		  git-repo home [--issues|--code|--prs]

		Options:
		  -h --help          Show this screen.
		  --version          Show version.
  `)

	arguments, _ := docopt.Parse(usage, nil, true, "Git Release 0.1.0", false)

	homeCommand := arguments["home"] == true
	repo := utils.GetRepo()

	if homeCommand {
		openRepo(repo, "home")
	}
}
