package server

import (
	"net/http"

	"github.com/IbzanJared/sistema-reservas/infraestructure/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() http.Handler {
	database, _ := db.ConnectDB()

	handlers := GetHandlers(database)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/reservas", func(r chi.Router) {
			r.Post("/", handlers.ReservaHandler.CrearReserva)
			r.Get("/", handlers.ReservaHandler.ObtenerReservas)
			r.Get("/{id}", handlers.ReservaHandler.ObtenerReservaPorID)
			r.Patch("/{id}", handlers.ReservaHandler.ModificarReserva)
			r.Delete("/{id}", handlers.ReservaHandler.EliminarReserva)
		})
	})

	return r
}
