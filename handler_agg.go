package main

import (
	"fmt"
	"context"
	"errors"
	"time"
	"database/sql"
	"github.com/Yishen1011/blog_aggregator/internal/database"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.Args) < 1 {
		return errors.New("The argument is empty\n")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenRequests)
	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func printRSS(rss *RSSFeed) {
	fmt.Printf("%s\n", rss.Channel.Title)
	fmt.Printf("%s\n", rss.Channel.Description)

	for _, item := range rss.Channel.Item {
		fmt.Printf("%s\n", item.Title)
		fmt.Printf("%s\n", item.Description)
		
	}
}

func scrapeFeeds(s *state) {

	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println("Get Feed from database is unsuccessful!")
		return
	}

	timeNow := sql.NullTime{
		Time:  time.Now().UTC(),
		Valid: true,
	}

	markedParams := database.MarkFeedFetchedParams{
		UpdatedAt:     time.Now().UTC(),
		LastFetchedAt: timeNow,
		ID:            feed.ID,
	}

	err = s.db.MarkFeedFetched(context.Background(), markedParams)
	if err != nil {
		fmt.Println("Marking Feed from database is unsuccessful!")
		return
	}

	rss, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Println("Fetching Feed from RSS is unsuccessful!")
		return
	}

	printRSS(rss)
}
