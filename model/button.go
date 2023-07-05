package model

import (
	"gopkg.in/telebot.v3"
)

var (
	Button        = &telebot.ReplyMarkup{}
	BtnSignal     = Button.Data("Serendipity Premium Signal", "premSignal")
	BtnPropFirm   = Button.Data("Prop Firm Phase Helper", "propHelper")
	BtnResiSignal = Button.Data("Upload Bukti Transfer", "uploadResiSignal")
	BtnResiProp   = Button.Data("Upload Bukti Transfer", "uploadResiProp")
)
