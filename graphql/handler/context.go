package handler

import (
	"context"
	"net/http"
)

const (
	UserInfoKey	string = "UserInfoKey"
)

func AddContext(ctx context.Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
