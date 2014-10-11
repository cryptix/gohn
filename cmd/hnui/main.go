package main

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/andlabs/ui"
	"github.com/cryptix/gohn"
	"github.com/skratchdot/open-golang/open"
)

type displayItem struct {
	Idx                        int
	Type, Title, Text, URL, By string
}

var (
	window ui.Window
	status ui.Label
	hn     *gohn.Client

	newItem chan *displayItem
)

func main() {

	go ui.Do(func() {
		reload := ui.NewButton("Reload")
		status = ui.NewStandaloneLabel("")

		table := ui.NewTable(reflect.TypeOf(displayItem{}))

		// on click > open url of selected item
		table.OnSelected(func() {
			idx := table.Selected()
			if idx >= 0 {
				table.RLock()
				td := table.Data().(*[]displayItem)
				open.Start((*td)[idx].URL)
				table.RUnlock()
			}
		})

		// iterates over chan and updates items in the table
		go func() {
			newItem = make(chan *displayItem)

			table.Lock()
			td := table.Data().(*[]displayItem)
			*td = make([]displayItem, 100)
			table.Unlock()

			for item := range newItem {
				status.SetText(fmt.Sprintf("got item %d", item.Idx))
				table.Lock()
				td := table.Data().(*[]displayItem)
				(*td)[item.Idx] = *item
				table.Unlock()
			}
		}()

		// construct main view stack
		stack := ui.NewVerticalStack(
			reload,
			status,
			table)
		stack.SetStretchy(2)

		window = ui.NewWindow("HNui - TopStories", 200, 500, stack)

		reload.OnClicked(func() {
			go updateItems(newItem)
		})

		window.OnClosing(func() bool {
			ui.Stop()
			return true
		})

		window.Show()
	})

	hn = gohn.NewClient(nil)

	err := ui.Go()
	if err != nil {
		panic(err)
	}
}

func updateItems(ichan chan<- *displayItem) {
	status.SetText("Update started...")

	// get the current top IDs
	ids, err := hn.Items.TopStoryIDs()
	if err != nil {
		status.SetText(fmt.Sprint("TopStoryIDs() Err: ", err.Error()))
		return
	}

	// start a go routine for each ID and push them to the displayItem channel
	var wg sync.WaitGroup
	for idx, id := range ids {
		wg.Add(1)
		go updateItem(&wg, ichan, idx, id)
	}

	wg.Wait()
	status.SetText("Update done!")
}

func updateItem(wg *sync.WaitGroup, ichan chan<- *displayItem, i, id int) {
	defer wg.Done()

	item, err := hn.Items.Item(id)
	// try again if it fails (EOF, timeout, ...)
	if err != nil {
		fmt.Println("Items() Err:", err)
		ichan <- &displayItem{
			Idx:   i,
			Type:  "Err",
			Title: item.Title,
			Text:  err.Error(),
		}

		wg.Add(1)
		go updateItem(wg, ichan, i, id)
		return
	}

	ichan <- &displayItem{
		Idx:   i,
		Type:  item.Type,
		Title: item.Title,
		Text:  item.Text,
		URL:   item.URL,
		By:    item.By,
	}
}
