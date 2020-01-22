package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

////////////////////////////////////////////////////////////////////////////////
// github

func TestGetRepoFromRemote_httpsGithub(t *testing.T) {
	r := getRepoFromRemote("https://github.com/josa42/git-repo")

	assert.Equal(t, "github", r.hoster)
	assert.Equal(t, "josa42", r.owner)
	assert.Equal(t, "git-repo", r.name)
}

func TestGetRepoFromRemote_httpsGithub_git(t *testing.T) {
	r := getRepoFromRemote("https://github.com/josa42/git-repo.git")

	assert.Equal(t, "github", r.hoster)
	assert.Equal(t, "josa42", r.owner)
	assert.Equal(t, "git-repo", r.name)
}

func TestGetRepoFromRemote_sshGithub(t *testing.T) {
	r := getRepoFromRemote("git@github.com:josa42/git-repo")

	assert.Equal(t, "github", r.hoster)
	assert.Equal(t, "josa42", r.owner)
	assert.Equal(t, "git-repo", r.name)
}

func TestGetRepoFromRemote_sshGithub_git(t *testing.T) {
	r := getRepoFromRemote("git@github.com:josa42/git-repo.git")

	assert.Equal(t, "github", r.hoster)
	assert.Equal(t, "josa42", r.owner)
	assert.Equal(t, "git-repo", r.name)
}

//
func TestGetRepoFromRemote_sshGithubWithDots_git(t *testing.T) {
	r := getRepoFromRemote("git@github.com:josa42/josa42.github.io.git")

	assert.Equal(t, "github", r.hoster)
	assert.Equal(t, "josa42", r.owner)
	assert.Equal(t, "josa42.github.io", r.name)
}

////////////////////////////////////////////////////////////////////////////////
// bitbucket

func TestGetRepoFromRemote_httpsBitbucket_git(t *testing.T) {
	r := getRepoFromRemote("https://user@bitbucket.org/owner/name.git")

	assert.Equal(t, "bitbucket", r.hoster)
	assert.Equal(t, "owner", r.owner)
	assert.Equal(t, "name", r.name)
}

func TestGetRepoFromRemote_httpsBitbucket(t *testing.T) {
	r := getRepoFromRemote("https://user@bitbucket.org/owner/name")

	assert.Equal(t, "bitbucket", r.hoster)
	assert.Equal(t, "owner", r.owner)
	assert.Equal(t, "name", r.name)
}

func TestGetRepoFromRemote_sshBitbucket_git(t *testing.T) {
	r := getRepoFromRemote("git@bitbucket.org:owner/name.git")

	assert.Equal(t, "bitbucket", r.hoster)
	assert.Equal(t, "owner", r.owner)
	assert.Equal(t, "name", r.name)
}

func TestGetRepoFromRemote_sshBitbucket(t *testing.T) {
	r := getRepoFromRemote("git@bitbucket.org:owner/name")

	assert.Equal(t, "bitbucket", r.hoster)
	assert.Equal(t, "owner", r.owner)
	assert.Equal(t, "name", r.name)
}

func TestGetRepoFromRemote_sshAltBitbucket_git(t *testing.T) {
	r := getRepoFromRemote("ssh://git@bitbucket.org/owner/name.git")

	assert.Equal(t, "bitbucket", r.hoster)
	assert.Equal(t, "owner", r.owner)
	assert.Equal(t, "name", r.name)
}

func TestGetRepoFromRemote_sshAltBitbucket(t *testing.T) {
	r := getRepoFromRemote("ssh://git@bitbucket.org/owner/name")

	assert.Equal(t, "bitbucket", r.hoster)
	assert.Equal(t, "owner", r.owner)
	assert.Equal(t, "name", r.name)
}

////////////////////////////////////////////////////////////////////////////////
// gitlab

func TestGetRepoFromRemote_httpsGitlab(t *testing.T) {
	r := getRepoFromRemote("https://gitlab.com/josa42/git-repo")

	assert.Equal(t, "gitlab", r.hoster)
	assert.Equal(t, "josa42", r.owner)
	assert.Equal(t, "git-repo", r.name)
}

func TestGetRepoFromRemote_httpsGitlab_git(t *testing.T) {
	r := getRepoFromRemote("https://gitlab.com/josa42/git-repo.git")

	assert.Equal(t, "gitlab", r.hoster)
	assert.Equal(t, "josa42", r.owner)
	assert.Equal(t, "git-repo", r.name)
}

func TestGetRepoFromRemote_sshGitlab(t *testing.T) {
	r := getRepoFromRemote("git@gitlab.com:josa42/git-repo")

	assert.Equal(t, "gitlab", r.hoster)
	assert.Equal(t, "josa42", r.owner)
	assert.Equal(t, "git-repo", r.name)
}

func TestGetRepoFromRemote_sshGitlab_git(t *testing.T) {
	r := getRepoFromRemote("git@gitlab.com:josa42/git-repo.git")

	assert.Equal(t, "gitlab", r.hoster)
	assert.Equal(t, "josa42", r.owner)
	assert.Equal(t, "git-repo", r.name)
}
