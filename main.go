package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pcauce/gator/commands"
	"github.com/pcauce/gator/internal/state"
	"os"
)

func main() {
	inputArgs := os.Args
	if len(inputArgs) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	appState, err := state.LoadAppState()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	commandList, err := commands.LoadCommands()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = commandList.RunCommand(&appState, commands.Command{Name: inputArgs[1], Args: inputArgs[2:]})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
