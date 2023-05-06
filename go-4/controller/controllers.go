package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pack/api/models"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ShowAllPersonas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}

func ShowPersona(w http.ResponseWriter, r *http.Request) {
	value := mux.Vars(r)
	id := value["id"]

	for _, value := range models.Personalidades {
		if strconv.Itoa(value.Id) == id {
			json.NewEncoder(w).Encode(value)
		}
	}
}
