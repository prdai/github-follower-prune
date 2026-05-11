package github

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const UserFollowersURI GithubUserListURI = "https://api.github.com/users/%s/followers"
const UserFollowingURI GithubUserListURI = "https://api.github.com/users/%s/following"

func (g *githubClient) getGitHubUsersPage(username string, uri GithubUserListURI) *[]GithubUser {
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
	var githubUsers []GithubUser
	g.extractResponseBodyContent(res, &githubUsers)
	return &githubUsers
}

func (g *githubClient) GetGitHubAllUsers(username string, uri GithubUserListURI) *[]GithubUser {

}
