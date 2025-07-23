package main

import (
	"github.com/Sheikh-Fahad-Ahmed/gator-rss/internal/config"
	"github.com/Sheikh-Fahad-Ahmed/gator-rss/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}
