package main

import (
	"context"
	"fmt"

	"github.com/chasenut/gator/internal/database"
)

type command struct {
	Name 	string
	Args 	[]string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
} 

func (c *commands) run(s *state, cmd command) error {
	command, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("No such command exists: %v\n", cmd.Name)
	}
	return command(s, cmd)
}

func (c *commands) register(name string, f func(s *state, cmd command) error) {
	c.registeredCommands[name] = f
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("failed to get current user: %w", err)
		}
		return handler(s, cmd, user)
	}
}
