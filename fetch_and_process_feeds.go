package main

import (
	"context"
	"fmt"

	"github.com/Sheikh-Fahad-Ahmed/gator-rss/internal/rss/api"
)


func FetchAndProcessFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	if err = s.db.MarkFeedFetched(context.Background(), feed.ID); err != nil {
		return err
	}

	feeds, err := api.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for i, item := range(feeds.Channel.Item) {
		fmt.Printf("\n%d. %s",(i + 1), item.Title)
	}
	return nil
}