package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"html"
	"log"
	"os"
	"time"

	"github.com/Sheikh-Fahad-Ahmed/gator-rss/internal/database"
	"github.com/Sheikh-Fahad-Ahmed/gator-rss/internal/rss/api"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func helperCreateUser(s *state, cmd command) error {
	id := uuid.New()
	currentTime := time.Now()
	name := cmd.arguments[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		fmt.Printf("%s already exists.\n", name)
		os.Exit(1)
	}

	params := database.CreateUserParams{
		ID:        id,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      name,
	}

	_, err = s.db.CreateUser(context.Background(), params)
	if err != nil {
		return errors.New("error when creating user")
	}
	return nil
}

func helperCreatedFeed(s *state, cmd command, user database.User) (*database.Feed, error) {
	id := uuid.New()
	currentTime := time.Now()
	feedName := cmd.arguments[0]
	feedURL := cmd.arguments[1]

	params := database.CreateFeedParams{
		ID: id,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name: feedName,
		Url: feedURL,
		UserID: user.ID,
	}

	feed,err := s.db.CreateFeed(context.Background(),params)
	if err != nil {
		return nil, err
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID: user.ID,
		FeedID: feed.ID,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return nil, err
	}

	return &feed, nil
}

func helperCreateFeedFollow(s *state, cmd command, user database.User) (*database.CreateFeedFollowRow, error) {
	id := uuid.New()
	currentTime := time.Now()
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.arguments[0])
	if err != nil {
		return nil, err
	}
	params := database.CreateFeedFollowParams{
		ID: id,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID: user.ID,
		FeedID: feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(),params)
	if err != nil {
		return nil, err
	}
	return &feedFollow, nil
}

func helperDeleteFeedFollow(s *state, cmd command, user database.User) error {
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}
	params := database.DeleteFeedFollowParams {
		FeedID: feed.ID,
		UserID: user.ID,
	}
	
	err = s.db.DeleteFeedFollow(context.Background(), params)
	return err
}

func helperCreatePost(s *state, feedId uuid.UUID, post *api.RSSItem) (*database.Post, error) {
	description := sql.NullString{
		String: html.UnescapeString(post.Description),
		Valid: post.Description != "",
	}
	layout := time.RFC1123

	publishedAt, err := time.Parse(layout, post.PubDate)
	if err != nil {
		log.Println("unable to parse published_at time")
		return nil, err
	}

	params := database.CreatePostParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title: html.UnescapeString(post.Title),
		Url: post.Link,
		Description: description,
		PublishedAt: publishedAt,
		FeedID: feedId,
	}
	resp, err := s.db.CreatePost(context.Background(),params)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505"{
			return nil, nil
		}
		log.Println("error creating a post:", err)
		return nil, err
	}
	return &resp, nil
}