package conn

import (
	"context"
	"fmt"
	"time"

	// "github.com/golang-migrate/migrate/v4"
	// "github.com/golang-migrate/migrate/v4/database/postgres"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
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
	Time time.Time
}
type BuyList struct {
	Name     string
	Weight   float32
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
	rows, err := s.conn.Query(context.Background(), `SELECT pl.name, pl.time FROM productList AS pl
														WHERE owner =`+owner)
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
	rows, err := s.conn.Query(context.Background(), `SELECT bl.name, bl.weight, bl.reminder FROM buyList AS bl
														WHERE owner =`+owner)
	if err != nil {
		return nil, fmt.Errorf("query err: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		p := BuyList{}
		if err := rows.Scan(&p.Name, &p.Weight, *&p.Reminder); err != nil {
			return nil, fmt.Errorf("scan err: %v\n", err)
		}
		list = append(list, p)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("scan err: %v\n", err)
	}
	return list, nil
}

// GetLastList selects LastList from table for current user
func (s *Store) GetLastList(owner string, time1 string, time2 string) ([]LastProducts, error) {
	list := make([]LastProducts, 0)
	var rows pgx.Rows
	var err error
	if time1 != "-1" && time2 != "-1" {
		rows, err = s.conn.Query(context.Background(), `SELECT ll.name, ll.status, ll.date FROM buyList AS ll
														WHERE owner =`+owner+` ll.date>=`+time1+` AND ll.date<=`+time2)
	} else if time1 == "-1" {
		rows, err = s.conn.Query(context.Background(), `SELECT ll.name, ll.status, ll.date FROM buyList AS ll
														WHERE owner =`+owner+` ll.date<=`+time2)
	} else if time2 == "-1" {
		rows, err = s.conn.Query(context.Background(), `SELECT ll.name, ll.status, ll.date FROM buyList AS ll
														WHERE owner =`+owner+` ll.date>=`+time1)
	} else {
		rows, err = s.conn.Query(context.Background(), `SELECT ll.name, ll.status, ll.date FROM buyList AS ll
														WHERE owner =`+owner)
	}

	if err != nil {
		return nil, fmt.Errorf("query err: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		p := LastProducts{}
		if err := rows.Scan(&p.Name, &p.Status, *&p.Date); err != nil {
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
	var err error
	if reminder != "-1" {
		_, err = s.conn.Query(context.Background(), `INSERT INTO buyList(name, weight, data, owner) 
														VALUE(`+name+`, `+weight+`, `+reminder+`, `+owner+`)`)
	} else {
		_, err = s.conn.Query(context.Background(), `INSERT INTO buyList(name, weight, owner) 
														VALUE(`+name+`, `+weight+`, `+owner+`)`)
	}

	if err != nil {
		return fmt.Errorf("found err: %w", err)
	}
	return nil
}

// migrates do migration of db
// func migrates(connString string) error {
// 	db, err := sql.Open("postgres", connString)
// 	if err != nil {
// 		panic(err)
// 	}
// 	driver, err := postgres.WithInstance(db, &postgres.Config{
// 		DatabaseName: "mitiushin",
// 		SchemaName:   "public",
// 	})
// 	if err != nil {
// 		return fmt.Errorf("migrate: %w", err)
// 	}
// 	m, err := migrate.NewWithDatabaseInstance("file://migrations", "mitiushin", driver)
// 	if err != nil {
// 		return fmt.Errorf("migrate: %w", err)
// 	}
// 	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
// 		return fmt.Errorf("migrate: %w", err)
// 	}
// 	return nil
// }
