package github

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/prdai/github-follower-prune/internal/types"
)

var SharedGhClientHeaders = map[string]string{
	"Accept":               "application/vnd.github+json",
	"X-GitHub-Api-Version": "2026-03-10",
	"Authorization":        fmt.Sprintf("Bearer %s", os.Getenv("GH_TOKEN")),
}

type githubClient struct {
	client        *http.Client
	sharedHeaders map[string]string
	userName      string
}

func InitGithubClient(config *types.Config) *githubClient {
	return &githubClient{
		client:        &http.Client{},
		sharedHeaders: SharedGhClientHeaders,
		userName:      config.UserName,
	}
}

func (g *githubClient) applySharedHeaders(request *http.Request) {
	for headerName, headerValue := range g.sharedHeaders {
		request.Header.Set(headerName, headerValue)
	}
}

func (g *githubClient) extractResponseBodyContent(response *http.Response, responseStructure any) {
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(0)
	}
	err = json.Unmarshal(body, responseStructure)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(0)
	}
}
