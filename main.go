package main

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

func main() {

}
