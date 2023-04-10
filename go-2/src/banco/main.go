package main

import (
	"fmt"

	cli "pack.com/go-2/pkg/clientes" // o "c" é uma apelido para os import
	c "pack.com/go-2/pkg/contas"
)

func main() {
	clienteVitor := cli.Titular{Nome: "Vitor", CPF: "23165498744", Profissao: "Analista DevOps"}
	contaDoVitor := c.ContaCorrente{Titular: clienteVitor, NumAgencia: 123, NumConta: 123456}

	contaDoVitor.Depositar(1000)
	fmt.Println(contaDoVitor.Visulizar())

}
