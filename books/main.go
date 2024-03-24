package main

import (
	. "cognito/books/config"
	"cognito/books/handlers"
	"fmt"
	"net/http"
	"strconv"
)

var bookConfig AppConfig

func main() {
	bookConfig = LoadConfig()

	router := handlers.GetHandler(bookConfig)

	server := http.Server{
		Addr:    ":" + strconv.Itoa(bookConfig.Server.Port),
		Handler: &router,
	}

	fmt.Printf("Server listening on port :%d", bookConfig.Server.Port)
	server.ListenAndServe()

}
