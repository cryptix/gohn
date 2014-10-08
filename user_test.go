package gohn

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/cryptix/encodedTime"
)

func TestUserService_Get(t *testing.T) {
	setup()
	defer teardown()

	want := &User{
		About:     "This is a test",
		Created:   encodedTime.NewUnix(1173923446),
		Delay:     0,
		ID:        "jl",
		Karma:     2937,
		Submitted: []int{},
	}

	var called bool
	mux.HandleFunc("/v0/user/jl.json", func(w http.ResponseWriter, r *http.Request) {
		called = true
		testMethod(t, r, "GET")

		writeJSON(w, want)
	})

	u, err := client.Users.Get(want.ID)
	if err != nil {
		t.Errorf("Users.Get returned error: %v", err)
	}

	if !called {
		t.Fatal("!called")
	}

	if !reflect.DeepEqual(u, want) {
		t.Errorf("Users.Get returned %+v, want %+v", u, want)
	}
}
