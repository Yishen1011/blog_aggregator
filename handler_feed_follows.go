package main

import (
	"fmt"
	"errors"

	"time"
	"context"
	"github.com/Yishen1011/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {

	if len(cmd.Args) < 1 {
		return errors.New("The argument is empty\n")
	}

	// Get current login in user
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	// Get feed from url
	feed, err := s.db.GetFeedForURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	args := database.CreateFeedFollowParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		UserID:     user.ID,
		FeedID:     feed.ID,
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), args)
	if err != nil {
		return err
	}

	fmt.Printf("Feed Follow is successfully added!\n")
	fmt.Printf("Feed: %s\nUser: %s\n", follow.FeedName, follow.UserName)
	
	return nil
}