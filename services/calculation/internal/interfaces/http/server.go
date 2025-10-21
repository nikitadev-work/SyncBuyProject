package http

import (
	"net/http"

	"github.com/nikitadev-work/SyncBuyProject/common/kit/logger"
)

func New(logger logger.LoggerInterface, mux *http.ServeMux, port string) *http.Server {
	return &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
}
