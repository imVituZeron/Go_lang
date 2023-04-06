package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	intro()
	registerLog("site-falso", false)
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
	//sites := []string{"https://www.alura.com.br", "https://random-status-code.herokuapp.com/", "https://www.caelum.com.br", "https://www.uol.com.br"}

	sites := readSitesFile()

	for i := 0; i < monitoramentos; i++ {
		for _, content := range sites {
			testSite(content)
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Println("---------------------------")
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "foi carregado com sucesso!")
		registerLog(site, true)
	} else {
		fmt.Println("O site", site, "esta com problemas. Status code:", resp.StatusCode)
		registerLog(site, false)
	}
}

func readSitesFile() []string {

	var sites []string
	file, err := os.Open("site.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(file) // pegando o arquivo e criando um leito para ela.
	for {
		linha, err_leitor := leitor.ReadString('\n') // pegando linha por linha do leitor separadas por \n
		linha = strings.TrimSpace(linha)             // tirando os \n nos finais das linhas

		sites = append(sites, linha)

		if err_leitor == io.EOF { //io.EOF indica que é o fim do arquivo, usado para fazer um break na leitura do arquivo
			break
		}

	}
	file.Close()

	return sites
}

func registerLog(site string, code bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(arquivo)
}
