package main

import (
	"fmt"
	"errors"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.Args) < 1 {
		return errors.New("The argument is empty\n")
	}

	if err := s.cfg.SetUser(cmd.Args[0]); err != nil {
		return err
	}

	fmt.Printf("%s has been set into the config\n", s.cfg.CurrentUserName)
	return nil
}