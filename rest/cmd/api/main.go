package main

import (
	"log"
	"net/http"

	"github.com/SvBrunner/thereandbackagain/internal/httpapi"
)

func main() {
	mux := httpapi.Router(myService)

	addr := ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
