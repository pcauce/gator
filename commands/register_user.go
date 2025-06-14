package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/pcauce/gator/internal/database"
	"github.com/pcauce/gator/internal/state"
	"time"
)

func RegisterUser(s *state.AppState, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("only one argument expected for register command")
	}

	user, err := s.DBQueries.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		return err
	}

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return err
	}
	err = s.Config.WriteCfgFile()

	fmt.Println("Username set to: ", user.Name, "\n", "User details: \n", user)
	return nil
}
