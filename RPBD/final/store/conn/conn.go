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
	Status   bool
	Priority bool
}
type Smena struct {
	Id          int
	User_id     int
	Started_at  time.Time
	Finished_at time.Time
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

// GetProductList selects list of product from table for current user
func (s *Store) GetProductList(owner string, date string) ([]ProductList, error) {
	var err error
	list := []ProductList{}
	if date == "-1" {
		err = s.conn.SelectContext(context.Background(), &list, `SELECT pl.name FROM productlist AS pl
														WHERE owner = $1
														ORDER BY pl.name DESC;`, owner)
	} else {
		err = s.conn.SelectContext(context.Background(), &list, `SELECT pl.name FROM productlist AS pl
														WHERE owner = $1 AND time = $2
														ORDER BY pl.name DESC;`, owner, date)
	}
	if err != nil {
		return nil, fmt.Errorf("query err: %w", err)
	}

	return list, nil
}

func (s *Store) AddProductToBuyList(name string, weight string, reminder string, owner string) error {
	var err error
	if reminder != "-1" {
		_, err = s.conn.ExecContext(context.Background(), `INSERT INTO buylist(name, weight, reminder, owner) 
														VALUES($1, $2, $3, $4);`, name, weight, reminder, owner)
	} else {
		_, err = s.conn.ExecContext(context.Background(), `INSERT INTO buylist(name, weight, owner) 
														VALUES($1, $2, $3, $4);`, name, weight, owner)
	}
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}

// DeleteProductFromBuyList deletes product from BuyList
func (s *Store) DeleteProductFromBuyList(name string, owner string) error {
	var err error
	_, err = s.conn.ExecContext(context.Background(), `DELETE FROM buylist
													WHERE name = $1 AND owner = $2;`, name, owner)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}

func (s *Store) ChangeProductFromProductList(name string, data string, owner string) error {
	var err error
	_, err = s.conn.ExecContext(context.Background(), `UPDATE productlist
													SET time = $1
													WHERE name = $2 AND owner = $3;`, data, name, owner)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}
