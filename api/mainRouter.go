package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func InitializeMainRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Group(func(router chi.Router) {
		router.Use(middleware.RealIP)
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)
		router.Mount("/healthz", HealthzRouter())
	})

	router.Group(func(router chi.Router) {
		router.Use(middleware.RequestID)
		router.Use(middleware.RealIP)
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)
		router.Mount("/", DummyRouter())
	})

	router.Group(func(router chi.Router) {
		router.Use(middleware.RequestID)
		router.Use(middleware.RealIP)
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)
		router.Mount("/secret", SecretRouter())
	})

	router.Group(func(router chi.Router) {
		router.Use(middleware.RequestID)
		router.Use(middleware.RealIP)
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)
		router.Mount("/qr", QRRouter())
	})

	return router
}
