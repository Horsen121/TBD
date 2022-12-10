package conn

import (
	"context"
	"fmt"
	"strings"
	"time"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	conn *sqlx.DB
}
type User struct {
	Id   int64
	Name string
}
type ProductList struct {
	Name string
	Time time.Time
}

func (p ProductList) Date() string {
	return fmt.Sprintf("%d-%s-%d", p.Time.Day(), p.Time.Month().String(), p.Time.Year())
}

type BuyList struct {
	Name     string
	Weight   float64
	Reminder time.Time
}

func (b BuyList) Date() string {
	return fmt.Sprintf("%d-%s-%d", b.Reminder.Day(), b.Reminder.Month().String(), b.Reminder.Year())
}

type LastProducts struct {
	Name   string
	Status bool
	Date   time.Time
}

func (l LastProducts) GetDate() string {
	return fmt.Sprintf("%d-%s-%d", l.Date.Day(), l.Date.Month().String(), l.Date.Year())
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

func (s *Store) GetUsers() ([]User, error) {
	var list []User
	err := s.conn.SelectContext(context.Background(), &list, `SELECT id, name FROM users`)
	if err != nil {
		return nil, fmt.Errorf("query err: %w", err)
	}

	return list, nil
}

func (s *Store) AddUser(name string, id int64) error {
	_, err := s.conn.ExecContext(context.Background(), `INSERT INTO users(id, name) VALUES($1, $2);`, id, name)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}

// GetProductList selects list of product from table for current user
func (s *Store) GetProductList(owner string, date *time.Time) ([]ProductList, error) {
	var err error
	list := []ProductList{}
	if date == nil {
		err = s.conn.SelectContext(context.Background(), &list, `SELECT pl.name, pl.time FROM product_list AS pl
														WHERE owner = $1
														ORDER BY pl.name DESC;`, owner)
	} else {
		err = s.conn.SelectContext(context.Background(), &list, `SELECT pl.name, pl.time  FROM product_list AS pl
														WHERE owner = $1 AND time <= $2
														ORDER BY pl.name DESC;`, owner, date)
	}
	if err != nil {
		return nil, fmt.Errorf("query err: %w", err)
	}

	return list, nil
}

// GetBuyList selects buy list from table for current user
func (s *Store) GetBuyList(owner string, date *time.Time) ([]BuyList, error) {
	var err error
	var list []BuyList

	if date == nil {
		err = s.conn.SelectContext(context.Background(), &list, `SELECT bl.name, bl.weight, bl.reminder FROM buy_list AS bl
														WHERE owner = $1;`, owner)
	} else {
		err = s.conn.SelectContext(context.Background(), &list, `SELECT bl.name, bl.weight, bl.reminder FROM buy_list AS bl
														WHERE owner = $1 AND reminder < $2 OR reminder = $2;`, owner, date)
	}
	if err != nil {
		return nil, fmt.Errorf("query err: %w", err)
	}

	return list, nil
}

// GetLastList selects LastList from table for current user
func (s *Store) GetLastList(owner string, from string, to string) ([]LastProducts, error) { // , time1 string, time2 string
	var list []LastProducts
	var err error

	if (from != "-1") && (to != "-1") {
		err = s.conn.SelectContext(context.Background(), &list, `SELECT ll.name, ll.status, ll.date FROM last_product AS ll
														WHERE owner = $1 AND ll.date >= $2 AND ll.date <= $3;`, owner, from, to)
	} else if (from == "-1") && (to != "-1") {
		err = s.conn.SelectContext(context.Background(), &list, `SELECT ll.name, ll.status, ll.date FROM last_product AS ll
														WHERE owner = $1 AND ll.date <= $2;`, owner, to)
	} else if (to == "-1") && (from != "-1") {
		err = s.conn.SelectContext(context.Background(), &list, `SELECT ll.name, ll.status, ll.date FROM last_product AS ll
														WHERE owner = $1 AND ll.date >= $2;`, owner, from)
	} else {
		err = s.conn.SelectContext(context.Background(), &list, `SELECT ll.name, ll.status, ll.date FROM last_product AS ll
														WHERE owner = $1;`, owner)
	}
	if err != nil {
		return nil, fmt.Errorf("query err: %w", err)
	}

	return list, nil
}

// AddProductToBuyList adds product to BuyList
func (s *Store) AddProductToBuyList(name string, weight string, reminder string, owner string) error {
	var err error
	if reminder != "-1" {
		_, err = s.conn.ExecContext(context.Background(), `INSERT INTO buy_list(name, weight, reminder, owner) 
														VALUES($1, $2, $3, $4);`, name, weight, reminder, owner)
	} else {
		data := strings.Split(time.Now().String(), " ")[0]
		_, err = s.conn.ExecContext(context.Background(), `INSERT INTO buy_list(name, weight, reminder, owner) 
														VALUES($1, $2, $3, $4);`, name, weight, data, owner)
	}
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}

// DeleteProductFromBuyList deletes product from BuyList
func (s *Store) DeleteProductFromBuyList(name string, owner string) error {
	var err error
	_, err = s.conn.ExecContext(context.Background(), `DELETE FROM buy_list
													WHERE name = $1 AND owner = $2;`, name, owner)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}

// AddProductToProductList adds product to productList
func (s *Store) AddProductToProductList(name string, data string, owner string) error {
	var err error
	_, err = s.conn.ExecContext(context.Background(), `INSERT INTO product_list(name, time, owner)
													VALUES($1, $2, $3);`, name, data, owner)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}

// ChangeProductFromProductList changes product to productList
func (s *Store) ChangeProductFromProductList(name string, data string, owner string) error {
	var err error
	_, err = s.conn.ExecContext(context.Background(), `UPDATE product_list
													SET time = $1
													WHERE name = $2 AND owner = $3;`, data, name, owner)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}

// DeleteProductFromProductList deletes product from productList
func (s *Store) DeleteProductFromProductList(name string, owner string) error {
	var err error
	_, err = s.conn.ExecContext(context.Background(), `DELETE FROM product_list
													WHERE name = $1 AND owner = $2;`, name, owner)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}

// AddProductToLastList adds product to lastList
func (s *Store) AddProductToLastList(name string, status bool, owner string) error {
	var err error
	_, err = s.conn.ExecContext(context.Background(), `INSERT INTO last_product(name, owner, status, date)
													VALUES($1, $2, $3, $4);`, name, owner, status, time.Now())
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}

	return nil
}
