package main

import (
	"net/http"
	"time"

	"github.com/aliakbarp/student-database/school/handler"
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

	r.Route("/upload", func(r chi.Router) {
		r.Post("/", handler.UploadHandler)
		// other
	})
	println("Listening at port 9000")
	http.ListenAndServe(":9000", r)
}
