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

func FollowFeed(s *state.AppState, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("only one argument expected for follow command")
	}

	userID, err := s.DBQueries.GetUserID(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}
	feed, err := s.DBQueries.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	_, err = s.DBQueries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    uuid.NullUUID{UUID: userID, Valid: true},
		FeedID:    uuid.NullUUID{UUID: feed.ID, Valid: true},
	})
	if err != nil {
		return err
	}

	fmt.Println("Feed added:", feed.Name, "// User:", s.Config.CurrentUserName)
	return nil
}
