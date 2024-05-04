package middlewares


import (
    "context"
    "net/http"

    "github.com/26thavenue/creditCardValidator/db"
)

func GormMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := context.WithValue(r.Context(), "db", db.DB)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}