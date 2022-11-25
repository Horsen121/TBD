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
func (s *Store) GetProductList(owner string, date string) ([]ProductList, error) {
	var rows pgx.Rows
	var err error
	list := make([]ProductList, 0)

	if date == "-1" {
		rows, err = s.conn.Query(context.Background(), `SELECT pl.name FROM productlist AS pl
														WHERE owner = $1;`, owner)
	} else {
		rows, err = s.conn.Query(context.Background(), `SELECT pl.name FROM productlist AS pl
														WHERE owner = $1 AND time = $2;`, owner, date)
	}
	if err != nil {
		return nil, fmt.Errorf("query err: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		p := ProductList{}
		if err := rows.Scan(&p.Name); err != nil {
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
func (s *Store) GetBuyList(owner string, date string) ([]BuyList, error) {
	var rows pgx.Rows
	var err error
	list := make([]BuyList, 0)

	if date == "-1" {
		rows, err = s.conn.Query(context.Background(), `SELECT bl.name, bl.weight FROM buylist AS bl
														WHERE owner = $1;`, owner)
	} else {
		rows, err = s.conn.Query(context.Background(), `SELECT bl.name, bl.weight FROM buylist AS bl
														WHERE owner = $1 AND reminder = $2;`, owner, date)
	}
	if err != nil {
		return nil, fmt.Errorf("query err: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		p := BuyList{}
		if err := rows.Scan(&p.Name, &p.Weight); err != nil {
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
			return nil, fmt.Errorf("1 scan err: %v\n", err)
		}
		list = append(list, p)
	}
	// if rows.Err() != nil {
	// 	return nil, fmt.Errorf("2 scan err: %v\n", err)
	// }
	return list, nil
}

// AddProductToBuyList adds product to BuyList
func (s *Store) AddProductToBuyList(name string, weight string, reminder string, owner string) error {
	var err error
	if reminder != "-1" {
		_, err = s.conn.Query(context.Background(), `INSERT INTO buylist(name, weight, reminder, owner) 
														VALUES($1, $2, $3, $4);`, name, weight, reminder, owner)
	} else {
		_, err = s.conn.Query(context.Background(), `INSERT INTO buylist(name, weight, owner) 
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
	_, err = s.conn.Query(context.Background(), `DELETE FROM buylist
													WHERE name = $1 AND owner = $2;`, name, owner)

	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}
	return nil
}

// AddProductToProductList adds product to productList
func (s *Store) AddProductToProductList(name string, data string, owner string) error {
	var err error
	_, err = s.conn.Query(context.Background(), `INSERT INTO productlist(name, data, owner)
													VALUES($1, $2, $3);`, name, data, owner)

	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}
	return nil
}

// ChangeProductFromProductList changes product to productList
func (s *Store) ChangeProductFromProductList(name string, data string, owner string) error {
	var err error
	_, err = s.conn.Query(context.Background(), `UPDATE productlist
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
	_, err = s.conn.Query(context.Background(), `DELETE FROM productlist
													WHERE name = $1 AND owner = $2;`, name, owner)

	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}
	return nil
}

// AddProductToLastList adds product to lastList
func (s *Store) AddProductToLastList(name string, status string, owner string) error {
	var err error
	_, err = s.conn.Query(context.Background(), `INSERT INTO lastproduct(name, owner, status, data)
													VALUES($1, $2, $3);`, name, owner, status, time.Now())

	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}
	return nil
}
