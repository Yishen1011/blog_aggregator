package main

import (
	"context"

	"github.com/Yishen1011/blog_aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	return func (s *state, cmd command) error {
		// Get current login in user
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}

}