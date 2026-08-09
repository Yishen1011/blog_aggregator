package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/Yishen1011/blog_aggregator/internal/database"
	"time"
)

func handlerFollow(s *state, cmd command, user database.User) error {

	if len(cmd.Args) < 1 {
		return errors.New("The argument is empty\n")
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

func handlerListFeedFollows(s *state, cmd command, user database.User) error {

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, follow := range feedFollows {
		fmt.Printf("* %s\n", follow.FeedName)
	}
	
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.Args) < 1 {
		return errors.New("The argument is empty\n")
	}

	// Get feed from url
	feed, err := s.db.GetFeedForURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	unfollowParams := database.DeleteFeedFollowsParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.db.DeleteFeedFollows(context.Background(), unfollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("Feed Follow is successfully deleted!\n")
	fmt.Printf("Feed: %s\nUser: %s\n", feed.Name, user.Name)
	
	return nil
}
