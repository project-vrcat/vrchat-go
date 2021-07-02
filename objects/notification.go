package objects

import (
	"time"
)

type (
	Notification struct {
		ID             string      `json:"id"`
		Type           string      `json:"type"`
		SenderUserID   string      `json:"senderUserId"`
		SenderUsername string      `json:"senderUsername"`
		ReceiverUserID string      `json:"receiverUserId"`
		Message        string      `json:"message"`
		Details        interface{} `json:"details"`
		CreatedAt      time.Time   `json:"created_at"`
	}
)

const (
	NotificationTypeInvite        = "invite"
	NotificationTypeFriendRequest = "friendRequest"
)
