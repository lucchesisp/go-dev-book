package routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	Path            string
	Method          string
	Handle          func(w http.ResponseWriter, r *http.Request)
	IsAuthenticated bool
}

func Config(r *mux.Router) *mux.Router {
	routers := usersRouter

	for _, router := range routers {
		r.HandleFunc(router.Path, router.Handle).Methods(router.Method)
	}

	return r
}
