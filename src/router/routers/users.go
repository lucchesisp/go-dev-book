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
		IsAuthenticated: true,
	},
	{
		Path:            "/users/{id}",
		Method:          http.MethodGet,
		Handle:          controllers.GetUser,
		IsAuthenticated: true,
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
		IsAuthenticated: true,
	},
	{
		Path:            "/users/{id}",
		Method:          http.MethodDelete,
		Handle:          controllers.DeleteUser,
		IsAuthenticated: true,
	},
	{
		Path:            "/users/{id}/update-password",
		Method:          http.MethodPost,
		Handle:          controllers.UpdatePassword,
		IsAuthenticated: true,
	},
	{
		Path:            "/users/{id}/follow",
		Method:          http.MethodPost,
		Handle:          controllers.FollowUser,
		IsAuthenticated: true,
	},
	{
		Path:            "/users/{id}/unfollow",
		Method:          http.MethodPost,
		Handle:          controllers.UnfollowUser,
		IsAuthenticated: true,
	},
	{
		Path:            "/users/{id}/followers",
		Method:          http.MethodGet,
		Handle:          controllers.FindFollowers,
		IsAuthenticated: true,
	},
	{
		Path:            "/users/{id}/following",
		Method:          http.MethodGet,
		Handle:          controllers.FindFollowing,
		IsAuthenticated: true,
	},
}
