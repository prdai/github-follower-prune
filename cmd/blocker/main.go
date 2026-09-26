package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/prdai/github-follower-prune/internal/github"
	"github.com/prdai/github-follower-prune/internal/types"
)

const (
	configFilePath = "config.json"
	envFilePath    = ".env"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if err := godotenv.Load(envFilePath); err != nil {
		return fmt.Errorf("loading %s: run `make init` first: %w", envFilePath, err)
	}
	config, err := loadConfig(configFilePath)
	if err != nil {
		return err
	}
	client, err := github.NewGithubClient(config)
	if err != nil {
		return err
	}
	result, err := client.PruneMassFollowers(config)
	if err != nil {
		return err
	}
	for _, login := range result.Blocked {
		fmt.Printf("pruned %s\n", login)
	}
	for login, err := range result.Failed {
		fmt.Fprintf(os.Stderr, "FAILED %s: %v\n", login, err)
	}
	fmt.Printf("done: %d pruned, %d failed\n", len(result.Blocked), len(result.Failed))
	return nil
}

func loadConfig(path string) (*types.Config, error) {
	buffer, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	var config types.Config
	if err := json.Unmarshal(buffer, &config); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}
	if config.UserName == "" {
		return nil, fmt.Errorf("%s: USER_NAME is required", path)
	}
	return &config, nil
}
