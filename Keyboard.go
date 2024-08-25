package main

import tele "gopkg.in/telebot.v3"

var (
	menu         = &tele.ReplyMarkup{ResizeKeyboard: true}
	selector     = &tele.ReplyMarkup{}
	btnHelp      = menu.Text("ℹ Help")
	btnAbout     = menu.Text("❓ About")
	btnInfo      = menu.Text("₿ Info")
	btnSubscribe = menu.Text("🔔 Subscribe")
	btnPrev      = selector.Data("⬅", "prev")
	btnNext      = selector.Data("➡", "next")
)

func InitKeyboard() {
	menu.Reply(
		menu.Row(btnHelp),
		menu.Row(btnAbout),
		menu.Row(btnInfo),
		menu.Row(btnSubscribe),
	)
	selector.Inline(
		selector.Row(btnPrev, btnNext),
	)
}
