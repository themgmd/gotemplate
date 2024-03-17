package dhttp

import (
	"github.com/go-chi/chi/v5"
)

type Auth struct {
	handler *Handler
}

func NewAuth(handler *Handler) *Auth {
	return &Auth{
		handler: handler,
	}
}

func (a Auth) SetupRoutes(router chi.Router) {
	router.HandleFunc("/auth/registration", a.handler.initRegistration)
	router.HandleFunc("/auth/registration/{identifier}", a.handler.finishRegistration)
	router.HandleFunc("/auth/login", a.handler.login)
}
