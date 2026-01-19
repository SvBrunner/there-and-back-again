package httpapi

import (
	"net/http"

	"github.com/SvBrunner/there-and-back-again/internal/service"
)

func Router(svc service.Service) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	runsHandler := NewRunsHandler(svc)
	mux.Handle("/runs", http.HandlerFunc(runsHandler.Handle))

	return mux
}
