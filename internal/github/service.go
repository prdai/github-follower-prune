package github

import (
	"fmt"
	"sync"

	"github.com/prdai/github-follower-prune/internal/types"
)

const maxConcurrentProfiles = 5

type PruneResult struct {
	Blocked []string
	Failed  map[string]error
}

func (g *githubClient) PruneMassFollowers(config *types.Config) (*PruneResult, error) {
	followers, err := g.GetGitHubFollowers(config.UserName)
	if err != nil {
		return nil, fmt.Errorf("listing followers of %s: %w", config.UserName, err)
	}
	result := &PruneResult{Failed: map[string]error{}}
	var (
		wg     sync.WaitGroup
		mu     sync.Mutex
		tokens = make(chan struct{}, maxConcurrentProfiles)
	)
	for _, follower := range followers {
		tokens <- struct{}{}
		wg.Go(func() {
			defer func() { <-tokens }()
			blocked, err := g.pruneMassFollower(follower.Login, config)
			mu.Lock()
			defer mu.Unlock()
			if err != nil {
				result.Failed[follower.Login] = err
				return
			}
			if blocked {
				result.Blocked = append(result.Blocked, follower.Login)
			}
		})
	}
	wg.Wait()
	return result, nil
}

func (g *githubClient) pruneMassFollower(login string, config *types.Config) (bool, error) {
	user, err := g.GetGitHubUser(login)
	if err != nil {
		return false, fmt.Errorf("fetching profile: %w", err)
	}
	if user.Followers <= config.FollowersThresholdToBlock || user.Following <= config.FollowingThresholdToBlock {
		return false, nil
	}
	if err := g.BlockGithubUser(login, blockUserURI); err != nil {
		return false, fmt.Errorf("blocking: %w", err)
	}
	if err := g.BlockGithubUser(login, unblockUserURI); err != nil {
		return false, fmt.Errorf("unblocking: %w", err)
	}
	return true, nil
}
