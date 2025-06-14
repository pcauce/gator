package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/pcauce/gator/internal/state"
)

func GetAllUsers(s *state.AppState, cmd Command) error {
	if len(cmd.Args) > 0 {
		return errors.New("no arguments expected for users command")
	}

	users, err := s.DBQueries.GetAllUsers(context.Background())
	if err != nil {
		return errors.New("couldn't reset database")
	}

	for _, user := range users {
		if user.Name == s.Config.CurrentUserName {
			fmt.Println(user.Name, "(current)")
			continue
		}
		fmt.Println(user.Name)
	}

	return nil
}
