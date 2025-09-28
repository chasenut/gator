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

	feed, err := s.db.CreateFeed(context.Background(), 
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
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), 
		database.CreateFeedFollowParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID: currentUser.ID,
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
		printFeedPro(feed, user.Name)
	}

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("%s:\n", feed.Name)
	fmt.Printf(" * ID:			%s\n", feed.ID)
	fmt.Printf(" * CreatedAt:	%v\n", feed.CreatedAt)
	fmt.Printf(" * UpdatedAt:	%s\n", feed.UpdatedAt)
	fmt.Printf(" * Name:		%s\n", feed.Name)
	fmt.Printf(" * Url:			%s\n", feed.Url)
	fmt.Printf(" * UserID: 		%s\n", feed.UserID)
}

func printFeedPro(feed database.Feed, username string) {
	printFeed(feed)
	fmt.Printf(" * UserName: 	%s\n", username)
}

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed to get specified feed: %w", err)
	}

	current, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), 
		database.CreateFeedFollowParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID: current.ID,
			FeedID: feed.ID,
		})

	printFeedFollow(feedFollow)

	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}

	return nil
}

func printFeedFollow(feedFollow database.CreateFeedFollowRow) {
	fmt.Printf("Created new feed follow:\n")
	fmt.Printf(" * User:	%s\n", feedFollow.UserName)
	fmt.Printf(" * Feed: 	%s\n", feedFollow.FeedName)
}

func handlerFollowing(s *state, cmd command) error {
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	followedFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)

	for _, f := range followedFeeds {
		fmt.Println("Followed feeds:")
		fmt.Printf(" * %s\n", f.FeedName)
	}

	return nil
}

