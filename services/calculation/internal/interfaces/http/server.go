package http

import (
	"calculation/internal/infra/logger"
	"net/http"
)

func New(logger logger.LoggerInterface, mux *http.ServeMux, port string) *http.Server {
	return &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
}
