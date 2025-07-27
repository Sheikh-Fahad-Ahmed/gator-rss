package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Sheikh-Fahad-Ahmed/gator-rss/internal/database"
	"github.com/google/uuid"
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

func helperCreatedFeed(s *state, cmd command) (*database.Feed, error) {
	id := uuid.New()
	currentTime := time.Now()
	feedName := cmd.arguments[0]
	feedURL := cmd.arguments[1]
	username := s.config.Current_user_name
	userInfo, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return nil, err
	}

	params := database.CreateFeedParams{
		ID: id,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name: feedName,
		Url: feedURL,
		UserID: userInfo.ID,
	}

	feed,err := s.db.CreateFeed(context.Background(),params)
	if err != nil {
		return nil, err
	}
	return &feed, nil
}

func helperCreateFeedFollow(s *state, cmd command) (*database.CreateFeedFollowRow, error) {
	id := uuid.New()
	currentTime := time.Now()
	user_id, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
	if err != nil {
		return nil, err
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.arguments[0])
	if err != nil {
		return nil, err
	}
	params := database.CreateFeedFollowParams{
		ID: id,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID: user_id.ID,
		FeedID: feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(),params)
	if err != nil {
		return nil, err
	}
	return &feedFollow, nil
}