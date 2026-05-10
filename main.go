package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type GithubUser struct {
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Login     string `json:"login"`
	ID        int64 `json:"id"`
	NodeID    string `json:"node_id"`
	AvatarURL string `json:"avatar_url"`
	GravatarID string `json:"gravatar_id,omitempty"`
	URL string `json:"url"`
	HTMLURL string `json:"html_url"`
	FollowersURL string `json:"followers_url"`
	FollowingURL string `json:"following_url"`
	GistsURL string `json:"gists_url"`
	StarredURL string `json:"starred_url"`
	SubscriptionsURL string `json:"subscriptions_url"`
	OrganizationsURL string `json:"organizations_url"`
	ReposURL string `json:"repos_url"`
	EventsURL string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type string `json:"type"`
	SiteAdmin bool `json:"site_admin"`
	StarredAt string `json:"starred_at"`
	UserViewType string `json:"user_view_type"`
}

type Config struct {
	FollowersThresholdToBlock int64 `json:"FOLLOWERS_THRESHOLD_TO_BLOCK"`
	UserName string `json:"USER_NAME"`
}

const UserFollowersURI = "https://api.github.com/users/%s/followers";
const ConfigFilePath = "./config.json";
const EnvFilePath = "./.env";

func main() {
	err := godotenv.Load(EnvFilePath);
	if err != nil {
		log.Fatal(err.Error());
		os.Exit(0);
	}
	buffer, err := os.ReadFile(ConfigFilePath)
	if err != nil {
		log.Fatal(err.Error());
		os.Exit(0);
	}
	var config Config;
	if err := json.Unmarshal(buffer, &config); err != nil {
		log.Fatal(err.Error());
		os.Exit(0);
	}
	fmt.Print("%v", config);
}






