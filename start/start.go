package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.heise.de/rss/heise-atom.xml")
	for i := range 5 {
		fmt.Println(feed.Items[i].Title)
		fmt.Println(feed.Items[i].Link)
		fmt.Println(feed.Items[i].Description)
	}
}
