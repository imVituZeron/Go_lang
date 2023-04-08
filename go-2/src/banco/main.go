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

func (c *ContaCorrente) depositar(dep float64) (string, float64) {
	if dep > 0 {
		c.saldo += dep
		return "Deposito feito com sucesso!", c.saldo
	} else {
		return "Não é permitido deposito de valores negativos!", c.saldo
	}

}

func (c *ContaCorrente) transferir(valor float64, dest *ContaCorrente) bool {
	if valor < 0 {
		fmt.Println("Não é permitido valores negativos!")
		return false
	} else {
		if valor < c.saldo {
			c.saldo -= valor
			dest.saldo += valor

			return true
		} else {
			fmt.Println("Saldo insuficiente!")
			return false
		}
	}
}

func main() {
	//conta1 := ContaCorrente{titular: "guilherm", numAgencia: 589, numConta: 5, saldo: 123.45} forma não obriga passar todos os params
	//conta1 := ContaCorrente{"guilherm", 589, 5, 123.45} forma obriga passar todos os params
	conta1 := ContaCorrente{titular: "Vitor", saldo: 1000}
	conta2 := ContaCorrente{titular: "Sara", saldo: 6000}

	fmt.Println(conta1, conta2)
	status := conta1.transferir(100, &conta2)
	fmt.Println(status, conta1, conta2)
}
