package api

import (
	"github.com/go-chi/chi"
	"net/http"
)

func HealthzRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", getHealthStatus)
	r.Get("/*", bye)
	return r
}

func getHealthStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
