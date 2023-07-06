package config

import (
	"log"
	"rakabot/handler"
	"rakabot/model"
	"rakabot/util"
	"time"

	"gopkg.in/telebot.v3"
)

func Configuration() (*telebot.Bot, error) {
	bot_token := util.DotEnv("BOT_TOKEN")

	if bot_token == "" {
		log.Fatal("Bot Token Can't Be Empty")
	} else {
		log.Println("Bot Token Environtment Variables OK")
	}

	preferences := telebot.Settings{
		Token:  bot_token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	chatbot, err := telebot.NewBot(preferences)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	chatbot.Handle("/start", handler.StartHandler())
	chatbot.Handle(&model.BtnSignal, handler.SignalHandler())
	chatbot.Handle(&model.BtnPropFirm, handler.PropFirmHandler())
	chatbot.Handle(&model.BtnResiSignal, handler.ResiSignalHandler())
	chatbot.Handle(&model.BtnResiProp, handler.ResiPropHandler())
	chatbot.Handle(telebot.OnPhoto, handler.PhotoHandler())

	return chatbot, nil
}
