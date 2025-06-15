package commands

import (
	"context"
	"errors"
	"github.com/pcauce/gator/internal/state"
)

type Command struct {
	Name string
	Args []string
}

type CommandHandler func(*state.AppState, Command) error

type CommandList map[string]CommandHandler

func LoadCommands() (CommandList, error) {
	commands := make(map[string]CommandHandler)
	commands["login"] = LoginUser
	commands["register"] = RegisterUser
	commands["users"] = GetAllUsers
	commands["reset"] = ResetUsers
	commands["agg"] = CollectFeeds
	commands["addfeed"] = AddFeed
	commands["feeds"] = ListFeeds
	commands["follow"] = FollowFeed
	commands["unfollow"] = UnfollowFeed
	commands["following"] = FollowedByUser
	commands["browse"] = BrowseSavedPosts

	return commands, nil
}

func (c *CommandList) RunCommand(s *state.AppState, cmd Command) error {
	if cmd.Name != "register" {
		exists, err := s.DBQueries.CheckUserExists(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return err
		}
		if !exists {
			return errors.New("user does not exist")
		}
	}

	err := (*c)[cmd.Name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}
