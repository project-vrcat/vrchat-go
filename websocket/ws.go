package websocket

import (
	"log"

	jsoniter "github.com/json-iterator/go"
	"github.com/project-vrcat/vrchat-go/events"
	"github.com/project-vrcat/vrchat-go/pkg/websocket"
)

type (
	Pipeline struct {
		ws     *websocket.Client
		handle func(string, interface{})
	}
	event struct {
		Type    string `json:"type"`
		Content string `json:"content"`
	}
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func New(proxyUrl, clientUserAgent string, handle func(string, interface{})) *Pipeline {
	client := websocket.New(proxyUrl)
	client.RequestHeader.Add("User-Agent", clientUserAgent)
	w := &Pipeline{
		ws:     client,
		handle: handle,
	}
	w.ws.OnConnected = func() {
		log.Println("OnConnected")
	}
	w.ws.OnDisconnected = func(err error) {
		log.Println("OnDisconnected", err)
	}
	w.ws.OnConnectError = func(err error) {
		log.Println("OnConnectError", err)
	}
	w.ws.OnTextMessage = func(message string) {
		onMessage(message, handle)
	}
	return w
}

func (w *Pipeline) Connect(authToken string) {
	w.ws.Connect("wss://pipeline.vrchat.cloud/?authToken=" + authToken)
}

func (w *Pipeline) Close() {
	w.ws.Close()
}

func onMessage(message string, handle func(string, interface{})) {
	var e event
	if json.UnmarshalFromString(message, &e) == nil {
		switch e.Type {
		case "friend-online":
			var info events.FriendOnlineParams
			if json.UnmarshalFromString(e.Content, &info) == nil {
				go handle(events.FriendOnline, info)
			}
		case "friend-offline":
			var info events.FriendOfflineParams
			if json.UnmarshalFromString(e.Content, &info) == nil {
				go handle(events.FriendOffline, info)
			}
		case "friend-active":
			var info events.FriendActiveParams
			if json.UnmarshalFromString(e.Content, &info) == nil {
				go handle(events.FriendActive, info)
			}
		case "friend-add":
			var info events.FriendAddParams
			if json.UnmarshalFromString(e.Content, &info) == nil {
				go handle(events.FriendAdd, info)
			}
		case "friend-delete":
			var info events.FriendDeleteParams
			if json.UnmarshalFromString(e.Content, &info) == nil {
				go handle(events.FriendDelete, info)
			}
		case "friend-update":
			var info events.FriendUpdateParams
			if json.UnmarshalFromString(e.Content, &info) == nil {
				go handle(events.FriendUpdate, info)
			}
		case "friend-location":
			var info events.FriendLocationParams
			if json.UnmarshalFromString(e.Content, &info) == nil {
				go handle(events.FriendLocation, info)
			}
		case "notification":
			var info events.NotificationParams
			if json.UnmarshalFromString(e.Content, &info) == nil {
				go handle(events.Notification, info)
			}
		default:
			log.Println(e)
		}
	}
}
