package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

type rabbitmq struct {
}

const (
	prefetchCount = 1
	queueName     = "Q_workergroup"
)

var rabbitmqURLs = map[string]string{
	"test":        "amqp://guest:guest@localhost:5672/V_localhost",
	"development": "amqp://guest:guest@localhost:5672/V_localhost",
}

func (r *rabbitmq) getAMQPURL() string {
	env := os.Getenv("env")
	if env == "" {
		env = "development"
	}
	return rabbitmqURLs[env]
}

func (r *rabbitmq) dial() (*amqp.Channel, func(), error) {
	url := r.getAMQPURL()
	connection, err := amqp.Dial(url)
	if err != nil {
		return nil, nil, errors.Wrap(err, "amqp dial")
	}
	close := func() {
		connection.Close()
	}

	ch, err := connection.Channel()
	if err != nil {
		return nil, close, errors.Wrap(err, "amqp conn channel")
	}
	close = func() {
		ch.Close()
		close()
	}

	_, err = ch.QueueDeclare(queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, close, errors.Wrap(err, "channel queue declare")
	}

	err = ch.Qos(prefetchCount, 0, false)
	if err != nil {
		return nil, close, errors.Wrap(err, "channel prefetch count")
	}

	return ch, close, nil

}
