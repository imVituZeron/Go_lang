package main

import (
	"fmt"

	cli "pack.com/go-2/pkg/clientes" // o "c" é uma apelido para os import
	c "pack.com/go-2/pkg/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	clienteVitor := cli.Titular{Nome: "Vitor", CPF: "23165498744", Profissao: "Analista DevOps"}
	contaCorrDoVitor := c.ContaCorrente{Titular: clienteVitor, NumAgencia: 123, NumConta: 123456}
	contaPoupDoVitor := c.ContaPoupanca{Titular: clienteVitor, NumAgencia: 123, NumConta: 123654}

	contaCorrDoVitor.Depositar(100)

	PagarBoleto(&contaCorrDoVitor, 60)

	contaPoupDoVitor.Depositar(100)
	PagarBoleto(&contaPoupDoVitor, 25)

	fmt.Println(contaCorrDoVitor.Visulizar())
	fmt.Println(contaPoupDoVitor.Visulizar())
}
