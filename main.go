package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	fmt.Println("Hello World")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT must be set")
	}
	fmt.Println("PORT: " + portString)
	router := chi.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Println("Server is running on port" + portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
