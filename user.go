package gohn

import (
	"github.com/bndr/gopencils"
	"github.com/cryptix/encodedTime"
)

// User is a user of HN
// see https://github.com/HackerNews/API#users for more
type User struct {
	About     string           `json:"about"`
	Created   encodedTime.Unix `json:"created"`
	Delay     int              `json:"delay"`
	ID        string           `json:"id"`
	Karma     int              `json:"karma"`
	Submitted []int            `json:"submitted"`
}

// UserService has all methods that the firebase api exposes for users
type UserService interface {
	Get(id string) (*User, error)
}

type userService struct {
	api *gopencils.Resource
}

func (u userService) Get(id string) (user *User, err error) {
	user = new(User)
	_, err = u.api.Res("user", user).Id(id).Get()
	return
}
