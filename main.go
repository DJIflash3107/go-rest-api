package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not set in the environment")
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
		MaxAge: 300,
	}))
	router.Get("/", handlerGetUser)
	router.Post("/user", handlerAddUser)
	server := &http.Server{
		Addr: ":" + portString,
		Handler: router,
	}
	log.Printf("Server starting on port %v", portString)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}