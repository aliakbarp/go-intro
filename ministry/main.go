package main

import (
	"net/http"
	"time"

	"github.com/aliakbarp/go-intro/ministry/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", http.HandlerFunc(handler.Hello))

	r.Route("/student", func(r chi.Router) {
		r.Post("/", handler.StudentHandler)
		// other
	})
	println("Listening at port 8000")
	http.ListenAndServe(":8000", r)
}
