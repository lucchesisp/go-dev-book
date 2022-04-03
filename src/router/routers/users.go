package routers

import (
	"github.com/lucchesisp/go-dev-book/src/controllers"
	"net/http"
)

var usersRouter = []Router{
	{
		Path:            "/users",
		Method:          http.MethodGet,
		Handle:          controllers.GetUsers,
		IsAuthenticated: false,
	},
	{
		Path:            "/users/{id}",
		Method:          http.MethodGet,
		Handle:          controllers.GetUser,
		IsAuthenticated: false,
	},
	{
		Path:            "/users",
		Method:          http.MethodPost,
		Handle:          controllers.CreateUser,
		IsAuthenticated: false,
	},
	{
		Path:            "/users/{id}",
		Method:          http.MethodPut,
		Handle:          controllers.UpdateUser,
		IsAuthenticated: false,
	},
	{
		Path:            "/users/{id}",
		Method:          http.MethodDelete,
		Handle:          controllers.DeleteUser,
		IsAuthenticated: false,
	},
}
