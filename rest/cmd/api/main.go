package main

import (
	"log"
	"net/http"

	"github.com/SvBrunner/there-and-back-again/internal/httpapi"
	"github.com/SvBrunner/there-and-back-again/internal/service"
)

func main() {
	svc := service.NewMemoryService()

	mux := httpapi.Router(svc)

	addr := ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
