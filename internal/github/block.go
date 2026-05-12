package github

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var BlockUserURI GithubBlockURI = GithubBlockURI{URI: "https://api.github.com/user/blocks/%s", Method: "PUT"}
var UnBlockUserURI GithubBlockURI = GithubBlockURI{URI: "https://api.github.com/user/blocks/%s", Method: "DELETE"}

func (g *githubClient) BlockGithubUser(blockingUserName string, uri GithubBlockURI) bool {
	req, err := http.NewRequest(uri.Method, fmt.Sprintf(string(uri.URI), blockingUserName), nil)
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
	if res.StatusCode != 204 {
		return false
	}
	return true
}
