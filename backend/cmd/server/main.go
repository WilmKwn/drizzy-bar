package main

import (
	"drizzy-backend/pkg/api"
	"drizzy-backend/pkg/db"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	log.Print("server started")

	pgdb, err := db.StartDB()
	if err != nil {
		log.Printf("error starting the database %v", err)
	}

	router := api.StartAPI(pgdb)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	handler := corsHandler.Handler(router)

	err = http.ListenAndServe(fmt.Sprintf(":%s", "8080"), handler)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}
