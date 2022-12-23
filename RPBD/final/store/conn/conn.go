package conn

import (
	"context"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	conn *sqlx.DB
	ctx  context.Context
}
type User struct {
	Id        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Surname   string `json:"surname" db:"surname"`
	Login     string `json:"login" db:"login"`
	Password  string `json:"password" db:"password"`
	Status    bool   `json:"status" db:"status"`
	Prioritet bool   `json:"prioritet" db:"prioritet"`
	User_type bool   `json:"user_type" db:"user_type"`
}
type Smena struct {
	Id          int       `json:"id" db:"id"`
	User_id     int       `json:"user_id" db:"user_id"`
	Started_at  time.Time `json:"started_at" db:"started_at"`
	Finished_at time.Time `json:"finished_at" db:"finished_at"`
}

func (s Smena) StartData() string {
	return fmt.Sprintf("%d-%s-%d %d:%d", s.Started_at.Day(), s.Started_at.Month().String(), s.Started_at.Year(),
		s.Started_at.Hour(), s.Started_at.Minute())
}
func (s Smena) FinishData() string {
	return fmt.Sprintf("%d-%s-%d %d:%d", s.Finished_at.Day(), s.Finished_at.Month().String(), s.Finished_at.Year(),
		s.Finished_at.Hour(), s.Finished_at.Minute())
}

// NewStore creates new database connection
func NewStore(connString string) (*Store, error) {
	conn, err := sqlx.Connect("postgres", connString)
	if err != nil {
		panic(err)
	}

	// migrates(connString)

	return &Store{
		conn: conn,
		ctx:  context.Background(),
	}, nil
}

// AddNewUser adds new user into users
func (s *Store) AddNewUser(name, surname, login, password string, status, prioritet, user_type bool) error {
	_, err := s.conn.ExecContext(s.ctx, `INSERT INTO users(name, surname, login, password, status, prioritet, user_type) 
											VALUES($1, $2, $3, $4, true, $5, $6);`, name, surname, login, password, prioritet, user_type)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}

// GetPasswordByLogin selects password of user from table by login
func (s *Store) GetUserByLogin(login string) (User, error) {
	user := []User{}
	err := s.conn.SelectContext(s.ctx, &user, `SELECT * FROM users 
												WHERE login=$1;`, login)
	if err != nil {
		return User{}, fmt.Errorf("query err: %w", err)
	}
	if len(user) == 0 {
		return User{}, fmt.Errorf("User not found")
	}

	return user[0], nil
}

// ChangePriority changes priority of user
func (s *Store) ChangePriority(user_id int, prioritet bool) error {
	_, err := s.conn.ExecContext(s.ctx, `UPDATE users 
											SET prioritet=$1 
											WHERE id=$2;`, prioritet, user_id)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}

// AddSmena adds new smena into timetable
func (s *Store) AddSmena(user_id int, started_at, finished_at time.Time) error {
	_, err := s.conn.ExecContext(s.ctx, `INSERT INTO timetable (user_id, started_at, finished_at) 
											VALUES ($1, $2, $3);`, user_id, started_at, finished_at)

	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}

// GetSmenaById get's smena by id
func (s *Store) GetSmenaById(smena_id int) (Smena, error) {
	smena := []Smena{}
	err := s.conn.SelectContext(s.ctx, &smena, `SELECT * FROM timetable 
													WHERE id=$1;`, smena_id)

	if err != nil {
		return Smena{}, fmt.Errorf("query err: %w", err)
	}
	if len(smena) == 0 {
		return Smena{}, fmt.Errorf("Smena not found")
	}

	return smena[0], nil
}

// AddChange adds new offer into change
func (s *Store) AddChange(smena_id, hoster_id int, coef float32, wonted_start, wonted_finish time.Time) error {

	smena, err := s.GetSmenaById(smena_id)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}
	_, err = s.conn.ExecContext(s.ctx,
		`INSERT INTO change (smena_id, started_at, finished_at, hoster_id, coef, wanted_start, wanted_finish, status) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, false);`, smena_id, smena.Started_at, smena.Finished_at, hoster_id, coef, wonted_start, wonted_finish)

	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}

// GetSmenaList selects list of smena from table
func (s *Store) GetSmenaList(user_id int) ([]Smena, error) {
	var err error
	list := []Smena{}
	if user_id != 0 {
		err = s.conn.SelectContext(s.ctx, &list, `SELECT started_at, finished_at FROM timetable 
																	WHERE user_id=$1;`, user_id)
	} else {
		err = s.conn.SelectContext(s.ctx, &list, `SELECT user_id, started_at, finished_at FROM timetable
																	ORDER BY started_at DESC
																	LIMIT 20;`)
	}

	if err != nil {
		return nil, fmt.Errorf("query err: %w", err)
	}

	return list, nil
}

// ChangeUserStatus changes status of user
func (s *Store) ChangeUserStatus(user_id int, status bool) error {
	_, err := s.conn.ExecContext(s.ctx, `UPDATE users 
											SET status=$1 
											WHERE id=$2;`, status, user_id)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}

// AddIll do transaction about ill
func (s *Store) AddIll(user_id int, started_at, finished_at time.Time, coef float64) error {
	var err error
	list := []Smena{}
	err = s.conn.SelectContext(s.ctx, &list, `SELECT id, started_at, finished_at FROM timetable 
												WHERE user_id=$1 AND started_at>=$2 AND finished_at<=$3;`, user_id, started_at, finished_at)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	_, err = s.conn.ExecContext(s.ctx, `INSERT INTO ill (user_id, d_start, d_finish) 
											VALUES ($1, $2, $3);`, user_id, started_at, finished_at)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	for _, smena := range list {
		tx := s.conn.MustBegin()
		tx.MustExec(`INSERT INTO change (smena_id, started_at, finished_at, hoster_id, coef, status)
						VALUES ($1, $2, $3, 1, $4, false);`, smena.Id, started_at, finished_at, coef)
		tx.MustExec(`DELETE FROM timetable WHERE user_id=$1 AND started_at>=$2 AND finished_at<=$3;`, user_id, started_at, finished_at)
		err = tx.Commit()

		if err != nil {
			return fmt.Errorf("found err: %w", err)
		}
	}

	return nil
}

// ChangeSmena do transaction about changing Smena
func (s *Store) ChangeSmena(smena_id, user_id int, started_at, finished_at, wonted_start, wonted_finish time.Time, ill bool) error {
	var err error

	if ill {
		tx := s.conn.MustBegin()
		tx.MustExec(`UPDATE change SET status=true 
						WHERE smena_id=$1;`, smena_id)
		tx.MustExec(`INSERT INTO timetable (user_id, started_at, finished_at) 
						VALUES ($1, $2, $3);`, user_id, started_at, finished_at)
		err = tx.Commit()
	} else {
		tx := s.conn.MustBegin()
		tx.MustExec(`UPDATE change SET status=true 
						WHERE smena_id=$1;`, smena_id)
		tx.MustExec(`UPDATE timetable SET started_at=$2, finished_at=$3
						WHERE id=$1;`, smena_id, wonted_start, wonted_finish)
		tx.MustExec(`UPDATE timetable SET started_at=$2, finished_at=$3 
						WHERE user_id=$1;`, user_id, started_at, finished_at)
		err = tx.Commit()
	}

	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}
