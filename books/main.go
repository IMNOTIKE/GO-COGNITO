package main

import (
	. "cognito/books/config"
	"cognito/books/handlers"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
)

func main() {
	bookConfig, err := LoadConfig()
	if err != nil {
		log.Err(err)
		os.Exit(1)
	}

	router := handlers.GetHandler(bookConfig)

	server := http.Server{
		Addr:    ":" + strconv.Itoa(bookConfig.Server.Port),
		Handler: &router,
	}

	fmt.Printf("Server listening on port :%d", bookConfig.Server.Port)
	server.ListenAndServe()

}
