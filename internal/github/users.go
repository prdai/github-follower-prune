package github

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const UserProfileURI GithubUserURI = "https://api.github.com/users/%s"
const UserFollowingURI GithubUserURI = "https://api.github.com/users/%s/following"

func (g *githubClient) GetGitHubUser(username string, uri GithubUserURI) GithubUser {
	req, err := http.NewRequest("GET", fmt.Sprintf(string(uri), username), nil)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(0)
	}
	g.applySharedHeaders(req)
	res, err := g.client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(0)
	}
	var githubUsers GithubUser
	g.extractResponseBodyContent(res, &githubUsers)
	return githubUsers
}

func (g *githubClient) getGitHubFollowingPage(username string, page int, uri GithubUserURI) *[]GithubDetailedUser {
	req, err := http.NewRequest("GET", fmt.Sprintf(string(uri), username), nil)
	queries := req.URL.Query()
	queries.Set("page", string(page))
	queries.Set("per_page", "100")
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(0)
	}
	g.applySharedHeaders(req)
	res, err := g.client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(0)
	}
	var githubUsers []GithubDetailedUser
	g.extractResponseBodyContent(res, &githubUsers)
	return &githubUsers
}

func (g *githubClient) GetGitHubFollowing(username string, uri GithubUserURI) *GithubDetailedUsers {
	ghDetailedUsers := GithubDetailedUsers{}
	return &ghDetailedUsers
}
