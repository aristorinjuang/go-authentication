package http

import (
	"log"
	"net/http"

	"github.com/aristorinjuang/go-authentication/internal/config"
	"github.com/aristorinjuang/go-authentication/internal/usecase"
	"github.com/gorilla/handlers"
)

func Run(c *config.Config, usecase usecase.Usecase) error {
	log.Printf("The HTTP server is running on port %s.", c.Port)

	return http.ListenAndServe(
		":"+c.Port,
		handlers.CompressHandler(Router(
			c,
			usecase,
		)),
	)
}
