package main

import "fmt"

type ContaCorrente struct { // estrutura == classe
	titular    string
	numAgencia int
	numConta   int
	saldo      float64
}

func main() {
	//conta1 := ContaCorrente{titular: "guilherm", numAgencia: 589, numConta: 5, saldo: 123.45} forma não obriga passar todos os params
	//conta1 := ContaCorrente{"guilherm", 589, 5, 123.45} forma obriga passar todos os params
	conta1 := ContaCorrente{}
	conta1.titular = "Guilherme"
	conta1.numAgencia = 568
	conta1.numConta = 05
	conta1.saldo = 123.35

	fmt.Println(conta1.titular)
	fmt.Println(conta1.numAgencia)
	fmt.Println(conta1.numConta)
	fmt.Println(conta1.saldo)
}
