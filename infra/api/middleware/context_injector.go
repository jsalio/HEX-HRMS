package middleware

import (
	"hrms/repository/postgress/repo"
	"net/http"
)

// middleware/context_injector.go
func InjectContextToGenericRepo[T any, G any](repo *repo.GenericCrud[T, G]) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			repo.WithContext(r.Context())
			next.ServeHTTP(w, r)
		})
	}
}
