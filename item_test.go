package gohn

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/cryptix/go/encodedTime"
)

func TestItemService_TopStoryIDs(t *testing.T) {
	setup()
	defer teardown()

	want := []int{23, 42, 1337}

	var called bool
	mux.HandleFunc("/v0/topstories.json", func(w http.ResponseWriter, r *http.Request) {
		called = true
		testMethod(t, r, "GET")

		writeJSON(w, want)
	})

	ids, err := client.Items.TopStoryIDs()
	if err != nil {
		t.Errorf("Items.TopStoryIDs returned error: %v", err)
	}

	if !called {
		t.Fatal("!called")
	}

	if !reflect.DeepEqual(ids, want) {
		t.Errorf("Items.TopStoryIDs returned %+v, want %+v", ids, want)
	}
}

func TestItemService_TopStories(t *testing.T) {
	setup()
	defer teardown()

	want := []*Item{
		&Item{
			By:    "testUser",
			ID:    23,
			Kids:  []int{},
			Score: 1,
			Time:  encodedTime.NewUnix(1),
			Title: "Wierd test story #1",
			Type:  "story",
			URL:   "http://www.testurl.fake/1",
		},
		&Item{
			By:    "testUser",
			ID:    24,
			Kids:  []int{},
			Score: 2,
			Time:  encodedTime.NewUnix(2),
			Title: "Wierd test story #2",
			Type:  "story",
			URL:   "http://www.testurl.fake/2",
		},
		&Item{
			By:    "testUser",
			ID:    128,
			Kids:  []int{},
			Score: 3,
			Time:  encodedTime.NewUnix(3),
			Title: "Wierd test story #3",
			Type:  "story",
			URL:   "http://www.testurl.fake/3",
		},
	}

	called := 0
	mux.HandleFunc("/v0/topstories.json", func(w http.ResponseWriter, r *http.Request) {
		called += 1
		testMethod(t, r, "GET")
		writeJSON(w, []int{23, 42, 128})
	})

	mux.HandleFunc("/v0/item/23.json", func(w http.ResponseWriter, r *http.Request) {
		called += 1
		testMethod(t, r, "GET")
		writeJSON(w, want[0])
	})

	mux.HandleFunc("/v0/item/42.json", func(w http.ResponseWriter, r *http.Request) {
		called += 1
		testMethod(t, r, "GET")
		writeJSON(w, want[1])
	})

	mux.HandleFunc("/v0/item/128.json", func(w http.ResponseWriter, r *http.Request) {
		called += 1
		testMethod(t, r, "GET")
		writeJSON(w, want[2])
	})

	stories, err := client.Items.TopStories()
	if err != nil {
		t.Errorf("Items.TopStoryIDs returned error: %v", err)
	}

	if called != 4 {
		t.Fatalf("called != 4. got %d", called)
	}

	if !reflect.DeepEqual(stories, want) {
		t.Errorf("Items.TopStoryIDs returned %+v, want %+v", stories, want)
	}

}

func TestItemService_MaxItemID(t *testing.T) {
	setup()
	defer teardown()

	var want = 1337

	var called bool
	mux.HandleFunc("/v0/maxitem.json", func(w http.ResponseWriter, r *http.Request) {
		called = true
		testMethod(t, r, "GET")

		writeJSON(w, want)
	})

	id, err := client.Items.MaxItemID()
	if err != nil {
		t.Errorf("Items.MaxItemID returned error: %v", err)
	}

	if !called {
		t.Fatal("!called")
	}

	if !reflect.DeepEqual(id, want) {
		t.Errorf("Items.MaxItemID returned %+v, want %+v", id, want)
	}
}

func TestItemService_ItemStory(t *testing.T) {
	setup()
	defer teardown()

	want := &Item{
		By:    "dhouston",
		ID:    8863,
		Kids:  []int{8952, 9224, 8917, 8884, 8887, 8943, 8869, 8958, 9005, 9671, 8940, 9067, 8908, 9055, 8865, 8881, 8872, 8873, 8955, 10403, 8903, 8928, 9125, 8998, 8901, 8902, 8907, 8894, 8878, 8870, 8980, 8934, 8876},
		Score: 111,
		Time:  encodedTime.NewUnix(1175714200),
		Title: "My YC app: Dropbox - Throw away your USB drive",
		Type:  "story",
		URL:   "http://www.getdropbox.com/u/2/screencast.html",
	}

	var called bool
	mux.HandleFunc("/v0/item/8863.json", func(w http.ResponseWriter, r *http.Request) {
		called = true
		testMethod(t, r, "GET")

		writeJSON(w, want)
	})

	story, err := client.Items.Item(want.ID)
	if err != nil {
		t.Errorf("Items.Item returned error: %v", err)
	}

	if !called {
		t.Fatal("!called")
	}

	if !reflect.DeepEqual(story, want) {
		t.Errorf("Items.Item returned %+v, want %+v", story, want)
	}
}

func TestItemService_ItemComment(t *testing.T) {
	setup()
	defer teardown()

	want := &Item{
		By:     "norvig",
		ID:     2921983,
		Kids:   []int{2922097, 2922429, 2924562, 2922709, 2922573, 2922140, 2922141},
		Parent: 2921506,
		Text:   "Aw shucks, guys ... you make me blush with your compliments.<p>Tell you what, Ill make a deal: I'll keep writing if you keep reading. K?",
		Time:   encodedTime.NewUnix(1314211127),
		Type:   "comment",
	}

	var called bool
	mux.HandleFunc("/v0/item/2921983.json", func(w http.ResponseWriter, r *http.Request) {
		called = true
		testMethod(t, r, "GET")

		writeJSON(w, want)
	})

	story, err := client.Items.Item(want.ID)
	if err != nil {
		t.Errorf("Items.Item returned error: %v", err)
	}

	if !called {
		t.Fatal("!called")
	}

	if !reflect.DeepEqual(story, want) {
		t.Errorf("Items.Item returned %+v, want %+v", story, want)
	}
}

func TestItemService_ItemPoll(t *testing.T) {
	setup()
	defer teardown()

	want := &Item{
		By:    "pg",
		ID:    126809,
		Kids:  []int{126822, 126823, 126993, 126824, 126934, 127411, 126888, 127681, 126818, 126816, 126854, 127095, 126861, 127313, 127299, 126859, 126852, 126882, 126832, 127072, 127217, 126889, 127535, 126917, 126875},
		Parts: []int{126810, 126811, 126812},
		Score: 46,
		Text:  "",
		Time:  encodedTime.NewUnix(1204403652),
		Title: "Poll: What would happen if News.YC had explicit support for polls?",
		Type:  "poll",
	}

	var called bool
	mux.HandleFunc("/v0/item/126809.json", func(w http.ResponseWriter, r *http.Request) {
		called = true
		testMethod(t, r, "GET")

		writeJSON(w, want)
	})

	story, err := client.Items.Item(want.ID)
	if err != nil {
		t.Errorf("Items.Item returned error: %v", err)
	}

	if !called {
		t.Fatal("!called")
	}

	if !reflect.DeepEqual(story, want) {
		t.Errorf("Items.Item returned %+v, want %+v", story, want)
	}
}

func TestItemService_ItemPollOpt(t *testing.T) {
	setup()
	defer teardown()

	want := &Item{
		By:     "pg",
		ID:     160705,
		Parent: 160704,
		Score:  335,
		Text:   "Yes, ban them; I'm tired of seeing Valleywag stories on News.YC.",
		Time:   encodedTime.NewUnix(1207886576),
		Type:   "pollopt",
	}

	var called bool
	mux.HandleFunc("/v0/item/160705.json", func(w http.ResponseWriter, r *http.Request) {
		called = true
		testMethod(t, r, "GET")

		writeJSON(w, want)
	})

	story, err := client.Items.Item(want.ID)
	if err != nil {
		t.Errorf("Items.Item returned error: %v", err)
	}

	if !called {
		t.Fatal("!called")
	}

	if !reflect.DeepEqual(story, want) {
		t.Errorf("Items.Item returned %+v, want %+v", story, want)
	}
}
