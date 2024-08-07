package main

import (
	"log"
	"net/http"
	apachepulsar "postgres-pulsar-rest/apache-pulsar"
	"postgres-pulsar-rest/handlers"
	"postgres-pulsar-rest/model"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize DB
	for {
		if err := model.InitDB(); err != nil {
			log.Println("Error Connecting to DB: ", err)
		} else {
			break
		}
	}

	log.Println("Connected to DB")

	model.Migrate()

	// Initialize Pulsar
	for {
		if err := apachepulsar.InitPulsar(); err != nil {
			log.Printf("Failed to initialize Pulsar: %v", err)
		} else {
			break
		}

	}
	log.Println("Connected to Apache Pulsar ")
	defer apachepulsar.Close()

	r := mux.NewRouter()
	r.HandleFunc("/produce", handlers.ProduceMessageHandler).Methods("POST")
	r.HandleFunc("/consume", handlers.ConsumeMessageHandler).Methods("GET")

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
