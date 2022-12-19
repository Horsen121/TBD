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

	return "New user was added successfuly"
}

func Autorisation(s *conn.Store, login string, password string) string {
	user, err := s.GetPasswordByLogin(login)
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

	return "User was autorised successfuly"
}
