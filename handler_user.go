package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Not enough positional arguments provided")
	}
	name := cmd.args[0]
	err := s.cfg.SetUser(name)
	if err != nil {
		return err
	}
	fmt.Printf("User %s has been registered!\n", name)
	return nil
}
