package vrchat

import (
	"strconv"
	"time"

	"github.com/project-vrcat/vrchat-go/objects"
)

func (v *VRCGO) Visits() (int, error) {
	resp, err := v.client.R().Get("visits")
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(resp.String())
}

func (v *VRCGO) RemoteConfig() (config objects.RemoteConfig, err error) {
	_, err = v.client.R().SetResult(&config).Get("config")
	return
}

func (v *VRCGO) Health() (bool, error) {
	var r struct {
		OK bool `json:"ok"`
	}
	_, err := v.client.R().SetResult(&r).Get("health")
	return r.OK, err
}

func (v *VRCGO) SystemTime() (t time.Time, err error) {
	resp, err := v.client.R().Get("time")
	if err != nil {
		return
	}
	return time.ParseInLocation(`"2006-01-02T15:04:05+00:00"`, resp.String(), time.UTC)
}
