package routes

import (
	"net/http"

	control "pack.com/loja/pkg/controllers"
)

func LoadingRoutes() {
	http.HandleFunc("/", control.Index)
	http.HandleFunc("/new", control.New)
	http.HandleFunc("/insert", control.Insert)
}
