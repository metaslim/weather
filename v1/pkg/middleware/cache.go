package middleware

import (
	"fmt"
	"net/http"
)

func BeforeMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			fmt.Println(req.URL)
			next.ServeHTTP(w, req)
		})
	}
}
