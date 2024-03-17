package middleware

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/swaggest/swgui"
	"github.com/swaggest/swgui/v5emb"
	"gotemplate/internal/config"
	"net/http"
	"os"
	"path/filepath"
)

type Middleware func(next http.Handler) http.Handler

type Manager struct {
	cfg    *config.Config
	router chi.Router
}

func NewManager(router chi.Router) *Manager {
	return &Manager{
		cfg:    config.Get(),
		router: router,
	}
}

func (m *Manager) ApplyMiddleware() {
	m.router.Use(middleware.Logger)
	m.router.Use(middleware.RealIP)
	m.router.Use(middleware.CleanPath)
	m.router.Use(middleware.RequestID)

	m.router.Use(m.cors())
	//m.router.Use(m.tracer)
	m.router.Use(middleware.Recoverer)

	m.swagger()
	m.healthcheck()
	m.router.Handle("/debug", middleware.Profiler())
}

func (m *Manager) cors() Middleware {
	return cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}

func (m *Manager) healthcheck() {
	m.router.Get("/healthz", healthness)
	m.router.Get("/readyz", readyness)
}

func (m *Manager) swagger() {
	wd, _ := os.Getwd()
	docDir := filepath.Join(wd, "api", "openapi")
	doc := http.FileServer(http.Dir(docDir))

	m.router.Mount("/static/", http.StripPrefix("/static/", doc))

	swagConfig := swgui.Config{
		Title:       "API Documentation",
		SwaggerJSON: "/static/gotemplate.yml",
		BasePath:    "/swagger",
		SettingsUI: map[string]string{
			"oauth2RedirectUrl": fmt.Sprintf("\"%s/%s\"", m.cfg.HTTP.Host, "swagger/oauth2-redirect.html"),
		},
	}

	swagHandler := v5emb.NewHandlerWithConfig(swagConfig)
	m.router.Mount("/swagger", swagHandler)
}
