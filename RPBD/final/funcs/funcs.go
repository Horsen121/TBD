package funcs

import (
	"fmt"
	"log"
	"time"

	"github.com/Horsen121/TBD/RPBD/final/store/conn"
)

//go:generate moq -out funcs_moq_test.go . Funcs
type Funcs interface {
	AddNewUser(name string, surname string, login string, password string, status bool, prioritet bool) error
	GetPasswordByLogin(login string) (conn.User, error)
	ChangePriority(user_id int, prioritet bool) error
	AddSmena(user_id int, started_at time.Time, finished_at time.Time) error
	AddChange(smena_id int, started_at time.Time, finished_at time.Time, hoster_id int,
		coef float32, wonted_start time.Time, wonted_finish time.Time) error
	GetSmenaList(user_id int) ([]conn.Smena, error)
	ChangeUserStatus(user_id int, status bool) error
	AddIll(user_id int, started_at time.Time, finished_at time.Time, coef float32) error
	ChangeSmena(smena_id, user_id int, started_at time.Time, finished_at time.Time,
		wonted_start time.Time, wonted_finish time.Time, ill bool) error
}

func NewUser(s *conn.Store, name string, surname string, login string, password string, prioritet bool) string {
	if err := s.AddNewUser(name, surname, login, password, true, prioritet); err != nil {
		log.Fatalf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}

	log.Fatal("New user was added successfuly")
	return "New user was added successfuly"
}

func Autorisation(s *conn.Store, login string, password string) string {
	user, err := s.GetUserByLogin(login)
	if err != nil {
		log.Fatalf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}

	if user.Id == -1 {
		log.Fatal("found err: user don't found!")
		return "found err: user don't found!"
	}

	if user.Password != password {
		log.Fatal("found err: password is uncorrect!")
		return "found err: password is uncorrect!"
	}

	if !user.Status {
		log.Fatal("found err: user was buned!")
		return "found err: user was buned!"
	}

	log.Fatal("User was autorised successfuly")
	return "User was autorised successfuly"
}

func ChangeUsersPriority(s *conn.Store, user_id int, prioritet bool) string {
	if err := s.ChangePriority(user_id, prioritet); err != nil {
		log.Fatalf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}

	log.Fatal("User's priority changed successfuly")
	return "User's priority changed successfuly"
}

func AddNewSmena(s *conn.Store, user_id int, started_at time.Time, finished_at time.Time) string {
	if err := s.AddSmena(user_id, started_at, finished_at); err != nil {
		log.Fatalf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}

	log.Fatal("New smena added successfuly")
	return "New smena added successfuly"
}

func AddChangeOffer(s *conn.Store, user_id int, smena_id int, wonted_start time.Time, wonted_finish time.Time) string {
	smena, err := s.GetSmenaById(smena_id)
	if err != nil {
		log.Fatalf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}
	if smena.Id == -1 {
		log.Fatalf("Smena not found")
		return fmt.Sprint("Smena not found")
	}

	if s.AddChange(smena_id, user_id, 1, wonted_start, wonted_finish); err != nil {
		log.Fatalf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}

	log.Fatal("New change offer added successfuly")
	return "New change offer added successfuly"
}

func GetListOfSmena(s *conn.Store, user_id int) string {
	list, err := s.GetSmenaList(user_id)
	if err != nil {
		log.Fatalf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}
	if len(list) == 0 {
		log.Fatal("List is empty")
		return "List is empty"
	}

	var res string
	for _, smena := range list {
		res += fmt.Sprintf("%d %s\t%s", smena.Id, smena.StartData(), smena.FinishData())
	}

	log.Fatal("List of smena got successfuly")
	return res
}

func ChangeStatus(s *conn.Store, user_id int, status bool) string {
	if err := s.ChangeUserStatus(user_id, status); err != nil {
		log.Fatalf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}

	log.Fatal("Status was changed successfuly")
	return "Status was changed successfuly"
}

func Ill(s *conn.Store, user_id int, started_at time.Time, finished_at time.Time, coef float32) string {
	if err := s.AddIll(user_id, started_at, finished_at, coef); err != nil {
		log.Fatalf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}

	log.Fatal("Transaction of illness successfuly")
	return "Transaction of illness successfuly"
}

func ChangingSmena(s *conn.Store, smena_id, user_id int, wanted_start time.Time, wanted_finish time.Time, ill bool) string {
	smena, err := s.GetSmenaById(smena_id)
	if err != nil {
		log.Fatalf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}

	if err := s.ChangeSmena(smena_id, user_id, smena.Started_at, smena.Finished_at, wanted_start, wanted_finish, ill); err != nil {
		log.Fatalf("found err: %s", err.Error())
		return fmt.Sprintf("found err: %s", err.Error())
	}

	log.Fatal("Transaction of illness successfuly")
	return "Transaction of illness successfuly"
}
