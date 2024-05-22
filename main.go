package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type CardHandler struct {
	
}

func(c CardHandler) Validate( w http.ResponseWriter, r *http.Request) {

	var card Card

	err:= json.NewDecoder(r.Body).Decode(&card)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	isValid := ValidateCard(card)

	response := make(map[string]string)
    if !isValid {
        response["message"] = "Invalid card number"
    } else {
        response["message"] = "Valid card number"
    } 

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
    if err != nil {
        http.Error(w, "Internal error", http.StatusInternalServerError)
        return
    }


}


func main() {
    r := chi.NewRouter()

    r.Use(middleware.Logger)

    handler := CardHandler{}

    r.Post("/", handler.Validate)

    r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })


    http.ListenAndServe(":8080", r)
}