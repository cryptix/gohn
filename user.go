package gohn

// User is a user of HN
// see https://github.com/HackerNews/API#users for more
type User struct {
	About     string `json:"about"`
	Created   int    `json:"created"`
	Delay     int    `json:"delay"`
	ID        string `json:"id"`
	Karma     int    `json:"karma"`
	Submitted []int  `json:"submitted"`
}
