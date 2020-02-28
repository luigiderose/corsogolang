package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("784952636:AAFgl0kxew_qADVC2iAxEV8DRFuE67955dc")

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		user := update.Message.Chat.FirstName

		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//scritto1 := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		scritto := update.Message.Text

		soprannome := update.Message.From.UserName

		//msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}

		//scrive sul server monitor
		log.Printf("L'utente %s %s ha scritto %s\n", user, soprannome, scritto)

		//generatore del file log
		f, err := os.OpenFile("text.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			log.Println(err)
		}

		defer f.Close()

		logger := log.New(f, "Defoult ", log.LstdFlags)
		loggertext := fmt.Sprintf("L'utente %s %s ha scritto %s", user, soprannome, scritto)
		logger.Println(loggertext)

	}
}
