package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/norbertgruszka/booking-app/product-api/pkg/config"
	"github.com/norbertgruszka/booking-app/product-api/pkg/routes"
)

func main() {
	config.ConnectToDb()

	r := chi.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
