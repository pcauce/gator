package commands

import (
	"context"
	"fmt"
	"github.com/pcauce/gator/internal/state"
	"github.com/pcauce/gator/rss"
)

func PrintFeed(s *state.AppState, cmd Command) error {
	feedStruct, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(feedStruct)
	return nil
}
