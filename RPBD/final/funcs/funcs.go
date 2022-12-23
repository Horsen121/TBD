package funcs

import (
	"fmt"
	"log"
	"time"

	"github.com/Horsen121/TBD/RPBD/final/store/conn"
	"golang.org/x/crypto/bcrypt"
)

//go:generate moq -out funcs_moq_test.go . Funcs
type Funcs interface {
	AddNewUser(name, surname, login, password string, status, prioritet bool) error
	GetUserByLogin(login string) (conn.User, error)
	ChangePriority(user_id int, prioritet bool) error
	AddSmena(user_id int, started_at, finished_at time.Time) error
	GetSmenaById(smena_id int) (conn.Smena, error)
	AddChange(smena_id, hoster_id int, coef float32, wonted_start, wonted_finish time.Time) error
	GetSmenaList(user_id int) ([]conn.Smena, error)
	ChangeUserStatus(user_id int, status bool) error
	AddIll(user_id int, started_at, finished_at time.Time, coef float32) error
	ChangeSmena(smena_id, user_id int, started_at, finished_at, wonted_start, wonted_finish time.Time, ill bool) error
}

var s *conn.Store

func Start(store *conn.Store) {
	s = store
}

func NewUser(admin, name, surname, login, password string, prioritet, user_type bool) string {
	user, err := s.GetUserByLogin(admin)
	if err != nil {
		log.Print(err.Error())
		return err.Error()
	}
	if !user.User_type {
		log.Print("You cannot add a new user. Contact to Admin")
		return "You cannot add a new user. Contact to Admin"
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {

	}

	if err := s.AddNewUser(name, surname, login, string(hash), true, prioritet, user_type); err != nil {
		log.Printf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}

	log.Print("New user was added successfuly")
	return "New user was added successfuly"
}

func Autorisation(login, password string) string {
	user, err := s.GetUserByLogin(login)
	if err != nil {
		if err.Error() == "User not found" {
			return err.Error()
		}
		log.Printf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Printf("found err: password is uncorrect!")
		return "found err: password is uncorrect!"
	}

	if !user.Status {
		log.Printf("found err: user was buned!")
		return "found err: user was buned!"
	}

	log.Print("User was autorised successfuly")
	return "User was autorised successfuly"
}

func ChangeUsersPriority(user_id int, prioritet bool) string {
	if err := s.ChangePriority(user_id, prioritet); err != nil {
		log.Printf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}

	log.Print("User's priority changed successfuly")
	return "User's priority changed successfuly"
}

func AddNewSmena(user_id int, started_at, finished_at time.Time) string {
	if err := s.AddSmena(user_id, started_at, finished_at); err != nil {
		log.Printf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}

	log.Print("New smena added successfuly")
	return "New smena added successfuly"
}

func AddChangeOffer(user_id, smena_id int, wonted_start, wonted_finish time.Time) string {
	smena, err := s.GetSmenaById(smena_id)
	if err != nil {
		log.Printf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}
	if smena.Id == -1 {
		log.Printf("Smena not found")
		return fmt.Sprint("Smena not found")
	}

	if s.AddChange(smena_id, user_id, 1, wonted_start, wonted_finish); err != nil {
		log.Printf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}

	log.Print("New change offer added successfuly")
	return "New change offer added successfuly"
}

func GetListOfSmena(user_id int) string {
	list, err := s.GetSmenaList(user_id)
	if err != nil {
		log.Printf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}
	if len(list) == 0 {
		log.Print("List is empty")
		return "List is empty"
	}

	var res string
	for _, smena := range list {
		res += fmt.Sprintf("%d %s\t%s\n", smena.Id, smena.StartData(), smena.FinishData())
	}

	log.Print("List of smena got successfuly")
	return res
}

func ChangeStatus(admin string, user_id int, status bool) string {
	log.Print("funcs")
	user, err := s.GetUserByLogin(admin)
	if err != nil {
		log.Print(err.Error())
		return err.Error()
	}
	if !user.User_type {
		log.Print("You cannot add a new user. Contact to Admin")
		return "You cannot add a new user. Contact to Admin"
	}
	if err := s.ChangeUserStatus(user_id, status); err != nil {
		log.Printf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}

	log.Print("Status was changed successfuly")
	return "Status was changed successfuly"
}

func Ill(user_id int, started_at, finished_at time.Time, coef float64) string {
	if err := s.AddIll(user_id, started_at, finished_at, coef); err != nil {
		log.Printf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}

	log.Print("Transaction of illness successfuly")
	return "Transaction of illness successfuly"
}

func ChangingSmena(smena_id, user_id int, wanted_start, wanted_finish time.Time, ill bool) string {
	smena, err := s.GetSmenaById(smena_id)
	if err != nil {
		log.Printf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}

	if err := s.ChangeSmena(smena_id, user_id, smena.Started_at, smena.Finished_at, wanted_start, wanted_finish, ill); err != nil {
		log.Printf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}

	log.Print("Transaction of illness successfuly")
	return "Transaction of illness successfuly"
}
