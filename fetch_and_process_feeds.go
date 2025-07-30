package main

import (
	"context"
	"log"


	"github.com/Sheikh-Fahad-Ahmed/gator-rss/internal/rss/api"
)


func FetchAndProcessFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	log.Printf("\nfound a feed to fetch: %s", feed.Name)
	if err = s.db.MarkFeedFetched(context.Background(), feed.ID); err != nil {
		return err
	}

	feedInfo, err := api.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range feedInfo.Channel.Item {
		_, err := helperCreatePost(s, feed.ID, &item)
		if err != nil {
			return err
		}
	}
	log.Printf("\n feed %s collected and %d posts found.", feed.Name, len(feedInfo.Channel.Item))
	return nil
}