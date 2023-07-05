package main

import (
	"fmt"
	"log"
	"os"
	"rakabot/config"
	"time"
)

func main() {

	currentTime := time.Now()
	timeFormat := currentTime.Format("2006-01-02_15-04-05")
	fileName := fmt.Sprintf("log_%s.txt", timeFormat)

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	log.SetOutput(file)

	chat, err := config.Configuration()
	if err != nil {
		log.Fatal(err)
	}

	chat.Start()
}
