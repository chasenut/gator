package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chasenut/rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

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

func handlerUnfollow(s *state, cmd command, user database.User) error {
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

func handlerListFeedFollows(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to parse feedFollows: %w", err)
	}

	fmt.Println("Followed feeds:")
	for _, f := range feedFollows {
		fmt.Printf(" * %s\n", f.FeedName)
	}

	return nil
}

func printFeedFollow(feedFollow database.CreateFeedFollowRow) {
	fmt.Printf("Created new feed follow:\n")
	fmt.Printf(" * User:	%s\n", feedFollow.UserName)
	fmt.Printf(" * Feed: 	%s\n", feedFollow.FeedName)
}
