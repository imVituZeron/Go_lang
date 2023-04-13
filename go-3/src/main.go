package main

import (
	"html/template"
	"net/http"

	prod "pack.com/loja/pkg/produtos"
)

var temp = template.Must(template.ParseGlob("./templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":9001", nil)
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
