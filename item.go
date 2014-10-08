package gohn

import "github.com/bndr/gopencils"

// Item represents a Story, Comments, Job, 'Ask HN' or a Poll
// see https://github.com/HackerNews/API#items for more
type Item struct {
	ID      int    `json:"id"`
	Deleted bool   `json:"deleted"`
	Type    string `json:"type"`
	By      string `json:"by"`
	Time    int    `json:"time"`
	Text    string `json:"text"`
	Dead    bool   `json:"dead"`
	Parent  int    `json:"parent"`
	Kids    []int  `json:"kids"`
	URL     string `json:"url"`
	Score   int    `json:"score"`
	Title   string `json:"title"`
	Parts   []int  `json:"parts"`
}

// ItemService has all methods that the firebase api exposes
type ItemService interface {
	TopStoryIDs() ([]int, error)
	MaxItemID() (int, error)
	Item(int) (*Item, error)
}

type itemService struct {
	api *gopencils.Resource
}

func (i itemService) TopStoryIDs() (ids []int, err error) {
	_, err = i.api.Res("topstories", &ids).Get()
	return
}

func (i itemService) MaxItemID() (id int, err error) {
	_, err = i.api.Res("maxitem", &id).Get()
	return
}

func (i itemService) Item(id int) (item *Item, err error) {
	item = new(Item)
	_, err = i.api.Res("item", item).Id(id).Get()
	return
}
