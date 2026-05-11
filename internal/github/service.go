package github

func FindMassFollowers(blockedUsersCh <-chan string, config *Config) {
	userFollowers := GetGitHubUsersPage(config.UserName, UserFollowersURI)
	var wg sync.WaitGroup
	for _, userFollower := range *userFollowers {
		wg.Go(func() {
			userFollowing := GetGitHubUsersPage(userFollower.Login, UserFollowingURI)
			noOfFollowing := len(*userFollowing)
			if noOfFollowing > config.FollowersThresholdToBlock {
				go BlockMassFollowers(blockedUsersCh, userFollower.Login)
			}
		})
	}
	wg.Wait()
}

func BlockMassFollowers(blockedUsersCh <-chan string, userToBlock string) {
	log.Printf("Blocking User: %s", userToBlock)
}
