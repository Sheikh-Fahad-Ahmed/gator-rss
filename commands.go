package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Sheikh-Fahad-Ahmed/gator-rss/internal/rss/api"
)

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commandsMap map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		fmt.Println("a username is required")
		os.Exit(1)
	}

	if len(cmd.arguments) > 1 {
		return errors.New("login command takes only one argument")
	}

	if _, err := s.db.GetUser(context.Background(), cmd.arguments[0]); err != nil {
		fmt.Println(cmd.arguments[0], " does not exist.")
		os.Exit(1)
	}

	if err := s.config.SetUser(cmd.arguments[0]); err != nil {
		return err
	}

	fmt.Printf("\n%s has logged in.\n", cmd.arguments[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		fmt.Println("a username is required")
		os.Exit(1)
	}
	if len(cmd.arguments) > 1 {
		return errors.New("register command takes only one argument")
	}
	err := helperCreateUser(s, cmd)
	if err != nil {
		return err
	}

	if err := s.config.SetUser(cmd.arguments[0]); err != nil {
		return err
	}
	fmt.Println(cmd.arguments[0], "has been registered as a user.")
	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.arguments) > 0 {
		return errors.New("reset command does not take any arguments")
	}
	err := s.db.DeleteAll(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("All data has been wiped.")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.arguments) > 0 {
		return errors.New("users command does not take any arguments")
	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, name := range users {
		fmt.Printf("\n* %s", name)
		if name == s.config.Current_user_name {
			fmt.Printf(" (current)")
		}
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) > 0 {
		return errors.New("agg command does not take any arguments")
	}
	feedURL := "https://www.wagslane.dev/index.xml"

	feed, err := api.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		fmt.Println("addFeed command takes 2 arguments: name url")
		os.Exit(1)
	}
	feed, err := helperCreatedFeed(s, cmd)
	if err != nil {
		return err
	}
	fmt.Println(feed.ID, feed.Name, feed.CreatedAt, feed.UpdatedAt, feed.Url, feed.UserID)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.arguments) > 0 {
		fmt.Println("feeds command does not take any arguments.")
		os.Exit(1)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for i, item := range feeds {
		fmt.Printf("\n%d. %s, %s", (i + 1), item.Name, item.Username)
	}
	return nil
}

func handlerFollow(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		fmt.Println("follow command only take one argument: url")
		os.Exit(1)
	}

	feedFollowRecord, err := helperCreateFeedFollow(s, cmd)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s	%s\n", feedFollowRecord.FeedName, feedFollowRecord.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		fmt.Println("following command does not take any argument")
		os.Exit(1)
	}

	user, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
	if err != nil {
		return err
	}
	followingFeeds, err := s.db.GetFeedFollowForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for i, item := range followingFeeds {
		fmt.Printf("\n%d. %s, %s", (i+1), item.FeedName, item.UserName)
	}
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.commandsMap[cmd.name]
	if !ok {
		return errors.New("command function not found")
	}

	if err := f(s, cmd); err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(s *state, cmd command) error) {
	if c.commandsMap == nil {
		c.commandsMap = make(map[string]func(*state, command) error)
	}
	_, exists := c.commandsMap[name]
	if exists {
		fmt.Println("command already exists")
		return
	}
	c.commandsMap[name] = f
}
