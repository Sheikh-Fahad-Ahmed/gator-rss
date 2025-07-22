package main

import (
	"fmt"
	"os"

	"github.com/Sheikh-Fahad-Ahmed/gator-rss/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing arguments")
		os.Exit(1)
	}

	cmd := command{
		name: os.Args[1],
		arguments: os.Args[2:],
	}
	
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error Read function")
	}

	s := &state{
		config: &cfg,
	}

	cmds := &commands{
		commandsMap: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	if err := cmds.run(s,cmd); err != nil {
		fmt.Println(err)
	}
}