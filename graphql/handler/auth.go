package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	logger "github.com/sirupsen/logrus"
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

func LoginCallback(serviceProvider *HandlerServiceProvider) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	state := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")

	userInfo, err := serviceProvider.UserService.HandleLogin(ctx, state, code)
	if err != nil {
	w.WriteHeader(401)
	log.Println(w, err)
}

	ctx = context.WithValue(ctx, UserInfoKey, userInfo)
	logger.Infof("%s", userInfo.RawIDToken.(string))

	fmt.Fprintf(w, "%s", userInfo.RawIDToken.(string))
})

}