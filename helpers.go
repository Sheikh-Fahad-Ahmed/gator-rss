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