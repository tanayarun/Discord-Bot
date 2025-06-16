package main

import (
	"fmt"

	"github.com/tanayarun/Discord-Bot/bot"
	"github.com/tanayarun/Discord-Bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	
	return
}
