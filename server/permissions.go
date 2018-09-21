package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

func (p *Plugin) checkUserPermissionsToRepo(userID string, repo *github.Repository) bool {
	info, err := p.getGitHubUserInfo(userID)
	if err != nil {
		fmt.Println(err.Message)
		return false
	}

	githubClient := p.githubConnect(*info.Token)
	fmt.Println(repo.GetID())
	newRepo, _, apiErr := githubClient.Repositories.Get(context.Background(), repo.GetOwner().GetLogin(), repo.GetName())
	if newRepo != nil {
		return true
	}

	if apiErr != nil {
		fmt.Println(apiErr.Error())
	}

	fmt.Println("nil return")
	return false
}
