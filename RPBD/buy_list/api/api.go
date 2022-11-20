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
		tgbotapi.NewKeyboardButton("Change status"),
		tgbotapi.NewKeyboardButton("Product list"),
		tgbotapi.NewKeyboardButton("Last products"),
		tgbotapi.NewKeyboardButton("Stats"),
	),
)

var AddToList = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Name", "name"),
		tgbotapi.NewInlineKeyboardButtonData("Weight", "weight"),
		tgbotapi.NewInlineKeyboardButtonData("Notice", "notice"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Add", "ToList"),
	),
)
var AddToRefrigerator = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("From list", "query"), /// query to Product list
		tgbotapi.NewInlineKeyboardButtonData("Another", "product"),
		tgbotapi.NewInlineKeyboardButtonData("Time", "time"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Add", "ToRefrigerator"),
	),
)
var OpenProduct = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Product", "product"), /// query to Product list
		tgbotapi.NewInlineKeyboardButtonData("New time", "time"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Open", "Open"),
	),
)
var ChangeStatus = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Product", "product"), /// query to Product list
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Done", "Done"),
		tgbotapi.NewInlineKeyboardButtonData("Cast", "Cast"),
	),
)
