package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"permit-proxy/handlers"
	"permit-proxy/internal/store"
)

func main() {
	s := store.New()
	h := handlers.NewPermitHandler(s)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/permits", h.HandlePermits)
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
