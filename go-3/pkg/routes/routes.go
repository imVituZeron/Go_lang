package routes

import (
	"net/http"

	control "pack.com/loja/pkg/controllers"
)

func LoadingRoutes() {
	http.HandleFunc("/", control.Index)
}
