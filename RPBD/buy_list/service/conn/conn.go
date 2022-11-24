package conn

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
)

type Store struct {
	conn *pgx.Conn
}
type Product struct {
	Name string
}
type ProductList struct {
	Name string
	Time string
}
type BuyList struct {
	Name     string
	Weight   float64
	Reminder time.Time
}
type LastProducts struct {
	Name   string
	Status bool
	Date   time.Time
}

// NewStore creates new database connection
func NewStore(connString string) (*Store, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		panic(err)
	}

	// migrates(connString)

	return &Store{
		conn: conn,
	}, nil
}

// GetProductList selects list of product from table for current user
func (s *Store) GetProductList(owner string) ([]ProductList, error) {
	list := make([]ProductList, 0)
	rows, err := s.conn.Query(context.Background(), `SELECT pl.name, pl.time FROM productlist AS pl
														WHERE owner = $1;`, owner)
	if err != nil {
		return nil, fmt.Errorf("query err: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		p := ProductList{}
		if err := rows.Scan(&p.Name, &p.Time); err != nil {
			return nil, fmt.Errorf("scan err: %v\n", err)
		}
		list = append(list, p)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("scan err: %v\n", err)
	}
	return list, nil
}

// GetBuyList selects buy list from table for current user
func (s *Store) GetBuyList(owner string) ([]BuyList, error) {
	list := make([]BuyList, 0)
	rows, err := s.conn.Query(context.Background(), `SELECT bl.name, bl.weight, bl.reminder FROM buylist AS bl
														WHERE owner = $1;`, owner)
	if err != nil {
		return nil, fmt.Errorf("query err: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		p := BuyList{}
		if err := rows.Scan(&p.Name, &p.Weight, &p.Reminder); err != nil {
			return nil, fmt.Errorf("1 scan err: %v\n", err)
		}
		list = append(list, p)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("2 scan err: %v\n", err)
	}
	return list, nil
}

// GetLastList selects LastList from table for current user
func (s *Store) GetLastList(owner string, time1 string, time2 string) ([]LastProducts, error) {
	list := make([]LastProducts, 0)
	var rows pgx.Rows
	var err error
	if time1 != "-1" && time2 != "-1" {
		rows, err = s.conn.Query(context.Background(), `SELECT ll.name, ll.status, ll.date FROM lastproduct AS ll
														WHERE owner = $1 AND ll.date >= $2 AND ll.date <= $3;`, owner, time1, time2)
	} else if time1 == "-1" {
		rows, err = s.conn.Query(context.Background(), `SELECT ll.name, ll.status, ll.date FROM lastproduct AS ll
														WHERE owner = $1 AND ll.date <= $2;`, owner, time2)
	} else if time2 == "-1" {
		rows, err = s.conn.Query(context.Background(), `SELECT ll.name, ll.status, ll.date FROM lastproduct AS ll
														WHERE owner == $1 AND ll.date >= $2;`, owner, time1)
	} else {
		rows, err = s.conn.Query(context.Background(), `SELECT ll.name, ll.status, ll.date FROM lastproduct AS ll
														WHERE owner = $1;`, owner)
	}

	if err != nil {
		return nil, fmt.Errorf("query err: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		p := LastProducts{}
		if err := rows.Scan(&p.Name, &p.Status, &p.Date); err != nil {
			return nil, fmt.Errorf("scan err: %v\n", err)
		}
		list = append(list, p)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("scan err: %v\n", err)
	}
	return list, nil
}

// AddProductToBuyList adds product to BuyList
func (s *Store) AddProductToBuyList(name string, weight string, reminder string, owner string) error {
	_, err := s.conn.Query(context.Background(), `INSERT INTO buylist(name, weight, reminder, owner) 
													VALUES($1, $2, $3, $4);`, name, weight, reminder, owner)
	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}
	return nil
}
