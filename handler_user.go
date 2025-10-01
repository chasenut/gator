package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/chasenut/rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command, ) error {
	if len(cmd.Args) == 0 {
		return errors.New("Not enough positional arguments provided")
	}
	name := cmd.Args[0]
	if _, err := s.db.GetUser(context.Background(), name); err != nil {
		return fmt.Errorf("no such user found, please register: %w", err)
	}
	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Printf("Successfully changed to the user: %s\n", name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <Name>", cmd.Name)
	}
	name := cmd.Args[0]
	if _, err := s.db.GetUser(context.Background(), name); err == nil {
		return fmt.Errorf("user already exists: %w", err)
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: name,
	})
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Println("User created successfully")
	printUser(user)
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list users: %w", err)
	}
	if len(users) == 0 {
		fmt.Println("List of users is empty")
		return nil
	}
	current := s.cfg.CurrentUserName

	for _, u := range users {
		fmt.Printf("* %s", u.Name)
		if u.Name == current {
			fmt.Printf(" (current)")
		}
		fmt.Printf("\n")
	}
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:		%v\n", user.ID)
	fmt.Printf(" * Name:	%v\n", user.Name)
}
