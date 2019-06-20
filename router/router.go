package router

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nodias/go-ApmExam2/api"
	"go.elastic.co/apm/module/apmgorilla"
)

func NewRouter() *mux.Router {
	return router()
}

func router() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/userInfo/{id}", getUserInfoHandler)
	router.HandleFunc("/users", getUsersHandler).Methods("GET")
	router.HandleFunc("/user/{id}", getUserHandler)
	router.HandleFunc("/hello/{name}", helloHandler)
	router.Use(apmgorilla.Middleware())
	return
}
func getUserInfoHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	data, err := api.ApiGetUserInfo(req.Context(), id)
	if err != nil {
		log.Printf("GetUser : %s", err)
		data = []byte(err.Error())
	}
	w.Write(data)
	return
}
func helloHandler(w http.ResponseWriter, req *http.Request) {
	name, err := fmt.Fprintf(w, "Hello, %s!\n", mux.Vars(req)["name"])
	if err != nil {
		log.Println(err)
	}
	data, err := api.UpdateRequestCount(req.Context(), name)
	if err != nil {
		log.Printf("GetUsers : %s", err)
		data = []byte(err.Error())
	}
	w.Write(data)
	return
}

func getUsersHandler(w http.ResponseWriter, req *http.Request) {
	data, err := api.ApiGetUsers(req.Context())
	if err != nil {
		log.Printf("GetUsers : %s", err)
		data = []byte(err.Error())
	}
	w.Write(data)
	return
}

func getUserHandler(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	uid_str := params["id"]
	uid, err := strconv.Atoi(uid_str)
	if err != nil {
		log.Fatalf("GetUser : %s", err)
	}
	data, err := api.ApiGetUser(req.Context(), uid)
	if err != nil {
		log.Printf("GetUser : %s", err)
		data = []byte(err.Error())
	}
	w.Write(data)
	return
}
