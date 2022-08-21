package http

import (
	"github.com/aristorinjuang/go-authentication/internal/config"
	"github.com/aristorinjuang/go-authentication/internal/repository"
	"github.com/gorilla/mux"
)

func Router(c *config.Config, repo repository.Repository) *mux.Router {
	h := NewHandler(c, repo)

	router := mux.NewRouter()
	router.HandleFunc("/login", h.Login)
	router.HandleFunc("/register", h.Register)
	router.HandleFunc("/token", h.Token)
	router.Use(JSON)

	authorized := router.PathPrefix("").Subrouter()
	authorized.HandleFunc("/me", h.Me)
	authorized.Use(JSON, Authentication(c.TokenSecret.Access))

	return router
}
