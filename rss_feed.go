package main

import (
	"io"
	"net/http"
	"encoding/xml"
	"context"
	"html"
	"fmt"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	request, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil) // Return (*Request, err)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", "gator")

	response, err := httpClient.Do(request) // Return (Response, err)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300{
		return nil, fmt.Errorf("unexpected status: %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	rss := RSSFeed{}
	if err := xml.Unmarshal(data, &rss); err != nil {
		return nil, err
	}

	cleanFeed(&rss)

	return &rss, nil
}

func cleanFeed(rss *RSSFeed) {
	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)

	for i := range rss.Channel.Item {
		rss.Channel.Item[i].Title = html.UnescapeString(rss.Channel.Item[i].Title)
		rss.Channel.Item[i].Description = html.UnescapeString(rss.Channel.Item[i].Description)
	}
}