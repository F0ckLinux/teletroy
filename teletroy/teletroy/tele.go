package telebot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Host struct {
	ip  string
	uid string
}

var ConnectedTargets map[string]Host
var nowinfo string = "init"
var nowuse string = "init"
var my_ip string = "init"

func Me() string {
	if my_ip == "init" {
		resp, err := http.Get("https://ifconfig.co/ip")
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err == nil {
			my_ip = strings.TrimSpace(string(body))
		}
	}
	return my_ip
}

func Wait(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	//log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}
