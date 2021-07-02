package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/project-vrcat/vrchat-go"
	"github.com/project-vrcat/vrchat-go/events"
)

func init() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
}

func main() {
	vrc := vrchat.NewClient()

	vrc.RegisterEvent(events.FriendUpdate, func(params interface{}) {
		p := params.(events.FriendUpdateParams)
		log.Println("Friend Update", p.User.DisplayName)
	})
	vrc.RegisterEvent(events.FriendLocation, func(params interface{}) {
		p := params.(events.FriendLocationParams)
		log.Println("Friend Location", p.User.DisplayName, p.World.Name)
	})
	vrc.RegisterEvent(events.Notification, func(params interface{}) {
		p := params.(events.NotificationParams)
		log.Println("Notification", p.SenderUsername, p.Type)
	})

	_, err := vrc.RemoteConfig()
	if err != nil {
		panic(err)
	}

	user, err := vrc.Login(os.Getenv("VRC_USN"), os.Getenv("VRC_PWD"))
	if err != nil {
		panic(err)
	}
	if user.Requires2FA() {
		var code string
		fmt.Print("2FA Code: ")
		_, _ = fmt.Scan(&code)
		if err = vrc.VerifyTOTP(code); err != nil {
			panic(err)
		}
		user, err = vrc.User.CurrentUser()
		if err != nil {
			panic(err)
		}
	}
	defer vrc.Logout()

	fmt.Println("Hello,", user.DisplayName)

	token, err := vrc.AuthToken()
	if err != nil {
		panic(err)
	}

	vrc.Pipeline.Connect(token)
	defer vrc.Pipeline.Close()

	fmt.Println("Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
