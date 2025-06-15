package commands

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/pcauce/gator/internal/state"
)

func UnfollowFeed(s *state.AppState, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("one argument expected for register command: url")
	}
	userID, err := s.DBQueries.GetUserID(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}

	err = s.DBQueries.Unfollow(context.Background(), uuid.NullUUID{UUID: userID, Valid: true})
	if err != nil {
		return err
	}

	return nil
}
