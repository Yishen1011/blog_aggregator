package main

import (
	"fmt"
	"context"
)

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