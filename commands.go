package main

import (
	"errors"
	"fmt"
	"os"
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
	if err := s.config.SetUser(cmd.arguments[0]); err != nil {
		return err
	}

	fmt.Println("A User has been set..")
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
