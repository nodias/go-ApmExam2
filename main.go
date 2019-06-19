package main

import (
	"log"
	"net/http"
	"strconv"

	"go.elastic.co/apm/module/apmgorilla"

	"github.com/gorilla/mux"
	"github.com/nodias/go-ApmExam2/api"
	"github.com/urfave/negroni"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", GetUser)
	router.Use(apmgorilla.Middleware())
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":7002")
}

func GetUsers(w http.ResponseWriter, req *http.Request) {
	data, err := api.ApiGetUsers()
	if err != nil {
		log.Printf("GetUsers : %s", err)
		data = []byte(err.Error())
	}
	w.Write(data)
	return
}

func GetUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	uid_str := params["id"]
	uid, err := strconv.Atoi(uid_str)
	if err != nil {
		log.Fatalf("GetUser : %s", err)
	}
	data, err := api.ApiGetUser(uid)
	if err != nil {
		log.Printf("GetUser : %s", err)
		data = []byte(err.Error())
	}
	w.Write(data)
	return
}
