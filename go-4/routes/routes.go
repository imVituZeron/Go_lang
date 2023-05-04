package routes

import (
	"log"
	"net/http"

	control "pack/api/controller"
)

func HandleRequest() {
	http.HandleFunc("/", control.Home)
	http.HandleFunc("/api/personalidades", control.ShowAllPersonas)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
