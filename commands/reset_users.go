package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/pcauce/gator/internal/state"
)

func ResetUsers(s *state.AppState, cmd Command) error {
	if len(cmd.Args) > 0 {
		return errors.New("no arguments expected for reset command")
	}

	err := s.DBQueries.DeleteAllUsers(context.Background())
	if err != nil {
		return errors.New("couldnt reset database")
	}

	fmt.Println("Database reset successfully")
	return nil
}
