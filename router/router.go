package router

import (
	"log"
	"net/http"

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
