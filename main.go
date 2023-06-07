package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println("Hello World")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT must be set")
	}
	fmt.Println("PORT: " + portString)
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
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
