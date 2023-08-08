package amqp

import (
	"fmt"
	"reflect"

	"github.com/cheeeasy2501/go-email-sender/config"
	"github.com/cheeeasy2501/go-email-sender/pkg/errs"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	cfg           *config.AMQP
	conn          *amqp.Connection
	devliveryChan <-chan amqp.Delivery
}

func CreateNewConsumer(cfg *config.AMQP) Consumer {
	return Consumer{
		cfg: cfg,
	}
}

// попытка подключения к очереди
func (c *Consumer) Connect() error {
	conn, err := amqp.Dial(c.cfg.GetConnectionString())

	if err != nil {
		return errs.NewError(
			"amqp.Consumer",
			"Connect",
			"AMPQ connection is't connected",
			fmt.Errorf("AMPQ connection is't connected - %w", err),
		)
	}

	c.conn = conn

	return nil
}

// открытие канала доставки
func (c *Consumer) OpenChannel() (*amqp.Channel, error) {
	ch, err := c.conn.Channel()
	if err != nil {
		return nil, errs.NewError(
			"amqp.Consumer",
			"OpenChannel",
			"Consumer channel not opened",
			fmt.Errorf("Consumer channel not opened - %w", err),
		)
	}

	return ch, nil
}

type ConsumerOptions struct {
	queue, consumerTag                  string
	autoAck, exclusive, noLocal, noWait bool
	args                                amqp.Table
}

func NewConsumerOptions(queue, consumerTag string, autoAck, exclusive, noLocal, noWait bool, args map[string]interface{}) ConsumerOptions {
	return ConsumerOptions{
		queue:       queue,
		consumerTag: consumerTag,
		autoAck:     autoAck,
		exclusive:   exclusive,
		noLocal:     noLocal,
		noWait:      noWait,
		args:        args,
	}
}

func (c *Consumer) DoDeliveryChannel(ch *amqp.Channel, o ConsumerOptions) error {
	delivery, err := ch.Consume(
		o.queue,
		o.consumerTag,
		o.autoAck,
		o.exclusive,
		o.noLocal,
		o.noWait,
		o.args,
	)

	if err != nil {
		return errs.NewError(
			"amqp.Receiver",
			"CreateTestConsumer",
			"AMPQ initialization issue code: 5",
			fmt.Errorf("AMPQ CreateTestConsumer issue - %w", err),
		)
	}

	c.devliveryChan = delivery

	return nil
}

// быстрое создание consumer
func FastCreateNewConsumer(cfg *config.AMQP, o ConsumerOptions) (Consumer, error) {
	sn := "amqp"
	fn := "FastCreateNewConsumer"
	m := "Consumer not created"

	c := Consumer{
		cfg: cfg,
	}

	err := c.Connect()
	if err != nil {
		return Consumer{}, errs.NewError(
			sn,
			fn,
			m,
			fmt.Errorf("Consumer can't open connect by amqp - %w", err),
		)
	}

	ch, err := c.OpenChannel()
	if err != nil {
		return Consumer{}, errs.NewError(
			sn,
			fn,
			m,
			fmt.Errorf("Consumer can't open channel - %w", err),
		)
	}

	// Дефолтные настройки
	if reflect.DeepEqual(o, ConsumerOptions{}) {
		o = NewConsumerOptions(
			"test-queue",
			"",
			false,
			false,
			false,
			false,
			nil,
		)
	}

	err = c.DoDeliveryChannel(ch, o)
	if err != nil {
		return Consumer{}, errs.NewError(
			sn,
			fn,
			m,
			fmt.Errorf("Consumer don't make delivery channel - %w", err),
		)
	}

	return c, nil
}

// Получение канала доставки
func (c *Consumer) Delivery() <-chan amqp.Delivery {
	return c.devliveryChan
}
