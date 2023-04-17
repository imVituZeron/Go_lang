package produtos

import (
	data "pack.com/loja/pkg/db"
)

type Produto struct {
	Id         int
	Nome       string
	Desc       string
	Preco      float64
	Quantidade int
}

func BringAllProduct() []Produto {
	db := data.ConectDb()
	selectProduct, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProduct.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProduct.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Desc = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}
