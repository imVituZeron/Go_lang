package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"pack.com/loja/pkg/routes"
)

func main() {
	routes.LoadingRoutes()
	http.ListenAndServe(":9001", nil)
}
