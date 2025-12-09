package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mr-rambling/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	ffRow, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed followed successfully:")
	printFeedFollow(ffRow.UserName, ffRow.FeedName)
	fmt.Println("=====================================")
	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	follows, err := s.db.GetFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feeds: %w", err)
	}

	if len(follows) == 0 {
		fmt.Println("No follows found.")
		return nil
	}

	printFollows(follows, user)
	return nil
}

func printFollows(follows []database.GetFollowsForUserRow, user database.User) {
	fmt.Printf("Found %d follows:\n", len(follows))
	fmt.Printf("* User:          %s\n", user.Name)
	for _, feed := range follows {
		fmt.Printf("* Feed:          %s\n", feed.FeedName)
	}
	fmt.Println("=====================================")
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	err = s.db.DeleteFeedFollow(
		context.Background(),
		database.DeleteFeedFollowParams{
			UserID: user.ID,
			FeedID: feed.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("couldn't delete feed follow: %w", err)
	}

	fmt.Println("Feed unfollowed successfully:")
	printFeedFollow(user.Name, feed.Name)
	fmt.Println("=====================================")
	return nil
}
