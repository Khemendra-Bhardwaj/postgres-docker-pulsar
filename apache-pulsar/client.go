package apachepulsar

import (
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
)

var Client pulsar.Client
var Producer pulsar.Producer
var Consumer pulsar.Consumer

func InitPulsar() error {
	var err error
	// Create a Client
	for {
		Client, err = pulsar.NewClient(pulsar.ClientOptions{
			URL: "pulsar://pulsar:6650",
		})
		if err == nil {
			break
		}
	}

	// Create a producer
	for {
		Producer, err = Client.CreateProducer(pulsar.ProducerOptions{
			Topic: "example-topic",
		})
		if err == nil {
			break
		}
	}

	log.Println("Created Producer ")
	// Create a consumer
	for {
		Consumer, err = Client.Subscribe(pulsar.ConsumerOptions{
			Topic:            "example-topic",
			SubscriptionName: "example-subscription",
		})
		if err == nil {
			break
		}
	}

	return nil
}

func Close() {
	if Producer != nil {
		Producer.Close()
	}
	if Consumer != nil {
		Consumer.Close()
	}
	if Client != nil {
		Client.Close()
	}
}
