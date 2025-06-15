package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/pcauce/gator/internal/database"
	"github.com/pcauce/gator/internal/state"
	"strconv"
)

func BrowseSavedPosts(s *state.AppState, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("only one argument expected for browse command: limit")
	}

	userID, err := s.DBQueries.GetUserID(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}

	limit, err := strconv.Atoi(cmd.Args[0])
	if err != nil {
		return err
	}
	switch {
	case limit < 1:
		return errors.New("limit must be greater than 0")
	case limit > 100:
		return errors.New("limit must be less than 100")
	}

	posts, err := s.DBQueries.GetUserSavedPosts(context.Background(), database.GetUserSavedPostsParams{
		UserID: uuid.NullUUID{
			UUID:  userID,
			Valid: true,
		},
		Limit: int32(limit),
	})
	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.Description)
		fmt.Println(post.PublishedAt)
		fmt.Println(post.Url)
		fmt.Println("------------")
	}

	return nil
}
