package main

import (
	"context"
	"fmt"
	"github.com/mr-rambling/gator/internal/database"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	lmt := int32(2)
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s <name> <optional: limit>", cmd.Name)
	} else if len(cmd.Args) == 1 {
		parsedLimit, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("invalid limit: %v", err)
		}
		lmt = int32(parsedLimit)
	}

	posts, err := s.db.GetPostsForUser(
		context.Background(),
		database.GetPostsForUserParams{
			UserID: user.ID,
			Limit:  lmt,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to get posts for user: %v", err)
	}

	fmt.Printf("Posts for user: %s\n", user.Name)
	for _, post := range posts {
		feed, err := s.db.GetFeedByID(context.Background(), post.FeedID)
		if err != nil {
			return fmt.Errorf("failed to get feed for post: %v", err)
		}
		printPost(post, feed)
		fmt.Println("=====================================")
	}

	return nil
}

func printPost(post database.GetPostsForUserRow, feed database.Feed) {
	fmt.Printf("* Feed:          %s\n", feed.Name)
	fmt.Printf("* Title:         %s\n", post.Title)
	fmt.Printf("* URL:           %s\n", post.Url)
	fmt.Printf("* Description:   %s\n", post.Description)
	fmt.Printf("* Created:       %v\n", post.CreatedAt)
	fmt.Printf("* Updated:       %v\n", post.UpdatedAt)
	fmt.Printf("* Published:     %v\n", post.PublishedAt)
}
