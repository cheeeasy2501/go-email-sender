package example

import (
	"context"
	"log"
	"time"

	"github.com/cheeeasy2501/go-email-sender/config"
)

func Start(ctx context.Context, cfg *config.Config) {
	// TODO: тест amqp
	r := NewReceiver(cfg)

	err := r.Connect()
	if err != nil {
		panic(err)
	}

	_, err = r.DeclareQueue()
	if err != nil {
		panic(err)
	}

	go StartConsumer(r)

	go StartPublisher(r)
}

func StartConsumer(r *Receiver) {
	consChan, err := r.CreateNewChannel()
	if err != nil {
		panic(err)
	}

	consumer, err := r.CreateTestConsumer(consChan)
	if err != nil {
		panic(err)
	}

	for d := range consumer {
		log.Printf("Received a message: %s\n", d.Body)
		t := time.Duration(5)
		time.Sleep(t * time.Second)
	}
}

func StartPublisher(r *Receiver) {
	pubChan, err := r.CreateNewChannel()
	if err != nil {
		panic(err)
	}

	for {
		err := r.AddTestPublish(pubChan)
		if err != nil {
			log.Printf("------- Error publishing ------- %w", err)
		} else {
			log.Println("------- Test message published -------")
		}

		time.Sleep(1 * time.Second)
	}
}
