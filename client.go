// Package gohn implements a client for the Hacker News REST API from firebaseio
package gohn

import (
	"net/http"

	"github.com/bndr/gopencils"
)

// Client exposes the different services for the HackerNews api
type Client struct {
	Items ItemService
	Users UserService
}

// NewClient returns a new api client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{}
	api := gopencils.Api("https://hacker-news.firebaseio.com/v0/", ".json", httpClient)
	c.Items = itemService{api}
	c.Users = userService{api}

	return c
}
