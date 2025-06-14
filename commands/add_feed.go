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

func AddFeed(s *state.AppState, cmd Command) error {
	if len(cmd.Args) != 2 {
		return errors.New("only one argument expected for addfeed command")
	}

	userID, err := s.DBQueries.GetUserID(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}
	addedFeed, err := s.DBQueries.StoreFeed(context.Background(), database.StoreFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    uuid.NullUUID{UUID: userID, Valid: true},
	})
	if err != nil {
		return err
	}

	fmt.Println(addedFeed)
	return nil
}
