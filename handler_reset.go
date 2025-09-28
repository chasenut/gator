package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background());
	if err != nil {
		return fmt.Errorf("failed to reset: %w", err)
	}
	fmt.Println("Reset successfully")
	return nil
}

