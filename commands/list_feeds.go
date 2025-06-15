package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/pcauce/gator/internal/state"
)

func ListFeeds(s *state.AppState, cmd Command) error {
	if len(cmd.Args) > 0 {
		return errors.New("no arguments expected for feeds command")
	}

	feeds, err := s.DBQueries.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		creator, err := s.DBQueries.GetUserName(context.Background(), feed.UserID.UUID)
		if err != nil {
			return err
		}

		fmt.Println("Feed:", feed.Name, "// URL:", feed.Url, "// Created by:", creator)
	}
	return nil
}
