package websocket

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Dialer        websocket.Dialer
	Conn          *websocket.Conn
	RequestHeader http.Header

	OnConnected     func()
	OnConnectError  func(err error)
	OnTextMessage   func(message string)
	OnBinaryMessage func(data []byte)
	OnDisconnected  func(err error)
}

func New(proxyUrl string) *Client {
	dialer := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
	}
	if proxyUrl != "" {
		proxyURL, err := url.Parse(proxyUrl)
		if err == nil {
			dialer.Proxy = http.ProxyURL(proxyURL)
		}
	}
	ws := &Client{
		Dialer:        dialer,
		RequestHeader: http.Header{},
	}
	return ws
}

func (c *Client) Connect(url string) {
	var err error
	c.Conn, _, err = c.Dialer.Dial(url, c.RequestHeader)
	if err != nil {
		if c.OnConnectError != nil {
			c.OnConnectError(err)
		}
		return
	}

	if c.OnConnected != nil {
		c.OnConnected()
	}

	defaultCloseHandler := c.Conn.CloseHandler()
	c.Conn.SetCloseHandler(func(code int, text string) error {
		if c.OnDisconnected != nil {
			c.OnDisconnected(errors.New(text))
		}
		return defaultCloseHandler(code, text)
	})

	go func() {
		for {
			messageType, message, err := c.Conn.ReadMessage()
			if err != nil {
				if c.OnDisconnected != nil {
					c.OnDisconnected(err)
				}
				return
			}

			switch messageType {
			case websocket.TextMessage:
				if c.OnTextMessage != nil {
					c.OnTextMessage(string(message))
				}
			case websocket.BinaryMessage:
				if c.OnBinaryMessage != nil {
					c.OnBinaryMessage(message)
				}
			}
		}
	}()
}

func (c *Client) Close() {
	err := c.Conn.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println(err)
	}
	_ = c.Conn.Close()
	if c.OnDisconnected != nil {
		c.OnDisconnected(err)
	}
}
