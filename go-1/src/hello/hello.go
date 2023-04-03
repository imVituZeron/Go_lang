package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	intro()
	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			initMonitoring()
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
}

func intro() {
	name := "Vitor"
	version := 1.1

	fmt.Println("Ola, Sr", name)
	fmt.Println("Este programa esta na versão", version)
}

func showMenu() {
	fmt.Println("---------------------------")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")
	fmt.Println("---------------------------")
}

func readCommand() int {
	var readedCommand int
	fmt.Scan(&readedCommand)
	fmt.Println("O comando escolhido foi", readedCommand)

	return readedCommand
}

func initMonitoring() {
	fmt.Println("---------------------------")
	fmt.Println("Monitorando.....")
	sites := []string{"https://www.alura.com.br", "https://random-status-code.herokuapp.com/", "https://www.caelum.com.br", "https://www.uol.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for _, content := range sites {
			testSite(content)
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Println("---------------------------")
}

func testSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O site", site, "esta com problemas. Status code:", resp.StatusCode)
	}
}
