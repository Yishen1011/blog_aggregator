package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/Yishen1011/gator/internal/database"
	"time"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {

	if len(cmd.Args) < 2 {
		return errors.New("You must include 2 arguments\n")
	}

	feedParams := database.CreateFeedParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		Name:           cmd.Args[0],
		Url:            cmd.Args[1], 
		UserID:         user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}

	followParams := database.CreateFeedFollowParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		UserID:     user.ID,
		FeedID:     feed.ID,
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return err
	}

	fmt.Printf("Feed is successfully added!\nName: %s\nUrl: %s\n", feed.Name, feed.Url)
	fmt.Printf("Feed Follow is successfully added!\n")
	fmt.Printf("Feed: %s\nUser: %s\n", follow.FeedName, follow.UserName)

	return nil
}


func handlerFeeds(s *state, cmd command) error {

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("Name: %s\n", feed.Name)
		fmt.Printf("Url: %s\n", feed.Url)

		user, err := s.db.GetUserWithID(context.Background(), feed.UserID)
			if err != nil {
			return err
		}
		fmt.Printf("Username: %s\n", user.Name)

		fmt.Println()
	}
	
	return nil
}
