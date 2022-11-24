package bot

import (
	"log"
	"strings"

	"github/Horsen121/TBD/RPBD/buy_list/api"
	"github/Horsen121/TBD/RPBD/buy_list/funcs"
	"github/Horsen121/TBD/RPBD/buy_list/service/conn"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start() {
	bot, err := tgbotapi.NewBotAPI("5757439968:AAEfKHe7vdACwxXayAGLsq20Z0DQxD64Cq4")
	if err != nil {
		log.Panic(err)
	}
	s, err := conn.NewStore("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin") // TODO: Change address to Docker
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	toBuyList := false
	blToPL := false
	anToPL := false
	openProduct := false
	changeStatus := false
	stats := false

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		user := update.Message.From.UserName
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ReplyMarkup = api.Buttons

			switch update.Message.Text {
			case "Add at list":
				toBuyList = true
				msg.Text = `Enter the name of the product, its weight and the time of the purchase reminder (if necessary) in the input line (through a space) and send the message`
				msg.ReplyMarkup = api.AddToList
			case "Add at refrigerator":
				msg.Text = `Where do you want to put the product from? (click the appropriate button)
				Enter the name of the product and its best before date in the input line (through a space) and send the message`
				msg.ReplyMarkup = api.AddToRefrigerator
			case "Open product":
				openProduct = true
				msg.Text = `Select product from the list and its new expiration date
				Enter the name of the product and its new best before date in the input line (through a space) and send the message`
				msg.ReplyMarkup = api.OpenProduct
			case "Change status":
				changeStatus = true
				msg.Text = `Select product from the list and select its status
				Enter the name of the product and its new status in the input line (through a space) and send the message`
				msg.ReplyMarkup = api.ChangeStatus
			case "Buy list":
				msg.Text = funcs.GetBuyList(s, user)
			case "Product list":
				msg.Text = funcs.GetProductList(s, user)
			case "Last products":
				msg.Text = funcs.GetLastProducts(s, user)
			case "Stats":
				stats = true
				msg.Text = `Shows stats of products for the specified period
				Enter the first and second date in the input line (through a space) and send the message`
			default:
				if toBuyList {
					// func ToBuyList
					data := strings.Split(msg.Text, " ")
					if err := funcs.AddToBuyList(s, data[0], data[1], data[2], user); err != "" {
						msg.Text = err
					}
					toBuyList = false
				} else if blToPL {
					// func ToProductList

					blToPL = false
				} else if anToPL {
					// func ToProductList

					anToPL = false
				} else if openProduct {
					// func OpenProduct

					openProduct = false
				} else if changeStatus {
					// func ChangeStatus

					changeStatus = false
				} else if stats {
					data := strings.Split(msg.Text, " ")
					msg.Text = funcs.GetStats(s, user, data[0], data[1])
					stats = false
				} else {
					msg.Text = "I don't know this command :("
				}
			}

			bot.Send(msg)
		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			// callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			// if _, err := bot.Request(callback); err != nil {
			// 	panic(err)
			// }
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

			switch update.CallbackQuery.Data {
			case "addFromList":
				blToPL = true
				msg.Text = funcs.GetProductList(s, user)
			case "addAnotherProduct":
				anToPL = true
			case "cancel":
				toBuyList = false
				blToPL = false
				anToPL = false
				openProduct = false
				changeStatus = false
				stats = false

				msg.Text = "Operation was canceled."
			}

			// And finally, send a message containing the data received.
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
