package main

import (
	"fmt"
	"context"
)

func handlerReset(s *state, cmd command) error {

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("All users have been deleted from the database\n")
	return nil
}