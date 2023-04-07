package main

import "fmt"

type ContaCorrente struct { // estrutura == classe
	titular    string
	numAgencia int
	numConta   int
	saldo      float64
}

func main() {

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
