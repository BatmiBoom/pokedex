package main

import (
	"github.com/BatmiBoom/pokedex/cmd/config"
	"github.com/BatmiBoom/pokedex/cmd/repl"
)

func main() {
	cfg := config.GetConfig()

	repl.StartRepl(&cfg)
}
