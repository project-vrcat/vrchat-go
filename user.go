package vrchat

import (
	"fmt"

	"github.com/project-vrcat/vrchat-go/objects"
)

type User struct {
	base
}

func (v *User) CurrentUser() (user objects.CurrentUser, err error) {
	_, err = v.client.R().
		SetResult(&user).
		Get("auth/user")
	return
}

func (v *User) Friends(offline bool) (list []objects.LimitedUser, err error) {
	_offline := "false"
	if offline {
		_offline = "true"
	}
	_, err = v.client.R().
		SetQueryParam("offline", _offline).
		SetResult(&list).
		Get("auth/user/friends")
	return
}

func (v *User) FriendStatus(uid string) (isFriend, outgoingRequest, incomingRequest bool, err error) {
	_url := fmt.Sprintf("user/%s/friendStatus", uid)
	var r struct {
		IsFriend        bool `json:"isFriend"`
		OutgoingRequest bool `json:"outgoingRequest"`
		IncomingRequest bool `json:"incomingRequest"`
	}
	_, err = v.client.R().
		SetResult(&r).
		Get(_url)
	return r.IsFriend, r.OutgoingRequest, r.IncomingRequest, err
}
