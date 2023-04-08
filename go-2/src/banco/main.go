package main

import (
	"fmt"

	c "pack.com/go-2/pkg/contas" // o "c" é uma apelido para os import
)

func main() {
	//conta1 := ContaCorrente{titular: "guilherm", numAgencia: 589, numConta: 5, saldo: 123.45} forma não obriga passar todos os params
	//conta1 := ContaCorrente{"guilherm", 589, 5, 123.45} forma obriga passar todos os params
	conta1 := c.ContaCorrente{Titular: "Vitor", Saldo: 1000}
	conta2 := c.ContaCorrente{Titular: "Sara", Saldo: 6000}

	fmt.Println(conta1, conta2)
	status := conta1.Transferir(100, &conta2)
	fmt.Println(status, conta1, conta2)

}
