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

	funcs.Sheduling(bot, s)

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
				msg.Text = `Enter the name of the product, its weight and the time of the purchase reminder (if necessary) in format 2022-11-25 in the input line (through a space) and send the message`
				msg.ReplyMarkup = api.Cancel
			case "Add at refrigerator":
				msg.Text = `Where do you want to put the product from? (click the appropriate button)
				Enter the name of the product and its best before date (in format 2022-11-25) in the input line (through a space) and send the message`
				msg.ReplyMarkup = api.AddToRefrigerator
			case "Open product":
				openProduct = true
				msg.Text = `Select product from the list and its new expiration date
				Enter the name of the product and its new best before date (in format 2022-11-25) in the input line (through a space) and send the message
				
				`
				msg.Text += funcs.GetProductList(s, user)
				msg.ReplyMarkup = api.Cancel
			case "Change status":
				changeStatus = true
				msg.Text = `Select product from the list and select its status
				Enter the name of the product and its new status ("done" or "cast") in the input line (through a space) and send the message
				`
				msg.Text += funcs.GetProductList(s, user)
				msg.ReplyMarkup = api.Cancel
			case "Buy list":
				msg.Text = funcs.GetBuyList(s, user)
			case "Product list":
				msg.Text = funcs.GetProductList(s, user)
			case "Last products":
				msg.Text = funcs.GetLastProducts(s, user)
			case "Stats":
				stats = true
				msg.Text = `Shows stats of products for the specified period
				Enter the first and second date in the input line (through a space) and send the message
				(if you if you don't need some date, specify -1)`
			case "From list":
				blToPL = true
				msg.ReplyMarkup = api.Cancel
				msg.Text = funcs.GetBuyList(s, user)
			case "Another":
				anToPL = true
				msg.ReplyMarkup = api.Cancel
			case "Cancel":
				toBuyList = false
				blToPL = false
				anToPL = false
				openProduct = false
				changeStatus = false
				stats = false

				msg.Text = "Operation was canceled."
			case "/start":
				msg.Text = funcs.Start(s, user, msg.ChatID)
			default:
				if toBuyList {
					// func ToBuyList
					data := strings.Split(msg.Text, " ")
					if len(data) == 3 {
						if err := funcs.AddToBuyList(s, data[0], data[1], data[2], user); err != "" {
							msg.Text = err
						}
					} else if len(data) == 2 {
						if err := funcs.AddToBuyList(s, data[0], data[1], "-1", user); err != "" {
							msg.Text = err
						}
					}
					msg.Text = "Product added successfully"

					toBuyList = false
				} else if blToPL {
					// func ToProductList
					data := strings.Split(msg.Text, " ")
					if err := funcs.AddToProductList(s, data[0], data[1], user, "bl"); err != "" {
						msg.Text = err
					}

					msg.Text = "Product added successfully"
					blToPL = false
				} else if anToPL {
					// func ToProductList
					data := strings.Split(msg.Text, " ")
					if err := funcs.AddToProductList(s, data[0], data[1], user, "-1"); err != "" {
						msg.Text = err
					}

					msg.Text = "Product added successfully"
					anToPL = false
				} else if openProduct {
					// func OpenProduct
					data := strings.Split(msg.Text, " ")
					if err := funcs.OpenProduct(s, data[0], data[1], user); err != "" {
						msg.Text = err
					}

					msg.Text = "Product's status added successfully"
					openProduct = false
				} else if changeStatus {
					// func ChangeStatus
					data := strings.Split(msg.Text, " ")
					if err := funcs.ChangeStatus(s, data[0], data[1], user); err != "" {
						msg.Text = err
					}

					msg.Text = "Product's status added successfully"
					changeStatus = false
				} else if stats {
					data := strings.Split(msg.Text, " ")
					msg.Text = funcs.GetStats(s, data[0], data[1], user)
					stats = false
				} else {
					msg.Text = "I don't know this command :("
				}
			}

			bot.Send(msg)
		}
	}
}
