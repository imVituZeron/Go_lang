package main

import (
	"fmt"

	cli "pack.com/go-2/pkg/clientes" // o "c" é uma apelido para os import
	c "pack.com/go-2/pkg/contas"
)

func main() {
	clienteVitor := cli.Titular{Nome: "Vitor", CPF: "23165498744", Profissao: "Analista DevOps"}
	contaCorrDoVitor := c.ContaCorrente{Titular: clienteVitor, NumAgencia: 123, NumConta: 123456}
	contaPoupDoVitor := c.ContaPoupanca{Titular: clienteVitor, NumAgencia: 123, NumConta: 123654}

	contaCorrDoVitor.Depositar(1000)
	fmt.Println(contaCorrDoVitor.Visulizar())
	fmt.Println(contaCorrDoVitor)

	contaPoupDoVitor.Depositar(1000)
	fmt.Println(contaPoupDoVitor.Visulizar())
	fmt.Println(contaPoupDoVitor)

	contaCorrDoVitor.Sacar(55)
	fmt.Println(contaCorrDoVitor.Visulizar())

}
