package bot

import (
	"log"

	"github/Horsen121/TBD/RPBD/buy_list/api"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start() {
	bot, err := tgbotapi.NewBotAPI("5757439968:AAEfKHe7vdACwxXayAGLsq20Z0DQxD64Cq4")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ReplyMarkup = api.Buttons

			switch update.Message.Text {
			case "Add at list":
			case "Add at refrigirator":
			case "Open product":
			case "Change status":
			case "Product list":
			case "Last products":
			case "Stats":
			}

			bot.Send(msg)
		}
	}
}
