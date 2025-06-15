package commands

import (
	"errors"
	"fmt"
	"github.com/pcauce/gator/internal/state"
	"github.com/pcauce/gator/rss"
	"os"
	"os/exec"
	"time"
)

func CollectFeeds(s *state.AppState, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("only one argument expected for agg command: time between requests")
	}

	c := exec.Command("clear")
	c.Stdout = os.Stdout

	refreshCount := 0

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		c.Run()
		refreshCount++
		fmt.Println("Refresh number:", refreshCount)

		err := rss.ScrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}
