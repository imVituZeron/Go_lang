package produtos

import (
	"fmt"

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

		p.Id = id
		p.Nome = nome
		p.Desc = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CreateNewProduct(nome string, desc string, preco float64, quant int) {
	db := data.ConectDb()
	query := "INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES ($1,$2,$3,$4);"

	insertDataInBase, err := db.Prepare(query)
	if err != nil {
		fmt.Println("ocorreu um erro: ", err)
	}

	insertDataInBase.Exec(nome, desc, preco, quant)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := data.ConectDb()

	deleteTheProduct := "DELETE FROM produtos WHERE id=$1"
	query, err := db.Prepare(deleteTheProduct)
	if err != nil {
		panic(err.Error())
	}

	query.Exec(id) // executando o query
	db.Close()
}
