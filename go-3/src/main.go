package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
	prod "pack.com/loja/pkg/produtos"
)

var temp = template.Must(template.ParseGlob("./templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":9001", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectDb()
	selectProduct, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}

	p := prod.Produto{}
	produtos := []prod.Produto{}

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

	temp.ExecuteTemplate(w, "index", produtos)
	defer db.Close()
}

func conectDb() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
