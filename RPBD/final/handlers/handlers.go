package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Horsen121/TBD/RPBD/final/funcs"
	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()

	res := funcs.NewUser(r.Form.Get("user"), r.Form.Get("name"), r.Form.Get("surname"), r.Form.Get("login"),
		r.Form.Get("password"), r.Form.Get("priority") == "true", r.Form.Get("type") == "true")
	json.NewEncoder(w).Encode(res)
}

func Autorising(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := mux.Vars(r)

	res := funcs.Autorisation(p["login"], p["password"])
	json.NewEncoder(w).Encode(res)
}

func SetPriority(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// r.ParseForm()
	var res string

	user_id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Print(err.Error())
		res = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res = funcs.ChangeUsersPriority(user_id, mux.Vars(r)["priority"] == "true")
	json.NewEncoder(w).Encode(res)
}

func CreateSmena(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	var res string

	user_id, err := strconv.Atoi(r.Form.Get("user"))
	if err != nil {
		log.Print(err.Error())
		res = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	tmp := strings.Replace(r.Form.Get("start"), " ", "T", 1) + "Z"
	start, err := time.Parse(time.RFC3339, tmp)
	if err != nil {
		res = err.Error()
		log.Print(err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}
	tmp = strings.Replace(r.Form.Get("finish"), " ", "T", 1) + "Z"
	finish, err := time.Parse(time.RFC3339, tmp)
	if err != nil {
		log.Print(err.Error())
		res = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res = funcs.AddNewSmena(user_id, start, finish)
	json.NewEncoder(w).Encode(res)
}

func CreateChange(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	var res string

	user_id, err := strconv.Atoi(r.Form.Get("user"))
	if err != nil {
		log.Print(err.Error())
		res = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	smena_id, err := strconv.Atoi(r.Form.Get("smena"))
	if err != nil {
		log.Print(err.Error())
		res = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	tmp := strings.Replace(r.Form.Get("start"), " ", "T", 1) + "Z"
	start, err := time.Parse(time.RFC3339, tmp)
	if err != nil {
		res = err.Error()
		log.Print(err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}
	tmp = strings.Replace(r.Form.Get("finish"), " ", "T", 1) + "Z"
	finish, err := time.Parse(time.RFC3339, tmp)
	if err != nil {
		log.Print(err.Error())
		res = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res = funcs.AddChangeOffer(user_id, smena_id, start, finish)
	json.NewEncoder(w).Encode(res)
}

func GetListOfSmenaForUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res string
	user_id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Print(err.Error())
		res = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res = funcs.GetListOfSmena(user_id)
	json.NewEncoder(w).Encode(res)
}

func GetListOfSmena(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res string

	res = funcs.GetListOfSmena(0)
	json.NewEncoder(w).Encode(res)
}

func BlockUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res string

	user_id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Print(err.Error())
		res = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res = funcs.ChangeStatus(mux.Vars(r)["user"], user_id, false)
	json.NewEncoder(w).Encode(res)
}

func Ill(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	var res string

	user_id, err := strconv.Atoi(r.Form.Get("user"))
	if err != nil {
		log.Print(err.Error())
		res = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	coef, err := strconv.ParseFloat(r.Form.Get("coef"), 32)
	if err != nil {
		log.Print(err.Error())
		res = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	tmp := strings.Replace(r.Form.Get("start"), " ", "T", 1) + "Z"
	start, err := time.Parse(time.RFC3339, tmp)
	if err != nil {
		res = err.Error()
		log.Print(err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}
	tmp = strings.Replace(r.Form.Get("finish"), " ", "T", 1) + "Z"
	finish, err := time.Parse(time.RFC3339, tmp)
	if err != nil {
		log.Print(err.Error())
		res = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res = funcs.Ill(user_id, start, finish, coef)
	json.NewEncoder(w).Encode(res)
}

func Change(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	var res string

	user_id, err := strconv.Atoi(r.Form.Get("user"))
	if err != nil {
		log.Print(err.Error())
		res = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	smena_id, err := strconv.Atoi(r.Form.Get("smena"))
	if err != nil {
		log.Print(err.Error())
		res = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	ill, err := strconv.ParseBool(r.Form.Get("ill"))
	if err != nil {
		log.Print(err.Error())
		res = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	tmp := strings.Replace(r.Form.Get("start"), " ", "T", 1) + "Z"
	start, err := time.Parse(time.RFC3339, tmp)
	if err != nil {
		res = err.Error()
		log.Print(err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}
	tmp = strings.Replace(r.Form.Get("finish"), " ", "T", 1) + "Z"
	finish, err := time.Parse(time.RFC3339, tmp)
	if err != nil {
		log.Print(err.Error())
		res = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res = funcs.ChangingSmena(smena_id, user_id, start, finish, ill)
	json.NewEncoder(w).Encode(res)
}
