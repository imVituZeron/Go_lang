package main

import (
	"fmt"

	route "pack/api/routes"
)

func main() {
	fmt.Println("Iniciando o servidore Rest com Go")
	route.HandleRequest()
}
