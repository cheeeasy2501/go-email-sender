package ampq

import (
	"context"
	"encoding/json"
	"fmt"
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
	q    *amqp.Queue
	cons <-chan amqp.Delivery
}

func NewReceiver(cfg *config.Config) *Receiver {
	return &Receiver{
		cfg: cfg.AMQP(),
	}
}

func (r *Receiver) GetQueue() *amqp.Queue {
	return r.q
}

func (r *Receiver) GetDeliveryChan() <-chan amqp.Delivery {
	return r.cons
}

func (r *Receiver) Connect() error {
	conn, err := amqp.Dial(r.cfg.GetConnectionString())
	if err != nil {
		return errs.NewError(
			"ampq.Receiver",
			"Connect",
			"AMPQ connection is't connected",
			fmt.Errorf("AMPQ connection is't connected - %w", err),
		)
	}

	r.conn = conn

	return nil
}

func (r *Receiver) DeclareQueue() error {
	ch, err := r.conn.Channel()
	if err != nil {
		return errs.NewError(
			"ampq.Receiver",
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
		return errs.NewError(
			"ampq.Receiver",
			"DeclareQueue",
			"AMPQ initialization issue code: 2",
			fmt.Errorf("AMPQ QueueDeclare issue - %w", err),
		)
	}
	r.q = &queue

	return nil
}

// TODO: составим тестовый слушатель, для тестирования
func (r *Receiver) CreateTestConsumer() error {
	ch, err := r.conn.Channel()
	if err != nil {
		return errs.NewError(
			"ampq.Receiver",
			"CreateTestConsumer",
			"AMPQ initialization issue code: 4",
			fmt.Errorf("AMPQ CreateTestConsumer issue - %w", err),
		)
	}

	cons, err := ch.Consume(
		qName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return errs.NewError(
			"ampq.Receiver",
			"CreateTestConsumer",
			"AMPQ initialization issue code: 5",
			fmt.Errorf("AMPQ CreateTestConsumer issue - %w", err),
		)
	}

	r.cons = cons

	return nil
}

// Тестовая структура для апи цитат
type TestPublishStruct struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
	SenderName  string `json:"senderName"`
	SenderLink  string `json:"senderLink"`
	QuoteLink   string `json:"quoteLink"`
}

func (r *Receiver) AddTestPublish() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.forismatic.com/api/1.0/?method=getQuote&format=json", nil)
	if err != nil {
		return errs.NewError(
			"ampq.Receiver",
			"AddTestPublish",
			"AMPQ publish issue code: 1",
			fmt.Errorf("AMPQ AddTestPublish issue - %w", err),
		)
	}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return errs.NewError(
			"ampq.Receiver",
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
			"ampq.Receiver",
			"AddTestPublish",
			"AMPQ publish issue code: 1",
			fmt.Errorf("AMPQ AddTestPublish issue - %w", err),
		)
	}

	return nil
}
