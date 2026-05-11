package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/prdai/github-follower-prune/internal/types"
)

const ConfigFilePath string = "config.json"
const EnvFilePath string = ".env"

func main() {
	err := godotenv.Load(EnvFilePath)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(0)
	}
	buffer, err := os.ReadFile(ConfigFilePath)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(0)
	}
	var config types.Config
	if err := json.Unmarshal(buffer, &config); err != nil {
		log.Fatal(err.Error())
		os.Exit(0)
	}
	blockedUsersCh := make(chan string)
	FindMassFollowers(blockedUsersCh, &config)
}
