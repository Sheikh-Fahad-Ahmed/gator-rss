package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Sheikh-Fahad-Ahmed/gator-rss/internal/rss/api"
)


func FetchAndProcessFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	log.Println("Found a feed to fetch.")
	if err = s.db.MarkFeedFetched(context.Background(), feed.ID); err != nil {
		return err
	}

	feedInfo, err := api.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for i, item := range feedInfo.Channel.Item {
		fmt.Printf("\n%d. %s",(i + 1), item.Title)
	}
	log.Printf("\n feed %s collected and %d posts found.", feed.Name, len(feedInfo.Channel.Item))
	return nil
}