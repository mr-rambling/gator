package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/mr-rambling/gator/internal/database"
)

func handlerAggregator(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name> <time_interval>", cmd.Name)
	}

	time_between_reqs := cmd.Args[0]
	timeBetweenRequests, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting next feed to fetch: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	fmt.Printf("Saved feed to database: %s\n", feed.Channel.Title)
	for _, item := range feed.Channel.Item {
		tm := parseTime(item)

		_, err = s.db.CreatePost(
			context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Title:       item.Title,
				Url:         item.Link,
				Description: item.Description,
				PublishedAt: tm,
				FeedID:      nextFeed.ID,
			},
		)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				// 23505 is the error code for unique_violation in PostgreSQL
				if pqErr.Code == "23505" {
					// This is a duplicate URL, just ignore it
					continue
				}
			}
			// If it's a different error, log it
			fmt.Printf("error creating post: %v", err)
			continue
		}

	}
	fmt.Println("=====================================")
	return nil
}

func parseTime(item RSSItem) time.Time {
	tm, err := time.Parse(time.RFC1123Z, item.PubDate)
	if err != nil {
		tm, err = time.Parse(time.RFC3339, item.PubDate)
	}
	if err != nil {
		log.Printf("error parsing publication date: %v\n", err)
		tm = time.Now()
	}
	return tm
}
