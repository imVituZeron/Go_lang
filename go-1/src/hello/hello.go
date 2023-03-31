package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	intro()
	showMenu()
	//	fmt.Println(devolveNomeIdade())

	command := readCommand()

	switch command {
	case 1:
		initMonitorin()
	case 2:
		fmt.Println("Exibindo logs")
	case 0:
		fmt.Println("Saindo ...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando!")
		os.Exit(-1)
	}
}

//func devolveNomeIdade() (string, int) {
//	name := "VItor"
//	idade := 23
//	return name, idade
//}

func intro() {
	name := "Vitor"
	version := 1.1

	fmt.Println("Ola, Sr", name)
	fmt.Println("Este programa esta na versão", version)
}

func showMenu() {
	fmt.Println("1 - Iniciar MOnitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")
}

func readCommand() int {
	var readedCommand int
	fmt.Scan(&readedCommand)
	fmt.Println("O comando escolhido foi", readedCommand)

	return readedCommand
}

func initMonitorin() {
	fmt.Println("Monitorando")
	site := "https://www.alura.com.br"
	fmt.Println(http.Get(site))
}
