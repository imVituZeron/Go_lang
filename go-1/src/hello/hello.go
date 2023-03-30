package main

import (
	"fmt"
)

func main() {
	name := "Vitor"
	version := 1.1

	fmt.Println("Ola, Sr", name)
	fmt.Println("Esta na versão ", version)

	fmt.Println("1 - Iniciar MOnitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")

	var command int
	fmt.Scan(&command)
	fmt.Println("O comando escolhido foi ", command)
}
