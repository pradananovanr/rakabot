package handler

import (
	"fmt"
	"log"
	"strings"

	"gopkg.in/telebot.v3"
)

type Admin struct {
	ChatID string
}

func (adm Admin) Recipient() string {
	return adm.ChatID
}

func ResiSignalHandler() telebot.HandlerFunc {
	return func(c telebot.Context) error {

		reply := fmt.Sprintf("Silahkan Upload Bukti Transfer/Resi Anda disertai dengan caption pada gambar (Pastikan untuk ditulis pada caption, tidak pada chat terpisah)\n\nCaption : %s", "PremSignal-"+fmt.Sprint(c.Sender().Username))

		err := c.Send(reply)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("sender id %d pressed resi signal button", c.Sender().ID)

		return nil
	}
}

func ResiPropHandler() telebot.HandlerFunc {
	return func(c telebot.Context) error {

		reply := fmt.Sprintf("Silahkan Upload Bukti Transfer/Resi Anda disertai dengan caption pada gambar (Pastikan untuk ditulis pada caption, tidak pada chat terpisah)\n\nCaption : %s", "PropFirm-"+fmt.Sprint(c.Sender().Username))

		err := c.Send(reply)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("sender id %d pressed resi prop firm button", c.Sender().ID)

		return nil
	}
}

func PhotoHandler() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		caption := c.Message().Caption

		if caption != "" {
			if strings.Contains(caption, "PropFirm") {
				err := c.Send("Terima kasih telah mengirimkan bukti transfer/resi untuk Jasa Prop Firm. Mohon ditunggu, Admin akan menghubungi Anda sesegera mungkin")

				log.Printf("sender id %d sending photo for prop firm", c.Sender().ID)

				if err != nil {
					log.Fatal(err)
				}
			} else if strings.Contains(caption, "PremSignal") {
				err := c.Send("Terima kasih telah mengirimkan bukti transfer/resi untuk SerendipityFX Premium Signal. Mohon ditunggu, Admin akan menghubungi Anda sesegera mungkin")

				log.Printf("sender id %d sending photo for premium signal", c.Sender().ID)

				if err != nil {
					log.Fatal(err)
				}
			} else {
				err := c.Send("Maaf, caption yang Anda masukkan tidak dikenali. Silahkan mengirimkan ulang dengan caption yang sesuai")

				log.Printf("sender id %d sending photo with unrecognized caption", c.Sender().ID)

				if err != nil {
					log.Fatal(err)
				}
			}

			if strings.Contains(caption, "PremSignal") || strings.Contains(caption, "PropFirm") {
				recipient := &telebot.User{ID: 693523350}

				err := c.ForwardTo(recipient)

				log.Printf("forwarding messages to admin from %d", c.Sender().ID)

				if err != nil {
					log.Fatal(err)
				}
			}

		} else {
			err := c.Send("Pastikan bukti transfer/resi terdapat caption seperti yang telah diinstruksikan sebelumnya")

			log.Printf("sender id %d sending photo without caption", c.Sender().ID)

			if err != nil {
				log.Fatal(err)
			}
		}

		return nil
	}
}

func ResponseHandler() telebot.HandlerFunc {
	return func(c telebot.Context) error {
		response := &telebot.CallbackResponse{
			Text:      "Premium Serendipity Signal Subscription",
			ShowAlert: true,
		}
		err := c.Respond(response)

		log.Printf("sender id %d pressed premium signal button", c.Sender().ID)

		if err != nil {
			log.Fatal(err)
		}

		return nil
	}
}
