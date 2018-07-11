package git_artifact

import (
	"fmt"
	"github.com/flant/dapp/pkg/git_repo"
)

type GitArtifact struct {
	LocalGitRepo  *git_repo.LocalGitRepo
	RemoteGitRepo *git_repo.RemoteGitRepo

	Name               string
	As                 string
	Branch             string
	Commit             string
	Cwd                string
	Owner              string
	Group              string
	IncludePaths       []string
	ExcludePaths       []string
	StagesDependencies map[string][]string
	Paramshash         string
}

func (ga *GitArtifact) GitRepo() (git_repo.GitRepo, error) {
	if ga.LocalGitRepo != nil {
		return ga.LocalGitRepo, nil
	} else if ga.RemoteGitRepo != nil {
		return ga.RemoteGitRepo, nil
	}
	return nil, fmt.Errorf("GitRepo not initialized")
}

func (ga *GitArtifact) LatestCommit() (string, error) {
	gitRepo, err := ga.GitRepo()
	if err != nil {
		return "", err
	}
	return gitRepo.LatestCommit(ga.Branch)
}
