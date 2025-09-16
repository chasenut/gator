package main

import (
	"fmt"
)

type command struct {
	name 	string
	args 	[]string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
} 

func (c *commands) run(s *state, cmd command) error {
	if c == nil {
		return fmt.Errorf("State in NULL\n")
	}
	command, ok := c.registeredCommands[cmd.name]
	if !ok {
		return fmt.Errorf("No such command exists: %v\n", cmd.name)
	}
	err := command(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(s *state, cmd command) error) {
	c.registeredCommands[name] = f
}
