package commands

import (
	"github.com/pcauce/gator/internal/state"
)

type Command struct {
	Name string
	Args []string
}

type CommandHandler func(*state.AppState, Command) error

type CommandList map[string]CommandHandler

func (c *CommandList) RunCommand(s *state.AppState, cmd Command) error {
	err := (*c)[cmd.Name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func LoadCommands() (CommandList, error) {
	commands := make(map[string]CommandHandler)
	commands["login"] = LoginUser
	commands["register"] = RegisterUser
	commands["users"] = GetAllUsers
	commands["reset"] = ResetUsers
	commands["agg"] = PrintFeed
	commands["addfeed"] = AddFeed

	return commands, nil
}
