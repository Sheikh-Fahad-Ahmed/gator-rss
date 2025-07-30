package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Sheikh-Fahad-Ahmed/gator-rss/internal/config"
	"github.com/Sheikh-Fahad-Ahmed/gator-rss/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing arguments")
		os.Exit(1)
	}

	cmd := command{
		name:      os.Args[1],
		arguments: os.Args[2:],
	}

	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error Read function")
	}

	db, err := sql.Open("postgres", cfg.Db_url)
	if err != nil {
		fmt.Println("Error while opening database: ", err)
	}

	dbQueries := database.New(db)

	s := &state{
		db:     dbQueries,
		config: &cfg,
	}

	cmds := &commands{
		commandsMap: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))
	if err := cmds.run(s, cmd); err != nil {
		fmt.Println(err)
	}
}
