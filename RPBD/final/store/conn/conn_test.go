package conn

import (
	// "context"

	"fmt"
	"time"

	// "strings"
	"testing"
	// "time"

	// "github.com/golang-migrate/migrate"
	// "github.com/golang-migrate/migrate/database/postgres"
	// "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	// "github.com/testcontainers/testcontainers-go"
	// "github.com/testcontainers/testcontainers-go/wait"
)

// func CreateDB() (testcontainers.Container, *sqlx.DB) { // testcontainers.Container,
// 	containerReq := testcontainers.ContainerRequest{
// 		Image:        "postgres:latest",
// 		ExposedPorts: []string{"4321/tcp"},
// 		WaitingFor:   wait.ForListeningPort("4321/tcp"),
// 		Env: map[string]string{
// 			"POSTGRES_DB":       "test",
// 			"POSTGRES_PASSWORD": "postgres",
// 			"POSTGRES_USER":     "postgres",
// 		},
// 	}
// 	dbContainer, err := testcontainers.GenericContainer(
// 		context.Background(),
// 		testcontainers.GenericContainerRequest{
// 			ContainerRequest: containerReq,
// 			Started:          true,
// 		})
// 	if err != nil {
// 		panic(err)
// 	}
// 	host, err := dbContainer.Host(context.Background())
// 	if err != nil {
// 		panic(err)
// 	}
// 	port, err := dbContainer.MappedPort(context.Background(), "4321")
// 	if err != nil {
// 		panic(err)
// 	}
// 	connString := fmt.Sprintf("postgres://postgres:postgres@%v:%v/testdb?sslmode=disable", host, port.Port())
// 	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
// 	store, err := NewStore(connString)
// 	if err != nil {
// 		panic(err)
// 	}
// 	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	m, err := migrate.NewWithDatabaseInstance(
// 		"file:./migrations/",
// 		"postgres", driver)
// 	if err != nil {
// 		panic(err)
// 	}
// 	m.Up()
// 	return dbContainer, store.conn
// }

func TestAddNewUser(t *testing.T) {
	// container, db := CreateDB()
	// defer container.Terminate(context.Background())

	connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	store, err := NewStore(connString)
	if err != nil {
		panic(err)
	}

	// store := Store{
	// 	conn: db,
	// }
	hash, err := bcrypt.GenerateFromPassword([]byte("pass"), 14)
	if err != nil {

	}
	if err = store.AddNewUser("test", "User", "login", string(hash), true, true, true); err != nil {
		t.Error(err)
	}
}

func TestGetUserByLogin1(t *testing.T) {
	// container, db := CreateDB()
	// defer container.Terminate(context.Background())

	connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	store, err := NewStore(connString)
	if err != nil {
		panic(err)
	}

	// store := Store{
	// 	conn: db,
	// }

	_, err = store.GetUserByLogin("login")
	if err != nil {
		t.Error(err)
	}
}

func TestGetUserByLogin2(t *testing.T) {
	// container, db := CreateDB()
	// defer container.Terminate(context.Background())

	connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	store, err := NewStore(connString)
	if err != nil {
		panic(err)
	}

	// store := Store{
	// 	conn: db,
	// }

	_, err = store.GetUserByLogin("login111")
	if err != nil {
		if err.Error() == "User not found" {
			return
		}
		t.Error(err)
	}
}

func TestChangePriority(t *testing.T) {
	// container, db := CreateDB()
	// defer container.Terminate(context.Background())

	connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	store, err := NewStore(connString)
	if err != nil {
		panic(err)
	}

	// store := Store{
	// 	conn: db,
	// }

	if err = store.ChangePriority(1, false); err != nil {
		t.Error(err)
	}
}

func TestAddSmena(t *testing.T) {
	// container, db := CreateDB()
	// defer container.Terminate(context.Background())

	connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	store, err := NewStore(connString)
	if err != nil {
		panic(err)
	}

	// store := Store{
	// 	conn: db,
	// }

	if err = store.AddSmena(1, time.Now(), time.Date(2022, 12, 22, 12, 45, 00, 0, time.Local)); err != nil {
		t.Error(err)
	}
}

func TestGetSmenaById(t *testing.T) {
	// container, db := CreateDB()
	// defer container.Terminate(context.Background())

	connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	store, err := NewStore(connString)
	if err != nil {
		panic(err)
	}

	// store := Store{
	// 	conn: db,
	// }

	if _, err = store.GetSmenaById(1); err != nil {
		if err.Error() == "Smena not found" {
			return
		}
		t.Error(err)
	}
}

func TestAddChange(t *testing.T) {
	// container, db := CreateDB()
	// defer container.Terminate(context.Background())

	connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	store, err := NewStore(connString)
	if err != nil {
		panic(err)
	}

	// store := Store{
	// 	conn: db,
	// }

	if err = store.AddChange(2, 1, 1, time.Date(2022, 12, 22, 12, 45, 00, 0, time.Local), time.Date(2022, 12, 23, 12, 45, 00, 0, time.Local)); err != nil {
		t.Error(err)
	}
}

func TestGetSmenaList1(t *testing.T) {
	// container, db := CreateDB()
	// defer container.Terminate(context.Background())

	connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	store, err := NewStore(connString)
	if err != nil {
		panic(err)
	}

	// store := Store{
	// 	conn: db,
	// }

	if _, err = store.GetSmenaList(1); err != nil {
		t.Error(err)
	}
}

func TestGetSmenaList2(t *testing.T) {
	// container, db := CreateDB()
	// defer container.Terminate(context.Background())

	connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	store, err := NewStore(connString)
	if err != nil {
		panic(err)
	}

	// store := Store{
	// 	conn: db,
	// }

	if _, err = store.GetSmenaList(0); err != nil {
		t.Error(err)
	}
}

func TestChangeUserStatus(t *testing.T) {
	// container, db := CreateDB()
	// defer container.Terminate(context.Background())

	connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	store, err := NewStore(connString)
	if err != nil {
		panic(err)
	}

	// store := Store{
	// 	conn: db,
	// }

	if err = store.ChangeUserStatus(1, false); err != nil {
		t.Error(err)
	}
}

func TestAddIll(t *testing.T) {
	// container, db := CreateDB()
	// defer container.Terminate(context.Background())

	connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	store, err := NewStore(connString)
	if err != nil {
		panic(err)
	}

	// store := Store{
	// 	conn: db,
	// }

	if err = store.AddIll(1, time.Date(2022, 12, 22, 12, 45, 00, 0, time.Local),
		time.Date(2022, 12, 27, 12, 45, 00, 0, time.Local), 1.2); err != nil {
		t.Error(err)
	}
}

func TestChangeSmena1(t *testing.T) {
	// container, db := CreateDB()
	// defer container.Terminate(context.Background())

	connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	store, err := NewStore(connString)
	if err != nil {
		panic(err)
	}

	// store := Store{
	// 	conn: db,
	// }

	if err = store.ChangeSmena(2, 1, time.Date(2022, 12, 22, 12, 45, 00, 0, time.Local),
		time.Date(2022, 12, 27, 12, 45, 00, 0, time.Local), time.Time{}, time.Time{}, true); err != nil {
		t.Error(err)
	}
}

func TestChangeSmena2(t *testing.T) {
	// container, db := CreateDB()
	// defer container.Terminate(context.Background())

	connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	store, err := NewStore(connString)
	if err != nil {
		panic(err)
	}

	// store := Store{
	// 	conn: db,
	// }

	if err = store.ChangeSmena(2, 1, time.Date(2022, 12, 22, 12, 45, 00, 0, time.Local),
		time.Date(2022, 12, 27, 12, 45, 00, 0, time.Local), time.Date(2022, 12, 26, 12, 45, 00, 0, time.Local),
		time.Date(2022, 12, 28, 12, 45, 00, 0, time.Local), false); err != nil {
		t.Error(err)
	}
}
