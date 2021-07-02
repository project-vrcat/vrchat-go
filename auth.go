package vrchat

import (
	"net/http/cookiejar"

	"github.com/project-vrcat/vrchat-go/objects"
)

func (v *VRCGO) Login(username, password string) (user objects.CurrentUser, err error) {
	v.base.jar, err = cookiejar.New(nil)
	if err != nil {
		return
	}
	v.client.SetCookieJar(v.base.jar)

	_, err = v.client.R().
		SetBasicAuth(username, password).
		SetResult(&user).
		Get("auth/user")
	v.CallEvent("Login", user)
	return
}

func (v *VRCGO) LoginWithSteam(steamTicket string) (user objects.CurrentUser, err error) {
	v.base.jar, err = cookiejar.New(nil)
	if err != nil {
		return
	}
	v.client.SetCookieJar(v.base.jar)

	_, err = v.client.R().
		SetBody(Map{"steamTicket": steamTicket}).
		SetResult(&user).
		Post("auth/steam")
	return
}

func (v *VRCGO) VerifyTOTP(code string) error {
	var r struct {
		Verified bool `json:"verified"`
	}
	_, err := v.client.R().
		SetBody(Map{"code": code}).
		SetResult(&r).
		Post("auth/twofactorauth/totp/verify")
	if !r.Verified {
		return Err2FAFailed
	}
	return err
}

func (v *VRCGO) VerifyOTP(code string) error {
	var r struct {
		Verified bool `json:"verified"`
	}
	_, err := v.client.R().
		SetBody(Map{"code": code}).
		SetResult(&r).
		Post("auth/twofactorauth/otp/verify")
	if !r.Verified {
		return Err2FAFailed
	}
	return err
}

func (v *VRCGO) Logout() error {
	_, err := v.client.R().Put("logout")
	return err
}

func (v *VRCGO) VerifyAuthToken() (token string, err error) {
	var r struct {
		Token string `json:"token"`
	}
	_, err = v.client.R().
		SetResult(&r).
		Get("auth")
	if err != nil {
		return
	}
	token = r.Token
	return
}

func (v *VRCGO) AuthToken() (token string, err error) {
	cookies, err := v.Cookies()
	if err != nil {
		return
	}
	err = ErrAuthTokenNotFound
	for _, cookie := range cookies {
		if cookie.Name == "auth" {
			return cookie.Value, nil
		}
	}
	return
}
