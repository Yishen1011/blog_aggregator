package main

import (
	"fmt"
	"context"
)

func handlerAgg(s *state, cmd command) error {

	// if len(cmd.Args) < 1 {
	// 	return errors.New("The argument is empty\n")
	// }

	url := "https://www.wagslane.dev/index.xml"

	rss, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	printRSS(rss)
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