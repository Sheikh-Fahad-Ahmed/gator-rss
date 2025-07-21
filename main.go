package main

import (
	"fmt"

	"github.com/Sheikh-Fahad-Ahmed/gator-rss/internal/config"
)

func main() {
	fmt.Println("Hello World!")
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error Read function")
	}
	fmt.Println(cfg)
}