package vrchat

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/project-vrcat/vrchat-go/websocket"
)

type (
	base struct {
		client *resty.Client
		jar    *cookiejar.Jar
	}
	VRCGO struct {
		base
		sync.Mutex
		events   map[string][]func(interface{})
		User     User
		Pipeline *websocket.Pipeline
	}
	Options struct {
		ClientUserAgent string
		Timeout         time.Duration
		RetryCount      int
		ProxyURL        string
	}
	Map map[string]interface{}
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var (
	Err2FAFailed         = errors.New("2FA verification failed")
	ErrAuthTokenNotFound = errors.New("AuthToken not found")
)

func NewClient() *VRCGO {
	return NewClientWithOptions(Options{
		Timeout:    time.Second * 5,
		RetryCount: 3,
	})
}

func NewClientWithOptions(opt Options) *VRCGO {
	client := resty.New()

	client.JSONMarshal = json.Marshal
	client.JSONUnmarshal = json.Unmarshal

	if opt.ClientUserAgent == "" {
		opt.ClientUserAgent = "vrchat-go/version 0.1.0"
	}
	if opt.Timeout == 0 {
		opt.Timeout = time.Second * 5
	}
	if opt.ProxyURL != "" {
		client.SetProxy(opt.ProxyURL)
	}

	client.
		SetHostURL("https://api.vrchat.cloud/api/1").
		SetHeader("User-Agent", opt.ClientUserAgent).
		SetTimeout(opt.Timeout).
		SetRetryCount(opt.RetryCount)

	base := base{
		client: client,
	}
	vrc := &VRCGO{
		base:   base,
		events: map[string][]func(interface{}){},
		User:   User{base},
	}
	vrc.Pipeline = websocket.New(opt.ProxyURL, opt.ClientUserAgent, vrc.CallEvent)

	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		switch resp.StatusCode() {
		case http.StatusOK:
			return nil
		default:
			return errors.New(errorMessage(resp.String()))
		}
	})
	return vrc
}

func (v *VRCGO) Cookies() ([]*http.Cookie, error) {
	return v.jar.Cookies(&url.URL{
		Scheme: "https",
		Host:   "api.vrchat.cloud",
	}), nil
}

func errorMessage(body string) (message string) {
	var r struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.UnmarshalFromString(body, &r); err != nil {
		return ""
	}
	return r.Error.Message
}
