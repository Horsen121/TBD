package conn

import (
	"context"

	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func CreateDB() (testcontainers.Container, *sqlx.DB) { // testcontainers.Container,
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"4321/tcp"},
		WaitingFor:   wait.ForListeningPort("4321/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       "test",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_USER":     "postgres",
		},
	}

	dbContainer, err := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})
	if err != nil {
		panic(err)
	}

	host, err := dbContainer.Host(context.Background())
	if err != nil {
		panic(err)
	}
	port, err := dbContainer.MappedPort(context.Background(), "4321")
	if err != nil {
		panic(err)
	}

	connString := fmt.Sprintf("postgres://postgres:postgres@%v:%v/testdb?sslmode=disable", host, port.Port())
	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	store, err := NewStore(connString)
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:./migrations/",
		"postgres", driver)
	if err != nil {
		panic(err)
	}
	m.Up()

	return dbContainer, store.conn
}

func TestAddUser(t *testing.T) {
	container, db := CreateDB()
	defer container.Terminate(context.Background())

	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	// store, err := NewStore(connString)
	// if err != nil {
	// 	panic(err)
	// }

	store := Store{
		conn: db,
	}

	if err := store.AddUser("testUser", 123456789); err != nil {
		t.Error(err)
	}
}

func TestGetUsers(t *testing.T) {
	container, db := CreateDB()
	defer container.Terminate(context.Background())

	store := Store{
		conn: db,
	}

	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	// store, err := NewStore(connString)
	// if err != nil {
	// 	panic(err)
	// }

	_, err := store.GetUsers()
	if err != nil {
		panic(err)
	}
}

func TestGetProductList1(t *testing.T) {
	container, db := CreateDB()
	defer container.Terminate(context.Background())

	store := Store{
		conn: db,
	}

	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	// store, err := NewStore(connString)
	// if err != nil {
	// 	panic(err)
	// }

	_, err := store.GetProductList("Horsen17", nil)
	if err != nil {
		t.Error(err)
	}
}

func TestGetProductList2(t *testing.T) {
	container, db := CreateDB()
	defer container.Terminate(context.Background())

	store := Store{
		conn: db,
	}

	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	// store, err := NewStore(connString)
	// if err != nil {
	// 	panic(err)
	// }

	// data := strings.Split(time.Now().String(), " ")[0]
	data := time.Now()
	_, err := store.GetProductList("Horsen17", &data)
	if err != nil {
		t.Error(err)
	}
}

func TestGetBuyList1(t *testing.T) {
	container, db := CreateDB()
	defer container.Terminate(context.Background())

	store := Store{
		conn: db,
	}

	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	// store, err := NewStore(connString)
	// if err != nil {
	// 	panic(err)
	// }

	_, err := store.GetBuyList("Horsen17", nil)
	if err != nil {
		t.Error(err)
	}
}

func TestGetBuyList2(t *testing.T) {
	container, db := CreateDB()
	defer container.Terminate(context.Background())

	store := Store{
		conn: db,
	}

	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	// store, err := NewStore(connString)
	// if err != nil {
	// 	panic(err)
	// }

	time := time.Now()
	_, err := store.GetBuyList("Horsen17", &time)
	if err != nil {
		t.Error(err)
	}
}

func TestGetLastList1(t *testing.T) {
	container, db := CreateDB()
	defer container.Terminate(context.Background())

	store := Store{
		conn: db,
	}

	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	// store, err := NewStore(connString)
	// if err != nil {
	// 	panic(err)
	// }

	data1 := strings.Split(time.Now().String(), " ")[0]
	data2 := strings.Split(time.Now().String(), " ")[0]
	_, err := store.GetLastList("Horsen17", data1, data2)
	if err != nil {
		t.Error(err)
	}
}

func TestGetLastList2(t *testing.T) {
	container, db := CreateDB()
	defer container.Terminate(context.Background())

	store := Store{
		conn: db,
	}

	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	// store, err := NewStore(connString)
	// if err != nil {
	// 	panic(err)
	// }

	data2 := strings.Split(time.Now().String(), " ")[0]
	_, err := store.GetLastList("Horsen17", "-1", data2)
	if err != nil {
		t.Error(err)
	}
}

func TestGetLastList3(t *testing.T) {
	container, db := CreateDB()
	defer container.Terminate(context.Background())

	store := Store{
		conn: db,
	}

	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	// store, err := NewStore(connString)
	// if err != nil {
	// 	panic(err)
	// }

	data1 := strings.Split(time.Now().String(), " ")[0]
	_, err := store.GetLastList("Horsen17", data1, "-1")
	if err != nil {
		t.Error(err)
	}
}

func TestGetLastList4(t *testing.T) {
	container, db := CreateDB()
	defer container.Terminate(context.Background())

	store := Store{
		conn: db,
	}

	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	// store, err := NewStore(connString)
	// if err != nil {
	// 	panic(err)
	// }

	_, err := store.GetLastList("Horsen17", "-1", "-1")
	if err != nil {
		t.Error(err)
	}
}

func TestAddProductToBuyList1(t *testing.T) {
	container, db := CreateDB()
	defer container.Terminate(context.Background())

	store := Store{
		conn: db,
	}

	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	// store, err := NewStore(connString)
	// if err != nil {
	// 	panic(err)
	// }

	data := strings.Split(time.Now().String(), " ")[0]
	err := store.AddProductToBuyList("test", "42", data, "Horsen17")
	if err != nil {
		t.Error(err)
	}
}

func TestAddProductToBuyList2(t *testing.T) {
	container, db := CreateDB()
	defer container.Terminate(context.Background())

	store := Store{
		conn: db,
	}

	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	// store, err := NewStore(connString)
	// if err != nil {
	// 	panic(err)
	// }

	err := store.AddProductToBuyList("test", "42", "-1", "Horsen17")
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteProductFromBuyList(t *testing.T) {
	container, db := CreateDB()
	defer container.Terminate(context.Background())

	store := Store{
		conn: db,
	}

	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	// store, err := NewStore(connString)
	// if err != nil {
	// 	panic(err)
	// }

	err := store.DeleteProductFromProductList("test", "Horsen17")
	if err != nil {
		t.Error(err)
	}
}

func TestAddProductToProductList(t *testing.T) {
	container, db := CreateDB()
	defer container.Terminate(context.Background())

	store := Store{
		conn: db,
	}

	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	// store, err := NewStore(connString)
	// if err != nil {
	// 	panic(err)
	// }

	data := strings.Split(time.Now().String(), " ")[0]
	err := store.AddProductToProductList("test", data, "Horsen17")
	if err != nil {
		t.Error(err)
	}
}

func TestChangeProductFromProductList(t *testing.T) {
	container, db := CreateDB()
	defer container.Terminate(context.Background())

	store := Store{
		conn: db,
	}

	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	// store, err := NewStore(connString)
	// if err != nil {
	// 	panic(err)
	// }

	data := strings.Split(time.Now().String(), " ")[0]
	err := store.ChangeProductFromProductList("test", data, "Horsen17")
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteProductFromProductList(t *testing.T) {
	container, db := CreateDB()
	defer container.Terminate(context.Background())

	store := Store{
		conn: db,
	}

	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	// store, err := NewStore(connString)
	// if err != nil {
	// 	panic(err)
	// }

	err := store.DeleteProductFromProductList("test", "Horsen17")
	if err != nil {
		t.Error(err)
	}
}

func TestAddProductToLastList(t *testing.T) {
	container, db := TestCreateDB()
	defer container.Terminate(context.Background())

	store := Store{
		conn: db,
	}

	// connString := fmt.Sprintf("postgres://mitiushin:PgDmnANIME10@95.217.232.188:7777/mitiushin")
	// store, err := NewStore(connString)
	// if err != nil {
	// 	panic(err)
	// }

	err := store.AddProductToLastList("test", true, "Horsen17")
	if err != nil {
		t.Error(err)
	}
}
