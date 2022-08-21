package http

import (
	"log"
	"net/http"

	"github.com/aristorinjuang/go-authentication/internal/config"
	"github.com/aristorinjuang/go-authentication/internal/repository"
	"github.com/gorilla/handlers"
)

func Run(c *config.Config, repo repository.Repository) error {
	log.Printf("The HTTP server is running on port %s.", c.Port)

	return http.ListenAndServe(
		":"+c.Port,
		handlers.CompressHandler(Router(
			c,
			repo,
		)),
	)
}
