package routes

import (
	"log"
	"net/http"

	control "pack/api/controller"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", control.Home)
	r.HandleFunc("/api/personalidades", control.ShowAllPersonas)
	log.Fatal(http.ListenAndServe(":8081", r))
}
