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

var AddToList = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Cancel", "cancel"),
	),
)

var AddToRefrigerator = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("From list", "addFromList"),
		tgbotapi.NewInlineKeyboardButtonData("Another", "addAnotherProduct"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Cancel", "cancel"),
	),
)

var OpenProduct = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Cancel", "cancel"),
	),
)

var ChangeStatus = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Cancel", "cancel"),
	),
)

var Stats = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Cancel", "cancel"),
	),
)
