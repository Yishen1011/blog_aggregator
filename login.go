package main

import (
	"fmt"
	"errors"

	"context"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.Args) < 1 {
		return errors.New("The argument is empty\n")
	}

	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(cmd.Args[0]); err != nil {
		return err
	}

	fmt.Printf("%s has been set into the config\n", s.cfg.CurrentUserName)
	return nil
}