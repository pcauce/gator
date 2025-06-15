package commands

import (
	"context"
	"fmt"
	"github.com/pcauce/gator/internal/state"
)

func FollowedByUser(s *state.AppState, cmd Command) error {
	userExists, err := s.DBQueries.CheckUserExists(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}
	if !userExists {
		return fmt.Errorf("user %s does not exist", s.Config.CurrentUserName)
	}

	feedsFollowed, err := s.DBQueries.GetFeedsFollowedByUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}

	fmt.Println(feedsFollowed)
	return nil
}
