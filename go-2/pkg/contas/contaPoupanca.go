package contas

import (
	cli "pack.com/go-2/pkg/clientes"
)

type ContaPoupanca struct {
	Titular                        cli.Titular
	NumAgencia, NumConta, Operacao int
	saldo                          float64
}

func (c *ContaPoupanca) Sacar(saque float64) string {
	positivo := saque > 0 && saque < c.saldo
	if positivo {
		c.saldo -= saque
		return "saque feito com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func (c *ContaPoupanca) Depositar(dep float64) (string, float64) {
	if dep > 0 {
		c.saldo += dep
		return "Deposito feito com sucesso!", c.saldo
	} else {
		return "Não é permitido deposito de valores negativos!", c.saldo
	}

}

func (c *ContaPoupanca) Visulizar() float64 {
	return c.saldo
}
