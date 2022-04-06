package routers

import (
	"github.com/lucchesisp/go-dev-book/src/controllers"
	"net/http"
)

var signinRouter = Router{
	Path:            "/signin",
	Method:          http.MethodPost,
	Handle:          controllers.Signin,
	IsAuthenticated: false,
}
