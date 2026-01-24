package main

import (
	"log"
	"net/http"

	"github.com/SvBrunner/there-and-back-again/internal/httpapi"
	"github.com/SvBrunner/there-and-back-again/internal/service"
)

func main() {
	dsn := "file:app.db?_foreign_keys=1&cache=shared&mode=rwc"
	svc, err := service.NewSqliteService(dsn)

	if err != nil {
		log.Fatalf("failed to create service: %v", err)
	}
	mux := httpapi.Router(svc)

	addr := ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
