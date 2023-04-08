package contas

import "fmt"

// Sempre que quiser deixar visivel os atributos e métodos só começar com a primeira letra maiuscula
type ContaCorrente struct { // estrutura == classe
	Titular    string
	NumAgencia int
	NumConta   int
	Saldo      float64
}

func (c *ContaCorrente) Sacar(saque float64) string {
	positivo := saque > 0 && saque < c.Saldo
	if positivo {
		c.Saldo -= saque
		return "saque feito com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func (c *ContaCorrente) Depositar(dep float64) (string, float64) {
	if dep > 0 {
		c.Saldo += dep
		return "Deposito feito com sucesso!", c.Saldo
	} else {
		return "Não é permitido deposito de valores negativos!", c.Saldo
	}

}

func (c *ContaCorrente) Transferir(valor float64, dest *ContaCorrente) bool {
	if valor < 0 {
		fmt.Println("Não é permitido valores negativos!")
		return false
	} else {
		if valor < c.Saldo {
			c.Saldo -= valor
			dest.Saldo += valor

			return true
		} else {
			fmt.Println("Saldo insuficiente!")
			return false
		}
	}
}
