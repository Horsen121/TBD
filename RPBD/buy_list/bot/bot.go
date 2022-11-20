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
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ReplyMarkup = api.Buttons

			switch update.Message.Text {
			case "Add at list":
				msg.Text = "Enter the name of the product, its weight and the time of the purchase reminder (if necessary)"
				msg.ReplyMarkup = api.AddToList
			case "Add at refrigerator":
				msg.Text = "Enter the name of the product (or select from the list) and its expiration date"
				msg.ReplyMarkup = api.AddToRefrigerator
			case "Open product":
				msg.Text = "Select product from the list and its new expiration date"
				msg.ReplyMarkup = api.OpenProduct
			case "Change status":
				msg.Text = "Select product from the list and select its status"
				msg.ReplyMarkup = api.ChangeStatus
			case "Product list":
				msg.Text = "" // query
			case "Last products":
				msg.Text = "" // query
			case "Stats":
				msg.Text = "" // query
			}

			bot.Send(msg)
		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			// And finally, send a message containing the data received.
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
