package main

import "fmt"

type ContaCorrente struct { // estrutura == classe
	titular    string
	numAgencia int
	numConta   int
	saldo      float64
}

func (c *ContaCorrente) sacar(saque float64) string {
	positivo := saque > 0 && saque < c.saldo

	if positivo {
		c.saldo -= saque
		return "saque feito com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func main() {
	//conta1 := ContaCorrente{titular: "guilherm", numAgencia: 589, numConta: 5, saldo: 123.45} forma não obriga passar todos os params
	//conta1 := ContaCorrente{"guilherm", 589, 5, 123.45} forma obriga passar todos os params
	conta1 := ContaCorrente{}
	conta1.numConta = 05
	conta1.saldo = 123.35

	fmt.Println(conta1.saldo)
	fmt.Println(conta1.sacar(60))
	fmt.Println(conta1.saldo)
}
