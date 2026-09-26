package github

type GithubUser struct {
	Login     string `json:"login"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
}

type GithubFollower struct {
	Login string `json:"login"`
}

type GithubUserURI string

type GithubBlockURI struct {
	URI    string
	Method string
}
