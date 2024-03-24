package handlers

import (
	"cognito/books/config"
	"cognito/books/entities"
	"cognito/books/service"
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid/v5"
)

var router http.ServeMux

func GetHandler(config config.AppConfig) http.ServeMux {

	service.Connect(config)
	defer service.Close()

	// Get all
	router.HandleFunc("GET /books", func(w http.ResponseWriter, r *http.Request) {
		requestedIsbn := r.URL.Query().Get("isbn")
		if requestedIsbn != "" {
			books, err := service.GetByISBN(requestedIsbn)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err)
			} else {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(books)
			}
		} else {
			books, err := service.GetAll()
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err)
			} else {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(books)
			}
		}
	})

	// Get by id
	router.HandleFunc("GET /books/{id}", func(w http.ResponseWriter, r *http.Request) {
		requestedId := r.PathValue("id")
		id, uuidErr := uuid.FromString(requestedId)
		if uuidErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(uuidErr)
		}
		books, err := service.GetById(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(books)
		}
	})

	// Add book
	router.HandleFunc("POST /books", func(w http.ResponseWriter, r *http.Request) {
		var newBook entities.Book
		bodyError := json.NewDecoder(r.Body).Decode(&newBook)
		if bodyError != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(bodyError)
		}

		err := service.Create(newBook)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode("Created succesfully")
		}
	})

	// Put by id
	router.HandleFunc("PUT /books/{id}", func(w http.ResponseWriter, r *http.Request) {
		var updatedBook entities.Book
		bodyError := json.NewDecoder(r.Body).Decode(&updatedBook)
		if bodyError != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(bodyError)
		}

		err := service.Update(updatedBook)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode("Updated succesfully")
		}
	})

	// Delete by id
	router.HandleFunc("DELETE /books/{id}", func(w http.ResponseWriter, r *http.Request) {
		requestedId := r.PathValue("id")
		id, uuidErr := uuid.FromString(requestedId)
		if uuidErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(uuidErr)
		}
		err := service.Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted successfully")
		}
	})

	return router
}
