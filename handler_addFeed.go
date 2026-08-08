package main

import (
	"fmt"
	"errors"

	"time"
	"context"
	"github.com/Yishen1011/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {

	if len(cmd.Args) < 2 {
		return errors.New("You must include 2 arguments\n")
	}

	// Get current login in user
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	args := database.CreateFeedParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		Name:       cmd.Args[0],
		Url:        cmd.Args[1], 
		UserID:     user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), args)
	if err != nil {
		return err
	}

	fmt.Printf("Feed is successfully added!\nName: %s\nUrl: %s\n", feed.Name, feed.Url)
	
	return nil
}
