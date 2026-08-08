package main

import (
	"fmt"
	"context"
)

func handlerFollowing(s *state, cmd command) error {

	// Get current login in user
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	followings, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, follow := range followings {
		fmt.Printf("%s is following %s\n", follow.UserName, follow.FeedName)
	}
	
	return nil
}