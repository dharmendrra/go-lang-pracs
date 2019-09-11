package main

import (
	"log"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

const consumerCount = 4

func main() {
	ch, rmqClose, err := new(rabbitmq).dial()
	if err != nil {
		panic(err)
	}
	defer rmqClose()

	start := time.Now()
	consumers(ch)
	log.Println(time.Now().Sub(start))
}

func consumers(ch *amqp.Channel) {
	wg := new(sync.WaitGroup)
	for i := 1; i <= consumerCount; i++ {
		wg.Add(1)
		go func(i int) {
			consume(ch)
			wg.Done()
		}(i)
	}
	log.Println("consumer running #d", consumerCount)
	wg.Wait()
}

func consume(ch *amqp.Channel) {
	msgs, err := ch.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	for m := range msgs {
		process(m)
	}
}

func process(m amqp.Delivery) {
	err := m.Ack(false)
	if err != nil {
		err = errors.Wrap(err, "ack")
		panic(err)
	}
	log.Println(m.Body)
}
