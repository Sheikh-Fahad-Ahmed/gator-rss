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
		fmt.Println("User already exists...")
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
