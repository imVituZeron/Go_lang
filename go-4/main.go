package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func HandleRequest() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	fmt.Println("Iniciando o servidore Rest com Go")
	HandleRequest()
}
