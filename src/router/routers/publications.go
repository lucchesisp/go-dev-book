package routers

import (
	"github.com/lucchesisp/go-dev-book/src/controllers"
	"net/http"
)

var publicationsRouter = []Router{
	{
		Path:            "/publications",
		Method:          http.MethodGet,
		Handle:          controllers.GetPublications,
		IsAuthenticated: true,
	},
	{
		Path:            "/publications/{id}",
		Method:          http.MethodGet,
		Handle:          controllers.GetPublication,
		IsAuthenticated: true,
	},
	{
		Path:            "/publications",
		Method:          http.MethodPost,
		Handle:          controllers.CreatePublication,
		IsAuthenticated: true,
	},
	{
		Path:            "/publications/{id}",
		Method:          http.MethodPut,
		Handle:          controllers.UpdatePublication,
		IsAuthenticated: true,
	},
	{
		Path:            "/publications/{id}",
		Method:          http.MethodDelete,
		Handle:          controllers.DeletePublication,
		IsAuthenticated: true,
	},
}
