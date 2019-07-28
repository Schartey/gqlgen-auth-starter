package handler

import (
	"net/http"
	"strings"
)

func Login(serviceProvider *HandlerServiceProvider) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		keycloakService := serviceProvider.KeycloakService
		url := keycloakService.GetAuthCodeURL()

		rawAccessToken := r.Header.Get("Authorization")
		if rawAccessToken == "" {
			http.Redirect(w, r, url, http.StatusFound)
			return
		}

		parts := strings.Split(rawAccessToken, " ")
		if len(parts) != 2 {
			w.WriteHeader(400)
			return
		}
		_, err := keycloakService.Verify(ctx, parts[1])

		if err != nil {
			http.Redirect(w, r, url, http.StatusFound)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
	})
}
