package funcs

import (
	"fmt"
	"github/Horsen121/TBD/RPBD/buy_list/service/conn"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var schedulerBL *gocron.Scheduler
var schedulerPL *gocron.Scheduler

func Start(s *conn.Store, user string, id int64) string {
	ischeated := false
	users, err := s.GetUsers()
	if err != nil {
		return "funcs: I'm sorry, but an error has occurred :("
	}

	for _, u := range users {
		if u.Name == user {
			ischeated = true
		}
	}
	if !ischeated {
		err := s.AddUser(user, id)
		if err != nil {
			return "funcs: I'm sorry, but an error has occurred :("
		}
	}

	return `Hello! I am a bot that will help you keep track of products from purchase to use.
				To start using me, click on the appropriate button and follow the instructions.`
}

func AddToBuyList(s *conn.Store, name string, weight string, reminder string, user string) string {
	_, err := time.Parse("2006-01-02", reminder)
	if err != nil {
		return "funcs: I'm sorry, but an error has occurred :("
	}

	if err := s.AddProductToBuyList(name, weight, reminder, user); err != nil {
		return "funcs: I'm sorry, but an error has occurred :("
	}

	return ""
}

func AddToProductList(s *conn.Store, name string, data string, user string, place string) string {
	_, err := time.Parse("2006-01-02", data)
	if err != nil {
		return "funcs: I'm sorry, but an error has occurred :("
	}

	if err := s.AddProductToProductList(name, data, user); err != nil {
		return "funcs: I'm sorry, but an error has occurred :("
	}
	if place == "bl" {
		if err := s.DeleteProductFromBuyList(name, user); err != nil {
			return "funcs: I'm sorry, but an error has occurred :("
		}
	}

	return ""
}

func ChangeStatus(s *conn.Store, name string, status string, user string) string {
	stat := false
	if status == "done" {
		stat = true
	}

	if err := s.AddProductToLastList(name, stat, user); err != nil {
		return "funcs: I'm sorry, but an error has occurred :("
	}
	if err := s.DeleteProductFromProductList(name, user); err != nil {
		return "funcs: I'm sorry, but an error has occurred :("
	}

	return ""
}

func OpenProduct(s *conn.Store, name string, data string, user string) string {
	if err := s.ChangeProductFromProductList(name, data, user); err != nil {
		return "funcs: I'm sorry, but an error has occurred :("
	}

	return ""
}

func GetBuyList(s *conn.Store, user string) string {
	var res string
	products, err := s.GetBuyList(user, "-1")
	if err != nil {
		res = "funcs: I'm sorry, but an error has occurred :("
	}
	if len(products) == 0 {
		return "List is empty"
	}

	for _, val := range products {
		time := strings.Split(val.Reminder.String(), " ")
		res += fmt.Sprintf("%s %v %s\n", val.Name, val.Weight, time[0])
	}

	return res
}

func GetProductList(s *conn.Store, user string) string {
	var res string
	products, err := s.GetProductList(user, "-1")
	if err != nil {
		res = "funcs: I'm sorry, but an error has occurred :("
	}
	if len(products) == 0 {
		return "List is empty"
	}

	for _, val := range products {
		time := strings.Split(val.Time.String(), " ")
		res += val.Name + " " + time[0] + "\n"
	}

	return res
}

func GetLastProducts(s *conn.Store, user string) string {
	var res string
	products, err := s.GetLastList(user, "-1", "-1")
	if err != nil {
		res = "funcs: I'm sorry, but an error has occurred :("
	}
	if len(products) == 0 {
		return "List is empty"
	}

	for _, val := range products {
		status := "done"
		if !val.Status {
			status = "casted"
		}
		res += val.Name + " - " + status + "\n"
	}

	return res
}

func GetStats(s *conn.Store, date1 string, date2 string, user string) string {
	var done, cast int
	products, err := s.GetLastList(user, date1, date2)
	if err != nil {
		return "funcs: I'm sorry, but an error has occurred :("
	}
	if len(products) == 0 {
		return "List is empty"
	}

	for _, val := range products {
		if val.Status {
			done++
		} else {
			cast++
		}
	}

	return fmt.Sprintf("Done! products - %v\nCasted products - %v", done, cast)
}

func Sheduling(bot *tgbotapi.BotAPI, s *conn.Store) {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		panic(err)
	}
	users, _ := s.GetUsers()
	for _, u := range users {
		user := u.Name
		chatId := u.Id

		schedulerBL = gocron.NewScheduler(loc)
		jbStart, err := schedulerBL.Every(1).Seconds().Do(CheckBuyList, bot, s, user, chatId)
		if err != nil {
			panic(err)
		}
		jbStart.LimitRunsTo(1)

		schedulerBL.Every(1).Day().At("09:00;18:45").Do(CheckBuyList, bot, s, user, chatId)
		schedulerBL.StartAsync()

		schedulerPL = gocron.NewScheduler(loc)
		jbStart1, err := schedulerBL.Every(1).Seconds().Do(CheckProductList, bot, s, user, chatId)
		if err != nil {
			panic(err)
		}
		jbStart1.LimitRunsTo(1)

		schedulerPL.Every(1).Day().At("09:00;18:45").Do(CheckProductList, bot, s, user, chatId)
		schedulerPL.StartAsync()
	}
}

func CheckBuyList(bot *tgbotapi.BotAPI, s *conn.Store, user string, id int64) {
	res := "You need to buy today:\n\n"
	data := strings.Split(time.Now().String(), " ")
	check := strings.Split(data[1], ":")
	if ((check[0] != "9") && (check[0] != "18")) || ((check[1] != "00") && (check[1] != "45")) {
		return
	}

	products, err := s.GetBuyList(user, data[0])
	if err != nil {
		res = "I'm sorry, but an error has occurred :("
	}
	if len(products) == 0 {
		return
	}

	for _, val := range products {
		res += fmt.Sprintf("%s %v \n", val.Name, val.Weight)
	}

	msg := tgbotapi.NewMessage(id, res)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func CheckProductList(bot *tgbotapi.BotAPI, s *conn.Store, user string, id int64) {
	res := "You need to check products today:\n\n"
	data := strings.Split(time.Now().String(), " ")
	check := strings.Split(data[1], ":")
	if ((check[0] != "9") && (check[0] != "18")) || ((check[1] != "00") && (check[1] != "45")) {
		return
	}

	products, err := s.GetProductList(user, data[0])
	if err != nil {
		res = "I'm sorry, but an error has occurred :("
	}

	for _, val := range products {
		res += fmt.Sprintf("%s \n", val.Name)
	}

	msg := tgbotapi.NewMessage(id, res)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
