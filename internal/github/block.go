package github

import (
	"fmt"
	"net/http"
)

var (
	blockUserURI   = GithubBlockURI{URI: "https://api.github.com/user/blocks/%s", Method: http.MethodPut}
	unblockUserURI = GithubBlockURI{URI: "https://api.github.com/user/blocks/%s", Method: http.MethodDelete}
)

func (g *githubClient) BlockGithubUser(username string, uri GithubBlockURI) error {
	request, err := g.newRequest(uri.Method, fmt.Sprintf(uri.URI, username))
	if err != nil {
		return err
	}
	return g.do(request, nil)
}
