package api

import (
	"github.com/go-chi/chi"
	"net/http"
	"specture/internal/config"
)

func DummyRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/*", bye)
	return r
}

func bye(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, config.GetDummyUrl(), http.StatusMovedPermanently)
}