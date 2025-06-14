package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/pcauce/gator/internal/state"
)

func LoginUser(s *state.AppState, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("only one argument expected for register command")
	}
	userName := cmd.Args[0]

	exists, err := s.DBQueries.CheckUserExists(context.Background(), userName)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("user does not exist")
	}

	err = s.Config.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Println("Username set to: ", s.Config.CurrentUserName)
	return nil
}
