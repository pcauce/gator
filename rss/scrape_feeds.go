package rss

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
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
		publishedAt, err := time.Parse(item.PubDate, "2013-Feb-03")
		if err != nil {
			return err
		}
		_, err = s.DBQueries.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: sql.NullTime{
				Time:  publishedAt,
				Valid: true,
			},
			FeedID: uuid.NullUUID{
				UUID:  feedToFetch.ID,
				Valid: true,
			},
		})
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code != "23505" {
				fmt.Println("Couldnt insert post:", item.Title, "//", pqErr.Code, "-", pqErr.Message)
			}
		}
	}

	return nil
}
