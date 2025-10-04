package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chasenut/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), 
		database.CreateFeedParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name: name,
			Url: url,
			UserID: user.ID,
		})
	if err != nil {
		return fmt.Errorf("failed to create new feed: %w", err)
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), 
		database.CreateFeedFollowParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID: user.ID,
			FeedID: feed.ID,
		})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}
	printFeed(feed)
	printFeedFollow(feedFollow)

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to list feeds: %w", err)
	}
	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("failed to fetch user data: %w", err)
		}
		printFeedPro(feed, user)
	}

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("%s:\n", feed.Name)
	fmt.Printf(" * ID:				%s\n", feed.ID)
	fmt.Printf(" * Created:			%v\n", feed.CreatedAt)
	fmt.Printf(" * Updated:			%s\n", feed.UpdatedAt)
	fmt.Printf(" * Name:			%s\n", feed.Name)
	fmt.Printf(" * Url:				%s\n", feed.Url)
	fmt.Printf(" * UserID: 			%s\n", feed.UserID)
	fmt.Printf(" * LastFetchedAt: 	%s\n", feed.LastFetchedAt.Time)
}

func printFeedPro(feed database.Feed, user database.User) {
	printFeed(feed)
	fmt.Printf(" * UserName: 	%s\n", user.Name)
}
