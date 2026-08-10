package main

import (
	"context"
	"database/sql"
	"fmt"
	"errors"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/Yishen1011/blog_aggregator/internal/database"
	"strconv"
	"time"
	
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

	// printRSS(rss)

	// Save posts
	for _, item := range rss.Channel.Item {

		_, err = s.db.GetPostFromURL(context.Background(), item.Link)
		if err == nil {
			fmt.Printf("Post already exist in the database. %s\n", item.Link)
			continue
		} else if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Creating post in the database.")
		} else {
			fmt.Println("Error retrieving posts from database:", err)
			continue
		}

		desc := sql.NullString{
			String: item.Description,
			Valid: true,
		}

		// Parse Published Date
		publishedAt := sql.NullTime{}
		var parsedSuccessfully bool
		layouts := []string{
			time.RFC1123,
			time.RFC1123Z,
			time.RFC822,
			time.RFC822Z,
			time.RFC3339,
		}

		for _, layout := range layouts {
			publishedDate, parseErr := time.Parse(layout, item.PubDate)
			if parseErr == nil {
					publishedAt = sql.NullTime{
					Time:  publishedDate,
					Valid: true,
				}
				parsedSuccessfully = true
				break
			}
		}

		if !parsedSuccessfully {
			fmt.Println("Error parsing date:", item.PubDate)
			publishedAt = sql.NullTime{
				Valid: false,
			}
		}
		
		postParams := database.CreatePostParams{
			ID:            uuid.New(),
			CreatedAt:     time.Now().UTC(),
			UpdatedAt:     time.Now().UTC(),
			Title:         item.Title,
			Url:           item.Link,
			Description:   desc,
			PublishedAt:   publishedAt,
			FeedID:        feed.ID,
		}

		_, err = s.db.CreatePost(context.Background(), postParams)
		if err != nil {
			var dbErr *pq.Error
			if errors.As(err, &dbErr) && dbErr.Code == "23505" {
				fmt.Printf("Post with link %s already exists.\n", item.Link)
				continue
			} 
			fmt.Println("Error: not able to save record:", err)
		}
	}
}

func handlerBrowse(s *state, cmd command, user database.User) error {

	var limit int32
	if len(cmd.Args) < 1 {
		limit = 2
	} else if len(cmd.Args) == 1 {
		limitVal, parseErr := strconv.Atoi(cmd.Args[0])
		if parseErr != nil {
			return errors.New("Unsuccessful! Parsing string to number.\n")
		}
		limit = int32(limitVal)
	}

	getPostParams := database.GetPostForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}

	posts, err := s.db.GetPostForUser(context.Background(), getPostParams)
	if err != nil {
		return errors.New("Unsuccessful! Getting posts from database.\n")
	}

	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Url: %s\n", post.Url)
		if post.Description.Valid {
			fmt.Printf("Description: %s\n", post.Description.String)
		}
		if post.PublishedAt.Valid {
			fmt.Printf("PublishedAt: %v\n", post.PublishedAt.Time)
		}
		fmt.Println()
	}

	return nil
}
