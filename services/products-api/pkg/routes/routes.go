package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/norbertgruszka/booking-app/product-api/pkg/controllers"
)

func RegisterRoutes(routes *chi.Mux) {
	routes.Get("/health", controllers.Health)
}
