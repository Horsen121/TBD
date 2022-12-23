package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Horsen121/TBD/RPBD/final/funcs"
	"github.com/Horsen121/TBD/RPBD/final/handlers"
	"github.com/Horsen121/TBD/RPBD/final/store/conn"
	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
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
	funcs.Start(s)

	r := mux.NewRouter()
	r.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/user/auto/{login}&{password}", handlers.Autorising).Methods("GET")
	r.HandleFunc("/user/priority/{id}&{priority}", handlers.SetPriority).Methods("PUT")
	r.HandleFunc("/smena", handlers.CreateSmena).Methods("POST")
	r.HandleFunc("/change/create", handlers.CreateChange).Methods("POST")
	r.HandleFunc("/smena/{id}", handlers.GetListOfSmenaForUser).Methods("GET")
	r.HandleFunc("/smena", handlers.GetListOfSmena).Methods("GET")
	r.HandleFunc("/user/block/{user}&{id}", handlers.BlockUser).Methods("PUT")
	r.HandleFunc("/ill", handlers.Ill).Methods("POST")
	r.HandleFunc("/change", handlers.Change).Methods("POST")

	log.Print("Server started")
	log.Print(http.ListenAndServe(":8080", r))
}
