package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pack/api/database"
	"pack/api/models"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ShowAllPersonas(w http.ResponseWriter, r *http.Request) {
	var P []models.Personalidade

	database.DB.Find(&P)
	json.NewEncoder(w).Encode(P)
}

func ShowPersona(w http.ResponseWriter, r *http.Request) {
	value := mux.Vars(r)
	id := value["id"]
	var P models.Personalidade

	database.DB.First(&P, id)
	json.NewEncoder(w).Encode(P)
}
