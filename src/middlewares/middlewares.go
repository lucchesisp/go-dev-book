package middlewares

import (
	"fmt"
	"github.com/lucchesisp/go-dev-book/src/authentication"
	"github.com/lucchesisp/go-dev-book/src/responses"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.Method, r.URL.Path, r.Host)
		next(w, r)
	}
}

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			responses.JSON(w, http.StatusUnauthorized, map[string]string{"message": "Unauthorized: " + err.Error()})
			return
		}

		next(w, r)
	}
}
