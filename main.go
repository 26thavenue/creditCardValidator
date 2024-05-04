package main

import (
    "net/http"

    "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
    "github.com/26thavenue/creditCardValidator/middlewares"
    "github.com/26thavenue/creditCardValidator/handlers"
)

func main() {
    r := chi.NewRouter()

    r.Use(middleware.Logger)
    r.Use(middlewares.GormMiddleware)

    // Define routes
    r.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.ListAllUsers)
		r.Post("/", handlers.CreateUser)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetUser)
			r.Put("/", handlers.UpdateUser)
			r.Delete("/", handlers.DeleteUser)
		})  
    })

	r.Route("/creditcards", func(r chi.Router) {
		r.Post("/", handlers.AddCreditCardHandler)
	})


    // Start the server
    http.ListenAndServe(":8080", r)
}