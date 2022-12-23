package service

import (
	"fmt"
	"log"
	"os"

	"github.com/Horsen121/TBD/RPBD/final/store/conn"
	"github.com/joho/godotenv"
)

func Start() {
	if err := godotenv.Load(); err != nil {
		log.Print(".env file not found")
		panic(err)
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DBNAME"), os.Getenv("DBPASSWORD"),
		os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"))
	s, err := conn.NewStore(connStr)
	if err != nil {
		log.Panic(err)
	}

	for {

	}
}
