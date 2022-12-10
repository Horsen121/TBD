package api

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Buttons = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Add at list"),
		tgbotapi.NewKeyboardButton("Add at refrigerator"),
		tgbotapi.NewKeyboardButton("Open product"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Buy list"),
		tgbotapi.NewKeyboardButton("Product list"),
		tgbotapi.NewKeyboardButton("Last products"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Change status"),
		tgbotapi.NewKeyboardButton("Stats"),
	),
)

var Cancel = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Cancel"),
	),
)

var AddToRefrigerator = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("From list"),
		tgbotapi.NewKeyboardButton("Another"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Cancel"),
	),
)
