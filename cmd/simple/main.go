package main

import (
	"fmt"
	"sync"

	"github.com/cryptix/gohn"
)

var hn *gohn.Client

func main() {
	// NewClient returns a new Hacker News API client.
	// If httpClient is nil, http.DefaultClient is used.
	hn = gohn.NewClient(nil)

	ids, err := hn.Items.TopStoryIDs()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for idx, id := range ids {

		wg.Add(1)
		go displayItem(&wg, idx, id)
	}

	wg.Wait()
}

func displayItem(wg *sync.WaitGroup, i, id int) {
	defer wg.Done()
	item, err := hn.Items.Item(id)
	if err != nil {
		fmt.Println("Err:", err)
		wg.Add(1)
		displayItem(wg, i, id)
		return
	}
	fmt.Print(i, "–", item.Title, "\n   ", item.URL, "\n")
}
