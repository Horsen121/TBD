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

func AddToProductList(s *conn.Store, name string, data string, user string, place string) string {
	_, err := time.Parse("2006-01-02", data)
	if err != nil {
		return fmt.Sprintf("err: %s", err.Error()) // "funcs: I'm sorry, but an error has occurred :("
	}

	if err := s.AddProductToProductList(name, data, user); err != nil {
		return fmt.Sprintf("found err: %s", err.Error())
	}
	if place == "bl" {
		if err := s.DeleteProductFromBuyList(name, user); err != nil {
			return fmt.Sprintf("found err: %s", err.Error())
		}
	}

	return ""
}

func ChangeStatus(s *conn.Store, name string, status string, user string) string {
	if err := s.AddProductToLastList(name, status, user); err != nil {
		return fmt.Sprintf("found err: %s", err.Error())
	}
	if err := s.DeleteProductFromProductList(name, user); err != nil {
		return fmt.Sprintf("found err: %s", err.Error())
	}

	return ""
}

func OpenProduct(s *conn.Store, name string, data string, user string) string {
	if err := s.ChangeProductFromProductList(name, data, user); err != nil {
		return fmt.Sprintf("found err: %s", err.Error())
	}

	return ""
}

func GetBuyList(s *conn.Store, user string) string {
	var res string
	products, err := s.GetBuyList(user, "-1")
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
	products, err := s.GetProductList(user, "-1")
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
		res = fmt.Sprintf("found err: %s", err.Error()) //"I'm sorry, but an error has occurred :("
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

func CheckBuyList(s *conn.Store, user string) string {
	res := "You need to buy today:\n\n"
	products, err := s.GetBuyList(user, "-1")
	if err != nil {
		res = fmt.Sprintf("err: %s", err.Error()) //"I'm sorry, but an error has occurred :("
	}

	for _, val := range products {
		res += fmt.Sprintf("%s %v \n", val.Name, val.Weight)
	}

	return res
}

func CheckProductList(s *conn.Store, user string) string {
	res := "You need to check products today:\n\n"
	products, err := s.GetProductList(user, "-1")
	if err != nil {
		res = fmt.Sprintf("err: %s", err.Error()) //"I'm sorry, but an error has occurred :("
	}

	for _, val := range products {
		res += fmt.Sprintf("%s \n", val.Name)
	}

	return res
}
