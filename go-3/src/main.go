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
	db := conectDb()
	http.HandleFunc("/", index)
	http.ListenAndServe(":9001", nil)
	defer db.Close()
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []prod.Produto{
		{Nome: "Camiseta", Desc: "Bem Bonita", Preco: 20.00, Quantidade: 10},
		{Nome: "Notebook", Desc: "Muito rápido", Preco: 1999, Quantidade: 2},
		{Nome: "Caneca", Desc: "beber nela", Preco: 29.00, Quantidade: 2},
		{Nome: "Teclado", Desc: "Mecânico", Preco: 290.00, Quantidade: 3},
	}

	temp.ExecuteTemplate(w, "index", produtos)
}
func conectDb() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
