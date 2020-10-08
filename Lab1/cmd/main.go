package main

import (
	"Lab1/internal/pgk/Person/Delivery"
	"Lab1/internal/pgk/Person/Repository"
	"Lab1/internal/pgk/Person/Usecase"
	"Lab1/internal/pgk/middleware"
	"context"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	connectionToBD, state := os.LookupEnv("forDataBase")
	if !state {
		log.Fatal("connection string was not found")
	}

	connectionToServer, errors := pgxpool.Connect(context.Background(), connectionToBD)
	if errors != nil {
		log.Fatal("database connection not established")
	}

	personToRepository := Repository.NewPersonRepository(*connectionToServer)
	personToUsecase := Usecase.NewPersonUsecase(personToRepository)
	personToDelivery := Delivery.NewPersonHandler(personToUsecase)

	router := mux.NewRouter()
	router.Use(middleware.InternalServerError)

	router.HandleFunc("/person/{personID}", personToDelivery.Read).Methods("GET")
	router.HandleFunc("/persons", personToDelivery.ReadAll).Methods("GET")
	router.HandleFunc("/person", personToDelivery.Create).Methods("POST")
	router.HandleFunc("/person/{personID}", personToDelivery.Update).Methods("PATSH")
	router.HandleFunc("/person/{personID}", personToDelivery.Delete).Methods("DELETE")

	srv := &http.Server{
		Handler:      router,
		Addr:         ":5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("Server running at ", srv.Addr)
	log.Fatal(srv.ListenAndServe())

}
