package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/chasenut/rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Not enough positional arguments provided")
	}
	name := cmd.args[0]
	if _, err := s.db.GetUser(context.Background(), name); err != nil {
		log.Fatal("Cannot login, no such user exists")
	}
	err := s.cfg.SetUser(name)
	if err != nil {
		return err
	}
	fmt.Println("Successfully changed the user")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Not enough positional arguments provided")
	}
	name := cmd.args[0]
	if _, err := s.db.GetUser(context.Background(), name); err == nil {
		log.Fatal("User already exists")
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

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background());
	if err != nil {
		log.Fatal("Failed to reset users")
	}
	fmt.Println("Successfully reset users")
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.ListUsers(context.Background())
	if err != nil {
		log.Fatal("Failed to access users data")
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
	fmt.Printf(" * ID: 		%v\n", user.ID)
	fmt.Printf(" * Name: 	%v\n", user.Name)
}
