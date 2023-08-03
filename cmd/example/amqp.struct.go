package example

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cheeeasy2501/go-email-sender/config"
	"github.com/cheeeasy2501/go-email-sender/pkg/errs"
	amqp "github.com/rabbitmq/amqp091-go"
)

var qName string = "sender-emails"

// TODO: добавим в качестве теста в структуру q и cons переменные, для тестирования
type Receiver struct {
	cfg  *config.AMQP
	conn *amqp.Connection
}

func NewReceiver(cfg *config.Config) *Receiver {
	return &Receiver{
		cfg: cfg.AMQP(),
	}
}

func (r *Receiver) CreateNewChannel() (*amqp.Channel, error) {
	ch, err := r.conn.Channel()
	if err != nil {
		return nil, errs.NewError(
			"amqp.Receiver",
			"Connect",
			"AMPQ connection is't connected",
			fmt.Errorf("AMPQ connection is't connected - %w", err),
		)
	}

	return ch, nil
}

func (r *Receiver) Connect() error {
	conn, err := amqp.Dial(r.cfg.GetConnectionString())

	if err != nil {
		return errs.NewError(
			"amqp.Receiver",
			"Connect",
			"AMPQ connection is't connected",
			fmt.Errorf("AMPQ connection is't connected - %w", err),
		)
	}

	r.conn = conn

	return nil
}

func (r *Receiver) DeclareQueue() (amqp.Queue, error) {
	ch, err := r.CreateNewChannel()
	if err != nil {
		return amqp.Queue{}, errs.NewError(
			"amqp.Receiver",
			"DeclareQueue",
			"AMPQ initialization issue code: 1",
			fmt.Errorf("AMPQ Channel issue - %w", err),
		)
	}
	defer ch.Close()

	// TODO: пока сделаем заглушку на создание очереди, потом решим в каком месте нужно создавать её
	queue, err := ch.QueueDeclare(
		qName, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return amqp.Queue{}, errs.NewError(
			"amqp.Receiver",
			"DeclareQueue",
			"AMPQ initialization issue code: 2",
			fmt.Errorf("AMPQ QueueDeclare issue - %w", err),
		)
	}

	return queue, nil
}

// TODO: составим тестовый слушатель, для тестирования
func (r *Receiver) CreateTestConsumer(ch *amqp.Channel) (<-chan amqp.Delivery, error) {
	cons, err := ch.Consume(
		qName,
		"unique-consumer-name",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, errs.NewError(
			"amqp.Receiver",
			"CreateTestConsumer",
			"AMPQ initialization issue code: 5",
			fmt.Errorf("AMPQ CreateTestConsumer issue - %w", err),
		)
	}

	return cons, nil
}

// Тестовая структура для апи цитат
type TestPublishStruct struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
	SenderName  string `json:"senderName"`
	SenderLink  string `json:"senderLink"`
	QuoteLink   string `json:"quoteLink"`
}

func (r *Receiver) AddTestPublish(ch *amqp.Channel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.forismatic.com/api/1.0/?method=getQuote&format=json", nil)
	if err != nil {
		return errs.NewError(
			"amqp.Receiver",
			"AddTestPublish",
			"AMPQ publish issue code: 1",
			fmt.Errorf("AMPQ AddTestPublish issue - %w", err),
		)
	}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return errs.NewError(
			"amqp.Receiver",
			"AddTestPublish",
			"AMPQ publish issue code: 1",
			fmt.Errorf("AMPQ AddTestPublish issue - %w", err),
		)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var ts TestPublishStruct
	err = decoder.Decode(&ts)
	if err != nil {
		return errs.NewError(
			"amqp.Receiver",
			"AddTestPublish",
			"AMPQ publish issue code: 1",
			fmt.Errorf("AMPQ AddTestPublish issue - %w", err),
		)
	}
	log.Printf("%v", ts)
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(ts)

	err = ch.PublishWithContext(ctx,
		"",    // exchange
		qName, // routing key
		false, // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         reqBodyBytes.Bytes(),
		})
	if err != nil {
		return errs.NewError(
			"amqp.Receiver",
			"AddTestPublish",
			"AMPQ publish issue code: 2",
			fmt.Errorf("AMPQ AddTestPublish issue - %w", err),
		)
	}

	return nil
}
