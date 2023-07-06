package main

import (
	"fmt"
	"log"
	"rakabot/config"
	"rakabot/util"
)

func main() {
	util.Logging()

	chat, err := config.Configuration()
	if err != nil {
		log.Fatal(err)
		fmt.Printf("bot error. %s", err)
	}

	chat.Start()
	fmt.Println("bot started")
}
