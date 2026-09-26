package github

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	userProfileURI   GithubUserURI = "https://api.github.com/users/%s"
	userFollowersURI GithubUserURI = "https://api.github.com/users/%s/followers"
	followersPerPage               = 100
)

func (g *githubClient) GetGitHubUser(username string) (GithubUser, error) {
	var user GithubUser
	request, err := g.newRequest(http.MethodGet, fmt.Sprintf(string(userProfileURI), username))
	if err != nil {
		return user, err
	}
	if err := g.do(request, &user); err != nil {
		return user, err
	}
	return user, nil
}

func (g *githubClient) GetGitHubFollowers(username string) ([]GithubFollower, error) {
	var followers []GithubFollower
	for page := 1; ; page++ {
		pageFollowers, err := g.getGitHubFollowersPage(username, page)
		if err != nil {
			return nil, err
		}
		followers = append(followers, pageFollowers...)
		if len(pageFollowers) < followersPerPage {
			return followers, nil
		}
	}
}

func (g *githubClient) getGitHubFollowersPage(username string, page int) ([]GithubFollower, error) {
	request, err := g.newRequest(http.MethodGet, fmt.Sprintf(string(userFollowersURI), username))
	if err != nil {
		return nil, err
	}
	query := request.URL.Query()
	query.Set("page", strconv.Itoa(page))
	query.Set("per_page", strconv.Itoa(followersPerPage))
	request.URL.RawQuery = query.Encode()
	var followers []GithubFollower
	if err := g.do(request, &followers); err != nil {
		return nil, err
	}
	return followers, nil
}
