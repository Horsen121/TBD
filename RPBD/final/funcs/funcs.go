package funcs

import (
	"time"
	// "github.com/Horsen121/TBD/RPBD/final/store/conn"
)

//go:generate moq -out funcs_moq_test.go . Funcs
type Funcs interface {
	AddNewUser(name string, surname string, login string, password string, status bool, prioritet bool) error
	GetPasswordByLogin(login string) (User, error)
	ChangePriority(user_id int, prioritet bool) error
	AddSmena(user_id int, started_at time.Time, finished_at time.Time) error
	AddChange(smena_id int, started_at time.Time, finished_at time.Time, hoster_id int,
		coef float32, wonted_start time.Time, wonted_finish time.Time) error
	GetSmenaList(user_id int) ([]Smena, error)
	ChangeUserStatus(user_id int, status bool) error
	AddIll(user_id int, started_at time.Time, finished_at time.Time, coef float32) error
	ChangeSmena(smena_id, user_id int, started_at time.Time, finished_at time.Time,
		wonted_start time.Time, wonted_finish time.Time, ill bool) error
}
