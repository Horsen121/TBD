package funcs

import (
	"fmt"
	"github/Horsen121/TBD/RPBD/buy_list/service/conn"
	"time"
)

func AddToBuyList(s *conn.Store, name string, weight string, reminder string, user string) string {
	_, err := time.Parse("2006-01-02", reminder)
	if err != nil {
		return fmt.Sprintf("err: %s", err.Error()) // "funcs: I'm sorry, but an error has occurred :("
	}

	if err := s.AddProductToBuyList(name, weight, reminder, user); err != nil {
		return fmt.Sprintf("found err: %s", err.Error())
	}

	return ""
}

func GetBuyList(s *conn.Store, user string) string {
	var res string
	products, err := s.GetBuyList(user)
	if err != nil {
		res = fmt.Sprintf("err: %s", err.Error()) //"I'm sorry, but an error has occurred :("
	}

	for _, val := range products {
		res += fmt.Sprintf("%s %v \n", val.Name, val.Weight)
	}

	return res
}

func GetProductList(s *conn.Store, user string) string {
	var res string
	products, err := s.GetProductList(user)
	if err != nil {
		res = "I'm sorry, but an error has occurred :("
	}

	for _, val := range products {
		res += val.Name + "\n"
	}

	return res
}

func GetLastProducts(s *conn.Store, user string) string {
	var res string
	products, err := s.GetLastList(user, "-1", "-1")
	if err != nil {
		res = "I'm sorry, but an error has occurred :("
	}

	for _, val := range products {
		res += val.Name + "\n"
	}

	return res
}

func GetStats(s *conn.Store, user string, date1 string, date2 string) string {
	var done, cast int
	products, err := s.GetLastList(user, date1, date2)
	if err != nil {
		return "I'm sorry, but an error has occurred :("
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
