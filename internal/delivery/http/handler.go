package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aristorinjuang/go-authentication/internal/config"
	"github.com/aristorinjuang/go-authentication/internal/entity"
	jsonFactory "github.com/aristorinjuang/go-authentication/internal/factory/json"
	"github.com/aristorinjuang/go-authentication/internal/repository"
	"github.com/aristorinjuang/go-authentication/internal/valueobject"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handler struct {
	config    *config.Config
	repo      repository.Repository
	validator *validator.Validate
}

func (h *handler) Login(res http.ResponseWriter, req *http.Request) {
	var body jsonFactory.Login
	json.NewDecoder(req.Body).Decode(&body)

	if err := h.validator.Struct(body); err != nil {
		log.Println(err)

		res.WriteHeader(http.StatusBadRequest)
		res.Write(jsonFactory.NewResponse("error", err.Error(), nil).JSON())

		return
	}

	user, err := h.repo.Get(valueobject.NewEmail(body.Email))
	if err != nil {
		log.Println(err)

		res.WriteHeader(http.StatusInternalServerError)
		res.Write(jsonFactory.NewResponse("error", "failed to login", nil).JSON())

		return
	}

	if user.Password.Verify(body.Password) {
		res.WriteHeader(http.StatusOK)
		res.Write(jsonFactory.NewResponse(
			"success",
			"",
			jsonFactory.NewUser(
				user,
				h.config.TokenDuration.Access,
				h.config.TokenSecret.Access,
				h.config.TokenSecret.Refresh,
			).Token(),
		).JSON())

		return
	}

	res.WriteHeader(http.StatusNotFound)
	res.Write(jsonFactory.NewResponse("fail", "user not found", nil).JSON())
}

func (h *handler) Register(res http.ResponseWriter, req *http.Request) {
	var body jsonFactory.Register
	json.NewDecoder(req.Body).Decode(&body)

	if err := h.validator.Struct(body); err != nil {
		log.Println(err)

		res.WriteHeader(http.StatusBadRequest)
		res.Write(jsonFactory.NewResponse("error", err.Error(), nil).JSON())

		return
	}

	email := valueobject.NewEmail(body.Email)
	name := valueobject.NewName(body.FirstName, body.LastName)
	password := valueobject.NewPassword("")

	password.SetHash(body.Password)

	user := entity.NewUser(email, name, password)
	err := h.repo.Create(user)
	if err != nil {
		log.Println(err)

		res.WriteHeader(http.StatusInternalServerError)
		res.Write(jsonFactory.NewResponse("error", "failed to register", nil).JSON())

		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(jsonFactory.NewResponse("success", "", nil).JSON())
}

func (h *handler) Token(res http.ResponseWriter, req *http.Request) {
	var body jsonFactory.Token
	json.NewDecoder(req.Body).Decode(&body)

	parsedToken, err := jwt.ParseWithClaims(body.Token, &jsonFactory.User{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.config.TokenSecret.Refresh), nil
	})

	if !parsedToken.Valid || err != nil {
		log.Println(err)

		res.WriteHeader(http.StatusForbidden)
		res.Write(jsonFactory.NewResponse("error", "failed to refresh token", nil).JSON())

		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(jsonFactory.NewResponse(
		"success",
		"",
		jsonFactory.NewUser(
			parsedToken.Claims.(*jsonFactory.User).Entity(),
			h.config.TokenDuration.Access,
			h.config.TokenSecret.Access,
			h.config.TokenSecret.Refresh,
		).AccessToken(),
	).JSON())
}

func (h *handler) Me(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write(jsonFactory.NewResponse(
		"success",
		"",
		User(req),
	).JSON())
}

func NewHandler(config *config.Config, repo repository.Repository) *handler {
	return &handler{
		config:    config,
		repo:      repo,
		validator: validator.New(),
	}
}
