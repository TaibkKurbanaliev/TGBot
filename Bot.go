package main

import (
	tele "gopkg.in/telebot.v3"
)

const ( // bot commands
	start = "/start"
	help  = "/help"
	about = "/about"
	info  = "/info"
)

type Bot struct {
	Name      string
	Token     string
	TeleBot   *tele.Bot
	Poller    tele.Poller
	IsStarted bool
}

func CreateBot(name string, token string, poller tele.Poller) (*Bot, error) {
	teleBot, err := tele.NewBot(tele.Settings{
		Token:  token,
		Poller: poller,
	})

	if err != nil {
		return nil, err
	}

	bot := &Bot{
		Name:      name,
		Token:     token,
		Poller:    poller,
		TeleBot:   teleBot,
		IsStarted: false,
	}

	bot.TeleBot.Handle(start, bot.start)
	bot.TeleBot.Handle(&btnHelp, bot.help)
	bot.TeleBot.Handle(&btnAbout, bot.about)
	bot.TeleBot.Handle(&btnInfo, bot.info)
	bot.TeleBot.Handle(&btnSubscribe, bot.subscribe)

	InitKeyboard()
	return bot, nil
}

func (b *Bot) start(c tele.Context) error {
	return c.Send("Hello "+c.Sender().FirstName, menu)
}

func (b *Bot) help(c tele.Context) error {
	return c.Send("Commands List:\n")
}

func (b *Bot) about(c tele.Context) error {
	// disks, err := GetDiskSpace()
	// if err != nil {
	// 	return c.Send(err.Error())
	// }

	// message := ""

	// for _, info := range disks {
	// 	message += info + "\n"
	// }

	return c.Send("Info")
}

func (b *Bot) info(c tele.Context) error {
	// coins := CurrentCoinsValues

	// var sb strings.Builder
	// sb.WriteString("Top coins:\n")

	// for _, coin := range coins {
	// 	sb.WriteString(fmt.Sprintf("  * %s (%s): $%.2f (Market Cap: $%.2f, Rank: %d)\n",
	// 		coin.Name, coin.Symbol, coin.CurrentPrice, float64(coin.MarketCap)/1e6, coin.MarketCapRank))
	// }

	return c.Send("Coins")
}

func (b *Bot) subscribe(c tele.Context) error {

	return c.Send("You have been subscribed")
}

func (b *Bot) Start() {
	b.IsStarted = true
	b.TeleBot.Start()
}
