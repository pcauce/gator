package rss

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pcauce/gator/internal/database"
	"github.com/pcauce/gator/internal/state"
	"time"
)

func ScrapeFeeds(s *state.AppState) error {
	feedToFetch, err := s.DBQueries.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.DBQueries.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:        feedToFetch.ID,
		UpdatedAt: time.Now(),
		LastFetched: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		return err
	}

	feedData, err := FetchFeed(context.Background(), feedToFetch.Url)
	if err != nil {
		return err
	}

	for _, item := range feedData.Channel.Items {
		fmt.Println(item.Title)
	}

	return nil
}
