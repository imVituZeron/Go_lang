package contas

import (
	"fmt"

	c "pack.com/go-2/pkg/clientes"
)

// Sempre que quiser deixar visivel os atributos e métodos só começar com a primeira letra maiuscula
type ContaCorrente struct { // estrutura == classe
	Titular              c.Titular
	NumAgencia, NumConta int
	saldo                float64
}

func (c *ContaCorrente) Sacar(saque float64) string {
	positivo := saque > 0 && saque < c.saldo
	if positivo {
		c.saldo -= saque
		return "saque feito com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func (c *ContaCorrente) Depositar(dep float64) (string, float64) {
	if dep > 0 {
		c.saldo += dep
		return "Deposito feito com sucesso!", c.saldo
	} else {
		return "Não é permitido deposito de valores negativos!", c.saldo
	}

}

func (c *ContaCorrente) Transferir(valor float64, dest *ContaCorrente) bool {
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

func (c *ContaCorrente) Visulizar() float64 {
	return c.saldo
}
