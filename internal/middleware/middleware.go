package middleware

import (
	"net/http"
)

type Middlewares func(http.Handler) http.Handler

func CreateStack(xs ...Middlewares) Middlewares {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}

type Middleware struct {
	Middleware interface {
	}
}

func NewMiddleware(env string) Middleware {
	return Middleware{
		Middleware: &MiddlewareStore{
			ENV: env,
		},
	}
}

type MiddlewareStore struct {
	ENV string
}
