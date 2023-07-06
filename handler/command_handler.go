package handler

import (
	"log"
	"rakabot/model"
	"rakabot/util"

	"gopkg.in/telebot.v3"
)

func StartHandler() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		log.Printf("sender id %d started bot", c.Sender().ID)

		button := model.Button
		button.Inline(
			button.Row(model.BtnSignal),
			button.Row(model.BtnPropFirm),
		)

		err := c.Send(util.DotEnv("REPLY_START"), button)

		if err != nil {
			log.Fatal(err)
		}

		return nil
	}
}

func SignalHandler() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		log.Printf("sender id %d pressed premium signal button", c.Sender().ID)

		button := model.Button

		button.Inline(
			button.Row(model.BtnResiSignal),
		)

		reply := "- PREMIUM SERENDIPITY FX SIGNAL SUBSCRIPTION -\n\nApabila Anda berminat untuk berlangganan SerendipityFX Premium Signal silahkan melakukan transfer ke rekening berikut :\n\nNama Bank : BCA\nNomor Rekening : 0391616433\n\nApabila sudah melakukan transfer, silahkan upload bukti transfer dengan menekan tombol di bawah. Admin akan segera memproses permintaan Anda"

		err := c.Send(reply, button)

		if err != nil {
			log.Fatal(err)
		}

		return nil
	}
}

func PropFirmHandler() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		log.Printf("sender id %d pressed prop firm button", c.Sender().ID)

		button := model.Button

		button.Inline(
			button.Row(model.BtnResiProp),
		)

		reply := "- JASA PROP FIRM -\n\nApabila Anda berminat terhadap Jasa Prop Firm agar dapat lolos pada kedua Phase silahkan melakukan transfer ke rekening berikut :\n\nNama Bank : BCA\nNomor Rekening : 0391616433\n\nApabila sudah melakukan transfer, silahkan upload bukti transfer dengan menekan tombol di bawah. Admin akan segera memproses permintaan Anda"

		err := c.Send(reply, button)

		if err != nil {
			log.Fatal(err)
		}

		return nil
	}
}
