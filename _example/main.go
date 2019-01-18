package main

import (
	"fmt"

	nnn "github.com/aoisensi/go-niconamanotify"
)

func main() {
	notifyer := &nnn.Notifyer{}
	notifyer.Handler = func(info *nnn.StreamInfo) {
		fmt.Println(info.StreamInfo.Title)
		fmt.Println(info.StreamInfo.Description)
		fmt.Print("https://nico.ms/")
		fmt.Println(info.RequestID)
		fmt.Println()
	}
	notifyer.Init(nil)
	defer notifyer.Close()
	notifyer.Listen()
}
