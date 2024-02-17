package dhttp

import "net/http"

type Auth struct {
	handler *Handler
}

func NewAuth(handler *Handler) *Auth {
	return &Auth{
		handler: handler,
	}
}

func (a Auth) SetupRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /auth/registration", a.handler.initRegistration)
	router.HandleFunc("POST /auth/registration/{identifier}", a.handler.finishRegistration)
	router.HandleFunc("POST /auth/login", a.handler.login)
}
