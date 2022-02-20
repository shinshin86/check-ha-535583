package main

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/mmcdole/gofeed"
)

type FeedItem struct {
	Title           string
	Link            string
	PublishedParsed string
}

func display(items []FeedItem) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].PublishedParsed > items[j].PublishedParsed
	})

	for _, item := range items {
		titleAndDate := item.Title + ": " + item.PublishedParsed + "\n"
		fmt.Printf("\x1b[36m%s\x1b[0m", titleAndDate)
		fmt.Println(item.Link)
		fmt.Println("=================================")
	}
}

func main() {
	var items []FeedItem

	feedURL := "https://a.hatena.ne.jp/shinshin86/rss?gid=535583"
	feed, err := gofeed.NewParser().ParseURL(feedURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, item := range feed.Items {
		items = append(items, FeedItem{
			Title:           item.Title,
			Link:            item.Link,
			PublishedParsed: item.PublishedParsed.Format(time.RFC3339),
		})
	}

	display(items)
}
