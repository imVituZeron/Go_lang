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
	r.HandleFunc("/api/personalidades", control.ShowAllPersonas).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", control.ShowPersona).Methods("Get")
	r.HandleFunc("/api/personalidades", control.CreateNewPersona).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", control.DeletePersona).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", control.EditPersona).Methods("Put")
	log.Fatal(http.ListenAndServe(":8081", r))
}
