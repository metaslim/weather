package di_container

import (
	"context"
	"net/http"
)

var dicCtxKey = &contextKey{"dic"}

type contextKey struct {
	name string
}

func ContextWithDIC(ctx context.Context, dic *DIContainer) context.Context {
	return context.WithValue(ctx, dicCtxKey, dic)
}

func DependencyInjectionMiddleware(dic *DIContainer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// put it in context
			ctx := ContextWithDIC(req.Context(), dic)
			// and call the next with our new context
			req = req.WithContext(ctx)
			next.ServeHTTP(w, req)
		})
	}
}

func DIC(ctx context.Context) *DIContainer {
	return ctx.Value(dicCtxKey).(*DIContainer)
}
