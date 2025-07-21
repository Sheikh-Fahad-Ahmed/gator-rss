package main

import (
	"errors"
	"fmt"
)

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commandsMap map[string]func(*State, command) error
}

func handlerLogin(s *State, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("no command arguments")
	}
	if len(cmd.arguments) > 1 {
		return errors.New("login command takes only one argument")
	}
	if err := s.Config.SetUser(cmd.arguments[0]); err != nil {
		return err
	}

	fmt.Println("A User has been set..")
	return nil
}

func (c *commands) run(s *State, cmd command) error {
	f, ok := c.commandsMap[cmd.name]
	if !ok {
		return errors.New("command function not found")
	}

	if err := f(s, cmd); err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(s *State, cmd command) error) {
	if c.commandsMap == nil {
		c.commandsMap = make(map[string]func(*State, command) error)
	}
	_, exists := c.commandsMap[name]
	if exists {
		fmt.Println("command already exists")
		return
	}
	c.commandsMap[name] = f
}
