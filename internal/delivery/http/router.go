package http

import (
	"github.com/aristorinjuang/go-authentication/internal/config"
	"github.com/aristorinjuang/go-authentication/internal/usecase"
	"github.com/gorilla/mux"
)

func Router(c *config.Config, usecase usecase.Usecase) *mux.Router {
	h := NewHandler(c, usecase)

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
