package events

import (
	"github.com/project-vrcat/vrchat-go/objects"
)

const (
	FriendOnline   = "EventFriendOnline"
	FriendOffline  = "EventFriendOffline"
	FriendActive   = "EventFriendActive"
	FriendAdd      = "EventFriendAdd"
	FriendDelete   = "EventFriendDelete"
	FriendUpdate   = "EventFriendUpdate"
	FriendLocation = "EventFriendLocation"
	Notification   = "EventNotification"
)

type (
	FriendOnlineParams struct {
		UserID string       `json:"userId"`
		User   objects.User `json:"user"`
	}
	FriendOfflineParams struct {
		UserID string `json:"userId"`
	}
	FriendActiveParams struct {
		UserID string       `json:"userId"`
		User   objects.User `json:"user"`
	}
	FriendAddParams struct {
		UserID string       `json:"userId"`
		User   objects.User `json:"user"`
	}
	FriendDeleteParams struct {
		UserID string `json:"userId"`
	}
	FriendUpdateParams struct {
		UserID string       `json:"userId"`
		User   objects.User `json:"user"`
	}
	FriendLocationParams struct {
		UserID           string        `json:"userId"`
		User             objects.User  `json:"user"`
		World            objects.World `json:"world"`
		Location         string        `json:"location"`
		Instance         string        `json:"instance"`
		CanRequestInvite bool          `json:"canRequestInvite"`
	}
	NotificationParams objects.Notification
)
