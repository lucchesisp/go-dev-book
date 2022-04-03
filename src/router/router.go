package router

import (
	"github.com/gorilla/mux"
	"github.com/lucchesisp/go-dev-book/src/router/routers"
)

// GenerateRouter returns a new router.
func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return routers.Config(r)
}
