package controllers

import (
	"net/http"
	"text/template"

	prod "pack.com/loja/pkg/produtos"
)

var temp = template.Must(template.ParseGlob("./templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProduct := prod.BringAllProduct()
	temp.ExecuteTemplate(w, "index", allProduct)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}
