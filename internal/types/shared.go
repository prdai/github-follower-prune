package types

type Config struct {
	FollowersThresholdToBlock int    `json:"FOLLOWERS_THRESHOLD_TO_BLOCK"`
	FollowingThresholdToBlock int    `json:"FOLLOWING_THRESHOLD_TO_BLOCK"`
	UserName                  string `json:"USER_NAME"`
}
