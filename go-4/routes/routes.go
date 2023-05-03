package routes

import (
	"log"
	"net/http"

	control "pack/api/controller"
)

func HandleRequest() {
	http.HandleFunc("/", control.Home)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
