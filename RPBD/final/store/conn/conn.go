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
}
type User struct {
	Id       int
	Name     string
	Surname  string
	Login    string
	Password string
	Status   bool
	Priority bool
	Type     bool
}
type Smena struct {
	Id          int
	User_id     int
	Started_at  time.Time
	Finished_at time.Time
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
	}, nil
}

// AddNewUser adds new user into users
func (s *Store) AddNewUser(name string, surname string, login string, password string, status bool, prioritet bool) error {
	_, err := s.conn.ExecContext(context.Background(), `INSERT INTO users(name, surname, login, password, status, prioritet) 
														VALUES($1, $2, $3, $4, true, $5);`, name, surname, login, password, prioritet)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}

// GetPasswordByLogin selects password of user from table by login
func (s *Store) GetUserByLogin(login string) (User, error) {
	user := User{}
	err := s.conn.SelectContext(context.Background(), &user, `SELECT * FROM users 
																WHERE login=$1;`, login)

	if err != nil {
		user.Id = -1
		return user, fmt.Errorf("query err: %w", err)
	}

	return user, nil
}

// ChangePriority changes priority of user
func (s *Store) ChangePriority(user_id int, prioritet bool) error {
	_, err := s.conn.ExecContext(context.Background(), `UPDATE users 
														SET prioritet=$1 
														WHERE id=$2;`, prioritet, user_id)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}

// AddSmena adds new smena into timetable
func (s *Store) AddSmena(user_id int, started_at time.Time, finished_at time.Time) error {
	_, err := s.conn.ExecContext(context.Background(), `INSERT INTO timetable (user_id, started_at, finished_at) 
														VALUES ($1, $2, $3);`, user_id, started_at, finished_at)

	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}

// GetSmenaById get's smena by id
func (s *Store) GetSmenaById(smena_id int) (Smena, error) {
	smena := Smena{}
	err := s.conn.SelectContext(context.Background(), &smena, `SELECT * FROM timetable 
																WHERE id=$1;`, smena_id)

	if err != nil {
		smena.Id = -1
		return smena, fmt.Errorf("query err: %w", err)
	}

	return smena, nil
}

// AddChange adds new offer into change
func (s *Store) AddChange(smena_id int, hoster_id int,
	coef float32, wonted_start time.Time, wonted_finish time.Time) error {

	smena, err := s.GetSmenaById(smena_id)
	_, err = s.conn.ExecContext(context.Background(),
		`INSERT INTO change (smena_id, started_at, finished_at, hoster_id, coef, wonted_start, wonted_finish, status) 
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
		err = s.conn.SelectContext(context.Background(), &list, `SELECT started_at, finished_at FROM timetable 
																	WHERE user_id=$1;`, user_id)
	} else {
		err = s.conn.SelectContext(context.Background(), &list, `SELECT user_id, started_at, finished_at FROM timetable
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
	_, err := s.conn.ExecContext(context.Background(), `UPDATE users 
														SET status=$1 
														WHERE id=$2;`, status, user_id)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}

// AddIll do transaction about ill
func (s *Store) AddIll(user_id int, started_at time.Time, finished_at time.Time, coef float32) error {
	var err error
	list := []Smena{}
	err = s.conn.SelectContext(context.Background(), &list, `SELECT smena_id, started_at, finished_at FROM timetable 
																WHERE user_id=$1 AND started_at>=$2 AND finished_at<=$3;`, user_id, started_at, finished_at)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	_, err = s.conn.ExecContext(context.Background(), `INSERT INTO illnes (user_id, started_at, finished_at) 
														VALUES ($1, $2, $3);`, user_id, started_at, finished_at)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	for _, smena := range list {
		_, err = s.conn.ExecContext(context.Background(),
			`BEGIN;
				INSERT INTO change (smena_id, started_at, finished_at, hoster_id, coef, status)
				VALUES ($4, $2, $3, 0, $5, false);
		
				DELETE FROM timetable WHERE user_id=$1 AND started_at>=$2 AND finished_at<=$3;
			END;`, user_id, started_at, finished_at, smena.Id, coef)

		if err != nil {
			return fmt.Errorf("found err: %w", err)
		}
	}

	return nil
}

// ChangeSmena do transaction about changing Smena
func (s *Store) ChangeSmena(smena_id, user_id int, started_at time.Time, finished_at time.Time,
	wonted_start time.Time, wonted_finish time.Time, ill bool) error {
	var err error

	if ill {
		_, err = s.conn.ExecContext(context.Background(),
			`BEGIN;
			UPDATE change SET status=true 
			WHERE smena_id=$1;

			INSERT INTO timetable (user_id, started_at, finished_at) 
			VALUES ($2, $3, $4);
		END;`, smena_id, user_id, started_at, finished_at)
	} else {
		_, err = s.conn.ExecContext(context.Background(),
			`BEGIN;
			UPDATE change SET status=true WHERE smena_id=$1;

			UPDATE timetable SET started_at=$5 AND finished_at=$6
			WHERE smena_id=$1;

			UPDATE timetable SET started_at=$3 AND finished_at=$4 
			WHERE user_id=$2;
		END;`, smena_id, user_id, started_at, finished_at, wonted_start, wonted_finish)
	}

	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}
