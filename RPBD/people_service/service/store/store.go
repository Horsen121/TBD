package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
)

type Store struct {
	conn *pgx.Conn
}
type People struct {
	ID   int
	Name string
}

// NewStore creates new database connection
func NewStore(connString string) (*Store, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		panic(err)
	}

	migrates(connString)

	return &Store{
		conn: conn,
	}, nil
}

// ListPeople selects list of people from table
func (s *Store) ListPeople() ([]People, error) {
	list := make([]People, 0)
	rows, err := s.conn.Query(context.Background(), "SELECT p.id, p.name FROM people AS p")
	if err != nil {
		return nil, fmt.Errorf("query err: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		p := People{}
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, fmt.Errorf("scan err: %v\n", err)
		}
		list = append(list, p)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("scan err: %v\n", err)
	}
	return list, nil
}

// GetPeopleByID selects person from table by id
func (s *Store) GetPeopleByID(id string) (People, error) {
	p := People{}
	err := s.conn.QueryRow(context.Background(), `SELECT p.id, p.name FROM people AS p 
													WHERE id=`+id).Scan(&p.ID, &p.Name)
	if err != nil {
		return p, fmt.Errorf("found err: %w", err)
	}
	return p, nil
}

// migrates do migration of db
func migrates(connString string) error {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{
		DatabaseName: "mitiushin",
		SchemaName:   "public",
	})
	if err != nil {
		return fmt.Errorf("migrate: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "mitiushin", driver)
	if err != nil {
		return fmt.Errorf("migrate: %w", err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migrate: %w", err)
	}
	return nil
}
