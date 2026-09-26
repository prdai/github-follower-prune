package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/prdai/github-follower-prune/internal/types"
)

const githubAPIVersion = "2026-03-10"

type githubClient struct {
	client        *http.Client
	sharedHeaders map[string]string
	userName      string
}

func NewGithubClient(config *types.Config) (*githubClient, error) {
	token := os.Getenv("GH_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("GH_TOKEN is not set: add it to .env or export it")
	}
	return &githubClient{
		client: &http.Client{},
		sharedHeaders: map[string]string{
			"Accept":               "application/vnd.github+json",
			"X-GitHub-Api-Version": githubAPIVersion,
			"Authorization":        fmt.Sprintf("Bearer %s", token),
		},
		userName: config.UserName,
	}, nil
}

func (g *githubClient) applySharedHeaders(request *http.Request) {
	for headerName, headerValue := range g.sharedHeaders {
		request.Header.Set(headerName, headerValue)
	}
}

func (g *githubClient) newRequest(method string, uri string) (*http.Request, error) {
	request, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("building %s %s request: %w", method, uri, err)
	}
	g.applySharedHeaders(request)
	return request, nil
}

func (g *githubClient) do(request *http.Request, responseBody any) error {
	response, err := g.client.Do(request)
	if err != nil {
		return fmt.Errorf("%s %s: %w", request.Method, request.URL, err)
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode > 299 {
		body, _ := io.ReadAll(io.LimitReader(response.Body, 512))
		return fmt.Errorf("%s %s: unexpected status %d: %s", request.Method, request.URL, response.StatusCode, body)
	}
	if responseBody == nil {
		return nil
	}
	if err := json.NewDecoder(response.Body).Decode(responseBody); err != nil {
		return fmt.Errorf("decoding %s %s response: %w", request.Method, request.URL, err)
	}
	return nil
}
