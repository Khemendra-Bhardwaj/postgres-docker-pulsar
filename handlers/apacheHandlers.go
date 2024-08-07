// handlers/handlers.go
package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	apachepulsar "postgres-pulsar-rest/apache-pulsar"
	"postgres-pulsar-rest/model"

	"github.com/apache/pulsar-client-go/pulsar"
)

// MessageRequest represents the request payload for producing a message
type MessageRequest struct {
	Content string `json:"content"`
}

// ProduceMessageHandler handles message production requests
func ProduceMessageHandler(w http.ResponseWriter, r *http.Request) {
	var request MessageRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Produce the message
	_, err := apachepulsar.Producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(request.Content),
	})
	if err != nil {
		http.Error(w, "Failed to produce message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message produced successfully"))
}

// ConsumeMessageHandler handles message consumption requests
func ConsumeMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Consume messages in a separate goroutine since it is a long-running process
	go func() {
		for {
			msg, err := apachepulsar.Consumer.Receive(context.Background())
			if err != nil {
				log.Printf("Failed to receive message: %v", err)
				continue
			}

			content := string(msg.Payload())
			if err := model.SaveMessage(content); err != nil {
				log.Printf("Failed to save message to DB: %v", err)
			} else {
				log.Printf("Message saved: %s", content)
			}

			apachepulsar.Consumer.Ack(msg)
		}
	}()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message consumption started"))
}
