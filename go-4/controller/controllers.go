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

func CreateNewPersona(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)

	database.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade)
}

func DeletePersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade

	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func EditPersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)

	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
