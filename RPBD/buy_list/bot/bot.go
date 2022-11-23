package bot

import (
	"log"

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

	user := bot.Self.UserName
	log.Printf("Authorized on account %s", user)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	toBuyList := false
	toProductList := false
	openProduct := false
	changeStatus := false
	getProductList := false
	getLastProductList := false
	getStats := false

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ReplyMarkup = api.Buttons

			switch update.Message.Text {
			case "Add at list":
				toBuyList = true
				msg.Text = `Enter the name of the product, its weight and the time of the purchase reminder (if necessary) 
								in the input line (through a space) and send the message`
				// msg.ReplyMarkup = api.AddToList
			case "Add at refrigerator":
				toProductList = true
				msg.Text = `Where do you want to put the product from? (click the appropriate button)
								\nEnter the name of the product and its best before date 
								in the input line (through a space) and send the message`
				msg.ReplyMarkup = api.AddToRefrigerator
			case "Open product":
				openProduct = true
				msg.Text = "Select product from the list and its new expiration date"
				// msg.ReplyMarkup = api.OpenProduct
			case "Change status":
				changeStatus = true
				msg.Text = "Select product from the list and select its status"
				// msg.ReplyMarkup = api.ChangeStatus
			case "Product list":
				getProductList = true
				msg.Text = "" // query
			case "Last products":
				getLastProductList = true
				msg.Text = "" // query lastProducts
			case "Stats":
				getStats = true
				msg.Text = "" // query to lastProducts
			default:
				var err error
				if toBuyList {
					// func ToBuyList

					toBuyList = false
				} else if toProductList {
					// func ToProductList

					toProductList = false
				} else if openProduct {
					// func OpenProduct

					openProduct = false
				} else if changeStatus {
					// func ChangeStatus

					changeStatus = false
				} else if getProductList {
					// func GetProductList
					msg.Text, err = funcs.GetProductList(s)
					getProductList = false
				} else if getLastProductList {
					// func GetLastProductList

					getLastProductList = false
				} else if getStats {
					// func GetStats

					getStats = false
				} else {
					msg.Text = "I don't know this command :("
				}
				if err != nil {
					msg.Text = "I'm sorry, but an error has occurred :("
				}
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
