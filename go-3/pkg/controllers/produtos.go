package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	prod "pack.com/loja/pkg/produtos"
)

var temp = template.Must(template.ParseGlob("./src/templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProduct := prod.BringAllProduct()
	temp.ExecuteTemplate(w, "index", allProduct)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConver, errprec := strconv.ParseFloat(preco, 64)
		if errprec != nil {
			fmt.Println("houve um erro!: ", errprec)
		}

		quantidadeConvert, errquant := strconv.Atoi(quantidade)
		if errquant != nil {
			fmt.Println("houve um erro!: ", errquant)
		}

		prod.CreateNewProduct(nome, descricao, precoConver, quantidadeConvert)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	prod.DeleteProduct(idProduct)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := prod.EditaProduto(idProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}
