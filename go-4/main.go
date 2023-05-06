package main

import (
	"fmt"

	models "pack/api/models"
	route "pack/api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "nome 1", Historia: "historia 1"},
		{Id: 2, Nome: "nome 2", Historia: "historia 2"},
	}

	fmt.Println("Iniciando o servidore Rest com Go")
	route.HandleRequest()
}
