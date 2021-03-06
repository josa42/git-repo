package utils

import (
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"

	git "github.com/josa42/go-gitutils"
	"github.com/skratchdot/open-golang/open"
)

// Repo :
type Repo struct {
	hoster string
	owner  string
	name   string
}

var hosterExps = []string{
	`^git@(github).com:([^/]+)/([^/]+)\.git$`,
	`^git@(github).com:([^/]+)/([^/]+)$`,
	`^https://(github)\.com/([^/]+)/([^/]+)\.git$`,
	`^https://(github)\.com/([^/]+)/([^/]+)$`,
	`ssh://git@(bitbucket).org/([^/.]+)/([^/]+)\.git$`,
	`ssh://git@(bitbucket).org/([^/.]+)/([^/]+)$`,
	`^git@(bitbucket)\.org:([^/]+)/([^/]+)\.git$`,
	`^git@(bitbucket)\.org:([^/]+)/([^/]+)$`,
	`^https://[a-zA-Z0-9]+@(bitbucket).org/([^/]+)/([^/]+)\.git$`,
	`^https://[a-zA-Z0-9]+@(bitbucket).org/([^/]+)/([^/]+)$`,
	`^git@(gitlab).com:([^/]+)/([^/]+)\.git$`,
	`^git@(gitlab).com:([^/]+)/([^/]+)$`,
	`^https://(gitlab).com/([^/]+)/([^/]+)\.git$`,
	`^https://(gitlab).com/([^/]+)/([^/]+)$`,
}

// GetRepo :
func GetRepo() Repo {
	remotes := git.Remotes()

	// TODO allow specifying a remote --remote=upstream

	// first try origin remote or upstream
	prefered := []string{"origin", "upstream"}

	for _, name := range prefered {
		if remote, ok := remotes[name]; ok {
			if repo, err := getRepoFromRemote(remote.Fetch); err == nil {
				return repo
			}
		}
	}

	// try any other remote
	for _, remote := range remotes {
		if repo, err := getRepoFromRemote(remote.Fetch); err == nil {
			return repo
		}
	}

	return Repo{}
}

func (repo Repo) Open(urlType string, arguments map[string]interface{}) {
	url := repo.URL(urlType, arguments)
	if url != "" {
		open.Run(url)
	} else {
		fmt.Println("error: Could not match remote")
	}
}

func getRepoFromRemote(remote string) (Repo, error) {
	repo := Repo{}

	if remote != "" {
		for _, exp := range hosterExps {
			re, _ := regexp.Compile(exp)
			result := re.FindStringSubmatch(remote)
			if result != nil {
				repo.hoster = result[1]
				repo.owner = result[2]
				repo.name = result[3]

				return repo, nil
			}
		}
	}

	return repo, errors.New("not found")

}

// URL :
func (repo *Repo) URL(urlType string, arguments map[string]interface{}) string {

	if urlType == "ci" {
		return ciURL(repo, arguments)
	}

	switch repo.hoster {
	case "github":
		base := "https://github.com/" + repo.owner + "/" + repo.name
		switch urlType {
		case "issues":
			return base + "/issues"
		case "prs":
			return base + "/pulls"
		case "pr":
			return base + "/compare/" + git.CurrentBranch() + "?expand=1"
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
		case "pr":
			return base + "/pull-requests/new?source=" + git.CurrentBranch()
		case "commits":
			return base + "/commits/all"
		case "compare":
			revA, revB := compareRevisions(arguments)
			return base + "/branches/compare/" + revB + ".." + revA + "#diff"
		case "home":
			return base
		}

	case "gitlab":
		base := "https://gitlab.com/" + repo.owner + "/" + repo.name
		switch urlType {
		case "issues":
			return base + "/issues"
		case "prs":
			return base + "/pull-requests/"
		case "pr":
			return base + "/merge_requests"
		case "commits":
			return base + "/commits/" + git.CurrentBranch()
		case "compare":
			revA, revB := compareRevisions(arguments)
			return base + "/branches/compare/" + revA + "..." + revB + "#diff"
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

func getCiType(repo *Repo, arguments map[string]interface{}) string {
	ciType := ""
	ciTypes := []string{"appveyor", "bitbucket", "circle", "gitlab", "jenkins", "travis", "github"}

	for _, key := range ciTypes {
		if arguments["--"+key] == true {
			ciType = key
		}
	}

	if ciType == "" {
		ciType = detectCiType(repo, arguments)
	}

	switch ciType {
	case "travis", "appveyor", "circle":
		if repo.hoster != "github" {
			return ""
		}
	case "gitlab", "bitbucket":
		if repo.hoster != ciType {
			return ""
		}
	}

	return ciType
}

func detectCiType(repo *Repo, arguments map[string]interface{}) string {
	jenkinsURL, _ := git.Exec("config", "--get", "git-repo.jenkins-url")
	if jenkinsURL != "" {
		return "jenkins"
	}

	switch repo.hoster {
	case "bitbucket":
		switch findFileConfigFile() {
		case "bitbucket-pipelines.yml":
			return "bitbucket"
		}
	case "github":
		switch findFileConfigFile() {
		case ".travis.yml":
			return "travis"
		case "appveyor.yml":
			return "appveyor"
		case "circle.yml":
			return "circle"
		case ".github/workflows":
			return "github"
		}
	}
	return ""
}

func ciURL(repo *Repo, arguments map[string]interface{}) string {

	switch getCiType(repo, arguments) {
	case "travis":
		return "https://travis-ci.org/" + repo.owner + "/" + repo.name
	case "appveyor":
		return "https://ci.appveyor.com/project/" + repo.owner + "/" + repo.name
	case "circle":
		return "https://circleci.com/gh/" + repo.owner + "/" + repo.name
	case "gitlab":
		return "https://gitlab.com/" + repo.owner + "/" + repo.name + "/pipelines"
	case "bitbucket":
		return "https://bitbucket.org/" + repo.owner + "/" + repo.name + "/addon/pipelines/home"
	case "jenkins":
		url, _ := git.Exec("config", "--get", "git-repo.jenkins-url")
		if url == "" {
			fmt.Println("Error: Add jenkins project url:\n> git config --add http://jenkins.example.org/job/example-job/")
			return ""
		}
		return url
	case "github":
		return "https://github.com/" + repo.owner + "/" + repo.name + "/actions"
	}

	return ""
}

func findFileConfigFile() string {

	filePaths := []string{
		".travis.yml",
		"bitbucket-pipelines.yml",
		"appveyor.yml",
		"circle.yml",
		".github/workflows",
	}

	for _, filePath := range filePaths {
		if fileExists(filePath) {
			return filePath
		}
	}

	return ""
}

func fileExists(filePath string) bool {
	rootPath, _ := git.Exec("rev-parse", "--show-toplevel")
	if rootPath == "" {
		return false
	}
	_, err := os.Stat(path.Join(rootPath, filePath))
	return os.IsNotExist(err) != true
}
