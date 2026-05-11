package github

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const BlockUserURI string = "https://api.github.com/user/blocks/%s"

func (g *githubClient) BlockGithubUser(blockingUserName string) bool {
	req, err := http.NewRequest("PUT", fmt.Sprintf(BlockUserURI, blockingUserName), nil)
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
