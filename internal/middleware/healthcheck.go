package middleware

import (
	"gotemplate/pkg/healthcheck"
	"net/http"
)

func healthness(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(nil)
}

func readyness(w http.ResponseWriter, _ *http.Request) {
	if healthcheck.Get().IsAlive() {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	_, _ = w.Write(nil)
}
