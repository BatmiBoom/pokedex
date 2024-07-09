package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/BatmiBoom/pokedex/cmd/commands"
	"github.com/BatmiBoom/pokedex/cmd/config"
)

func StartRepl(cfg *config.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanned := scanner.Scan()
		if scanned == false {
			fmt.Println("ERROR: There was an error reaing the command")
		}

		command, err := commands.GetCommand(scanner.Text())
		if err != nil {
			fmt.Printf("%v", err)
		}

		err = command.Callback(cfg)
		if err != nil {
			fmt.Println("There was something wrong with the command ", command.Name)
		}
	}
}
