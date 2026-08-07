package main

import (
	"fmt"
	"errors"

	"time"
	"context"
	"github.com/Yishen1011/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {

	if len(cmd.Args) < 1 {
		return errors.New("The argument is empty\n")
	}

	args := database.CreateUserParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		Name:       cmd.Args[0],
	}

	user, err := s.db.CreateUser(context.Background(), args)
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("%s has been register into the database\n", s.cfg.CurrentUserName)
	return nil
}