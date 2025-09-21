package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chasenut/rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAggregator(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to fetch the feed: %w", err)
	}
	fmt.Printf("%+v\n", feed)

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	dbFeed, err := s.db.CreateFeed(context.Background(), 
		database.CreateFeedParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name: name,
			Url: url,
			UserID: currentUser.ID,
		})
	if err != nil {
		return fmt.Errorf("failed to create new feed: %w", err)
	}
	printFeed(dbFeed)

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("%s:\n", feed.Name)
	fmt.Printf(" * ID: %s\n", feed.ID)
	fmt.Printf(" * CreatedAt: %v\n", feed.CreatedAt)
	fmt.Printf(" * UpdatedAt: %s\n", feed.UpdatedAt)
	fmt.Printf(" * Name: %s\n", feed.Name)
	fmt.Printf(" * Url: %s\n", feed.Url)
	fmt.Printf(" * UserID: %s\n", feed.UserID)
}
