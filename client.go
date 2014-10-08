package gohn

import (
	"net/http"

	"github.com/bndr/gopencils"
)

type Client struct {
	Items ItemService
	// Users UserService
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{}
	c.Items = itemService{gopencils.Api("https://hacker-news.firebaseio.com/v0/", ".json", httpClient)}

	return c
}
