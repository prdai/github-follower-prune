package github

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const UserFollowersURI string = "https://api.github.com/users/%s/followers"
const UserFollowingURI string = "https://api.github.com/users/%s/following"

func (g *githubClient) GetGitHubUsersPage(username string, uri string) *[]GithubUser {
	req, err := http.NewRequest("GET", fmt.Sprintf(uri, username), nil)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(0)
	}
	g.applySharedHeaders(req);
	res, err := g.client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(0)
	}
	var githubUsers []GithubUser
	g.extractResponseBodyContent(res, &githubUsers); 
	return &githubUsers
}

