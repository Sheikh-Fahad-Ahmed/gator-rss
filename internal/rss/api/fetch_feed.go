package api

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func fetchFeed(context context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(context, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rssFeed RSSFeed

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := xml.Unmarshal(body, &rssFeed); err != nil {
		return nil, err
	}

	html.UnescapeString(rssFeed.Channel.Title)
	html.UnescapeString(rssFeed.Channel.Description)
	
	for _, item := range(rssFeed.Channel.Item) {
		html.UnescapeString(item.Title)
		html.UnescapeString(item.Description)
	}


	return &rssFeed, nil
}
