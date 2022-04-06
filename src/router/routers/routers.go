package routers

import (
	"github.com/gorilla/mux"
	"github.com/lucchesisp/go-dev-book/src/middlewares"
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
	routers = append(routers, signinRouter)
	routers = append(routers, publicationsRouter...)

	for _, router := range routers {

		if router.IsAuthenticated {
			r.HandleFunc(router.Path,
				middlewares.Logger(
					middlewares.Authentication(router.Handle)),
			).Methods(router.Method)
		} else {
			r.HandleFunc(router.Path, middlewares.Logger(router.Handle)).Methods(router.Method)
		}
	}

	return r
}
