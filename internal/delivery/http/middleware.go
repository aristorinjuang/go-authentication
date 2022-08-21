package http

import (
	"context"
	"log"
	"net/http"
	"strings"

	jsonFactory "github.com/aristorinjuang/go-authentication/internal/factory/json"
	"github.com/golang-jwt/jwt/v4"
)

type ContextKey string

func User(r *http.Request) *jsonFactory.User {
	value := r.Context().Value(ContextKey("user"))
	if value == nil {
		return nil
	}
	return value.(*jsonFactory.User)
}

func Authentication(accessTokenSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rContext := r.Context()
			token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
			parsedToken, err := jwt.ParseWithClaims(token, &jsonFactory.User{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(accessTokenSecret), nil
			})

			if !parsedToken.Valid || err != nil {
				log.Println(err)

				w.WriteHeader(http.StatusForbidden)
				w.Write(jsonFactory.NewResponse("error", "failed to verify the token", nil).JSON())

				return
			}

			rContext = context.WithValue(rContext, ContextKey("user"), parsedToken.Claims.(*jsonFactory.User))

			next.ServeHTTP(w, r.WithContext(rContext))
		})
	}
}

func JSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
