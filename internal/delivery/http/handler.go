package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aristorinjuang/go-authentication/internal/config"
	jsonFactory "github.com/aristorinjuang/go-authentication/internal/factory/json"
	"github.com/aristorinjuang/go-authentication/internal/usecase"
	"github.com/aristorinjuang/go-authentication/internal/valueobject"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handler struct {
	config    *config.Config
	usecase   usecase.Usecase
	validator *validator.Validate
}

func (h *handler) Login(res http.ResponseWriter, req *http.Request) {
	var body Login
	json.NewDecoder(req.Body).Decode(&body)

	if err := h.validator.Struct(body); err != nil {
		log.Println(err)

		res.WriteHeader(http.StatusBadRequest)
		res.Write(NewResponse("error", err.Error(), nil).JSON())

		return
	}

	user, err := h.usecase.Login(
		valueobject.NewEmail(body.Email),
		body.Password,
	)
	if err != nil {
		log.Println(err)

		res.WriteHeader(http.StatusInternalServerError)
		res.Write(NewResponse("error", "failed to login", nil).JSON())

		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(NewResponse(
		"success",
		"",
		jsonFactory.NewUser(
			user,
			h.config.TokenDuration.Access,
			h.config.TokenSecret.Access,
			h.config.TokenSecret.Refresh,
		).Token(),
	).JSON())
}

func (h *handler) Register(res http.ResponseWriter, req *http.Request) {
	var body Register
	json.NewDecoder(req.Body).Decode(&body)

	if err := h.validator.Struct(body); err != nil {
		log.Println(err)

		res.WriteHeader(http.StatusBadRequest)
		res.Write(NewResponse("error", err.Error(), nil).JSON())

		return
	}

	err := h.usecase.Register(
		valueobject.NewEmail(body.Email),
		valueobject.NewName(body.FirstName, body.LastName),
		valueobject.NewPasswordFromPlain(body.Password),
	)
	if err != nil {
		log.Println(err)

		res.WriteHeader(http.StatusInternalServerError)
		res.Write(NewResponse("error", "failed to register", nil).JSON())

		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(NewResponse("success", "", nil).JSON())
}

func (h *handler) Token(res http.ResponseWriter, req *http.Request) {
	var body Token
	json.NewDecoder(req.Body).Decode(&body)

	if err := h.validator.Struct(body); err != nil {
		log.Println(err)

		res.WriteHeader(http.StatusBadRequest)
		res.Write(NewResponse("error", err.Error(), nil).JSON())

		return
	}

	parsedToken, err := jwt.ParseWithClaims(body.Token, &jsonFactory.User{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.config.TokenSecret.Refresh), nil
	})

	if !parsedToken.Valid || err != nil {
		log.Println(err)

		res.WriteHeader(http.StatusForbidden)
		res.Write(NewResponse("error", "failed to refresh token", nil).JSON())

		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(NewResponse(
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
	res.Write(NewResponse(
		"success",
		"",
		User(req),
	).JSON())
}

func NewHandler(config *config.Config, usecase usecase.Usecase) *handler {
	return &handler{
		config:    config,
		usecase:   usecase,
		validator: validator.New(),
	}
}
