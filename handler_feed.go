package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chasenut/rss-feed-aggregator/internal/database"
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
	fmt.Printf(" * ID:			%s\n", feed.ID)
	fmt.Printf(" * CreatedAt:	%v\n", feed.CreatedAt)
	fmt.Printf(" * UpdatedAt:	%s\n", feed.UpdatedAt)
	fmt.Printf(" * Name:		%s\n", feed.Name)
	fmt.Printf(" * Url:			%s\n", feed.Url)
	fmt.Printf(" * UserID: 		%s\n", feed.UserID)
}

func printFeedPro(feed database.Feed, user database.User) {
	printFeed(feed)
	fmt.Printf(" * UserName: 	%s\n", user.Name)
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed to get specified feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), 
		database.CreateFeedFollowParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID: user.ID,
			FeedID: feed.ID,
		})

	printFeedFollow(feedFollow)

	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}

	return nil
}

func handlerDeleteFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to get feed by url: %w", err)
	}
	err = s.db.DeleteFeedFollow(context.Background(), 
		database.DeleteFeedFollowParams{
			UserID: user.ID,
			FeedID: feed.ID,
		})
	if err != nil {
		return fmt.Errorf("failed to delete feed follow: %w", err)
	}
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	followedFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to parse followedFeeds: %w", err)
	}

	for _, f := range followedFeeds {
		fmt.Println("Followed feeds:")
		fmt.Printf(" * %s\n", f.FeedName)
	}

	return nil
}

func printFeedFollow(feedFollow database.CreateFeedFollowRow) {
	fmt.Printf("Created new feed follow:\n")
	fmt.Printf(" * User:	%s\n", feedFollow.UserName)
	fmt.Printf(" * Feed: 	%s\n", feedFollow.FeedName)
}
