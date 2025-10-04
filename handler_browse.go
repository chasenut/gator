package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/chasenut/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) > 0 {
		if i, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = i
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), 
		database.GetPostsForUserParams{
			UserID: user.ID,
			Limit: int32(limit),
		})
	if err != nil {
		return fmt.Errorf("failed to find posts for user: %w", err)
	}
	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("	%v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Printf("==================================\n")
	}
	return nil
}
