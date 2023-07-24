package ampq

import (
	"fmt"

	"github.com/cheeeasy2501/go-email-sender/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Receiver struct {
	cfg  config.AMPQ
	conn *amqp.Connection
}

func (r *Receiver) Connect() error {
	conn, err := amqp.Dial(r.cfg.GetConnectionString())
	if err != nil {
		return fmt.Errorf("Reciver - Connected. AMPQ connection is't connected - %w", err)
	}

	r.conn = conn

	return nil
}
