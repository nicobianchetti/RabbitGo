package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	qName := "gophers"

	chDelivery, err := ch.Consume(
		qName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	//dejar un evento eternamente consumiendo
	//hacer un chanel que siempre escuche para bloquear la goroutine princial(main)

	noStop := make(chan bool)

	go func() {
		for delivery := range chDelivery {
			fmt.Println("msg: " + string(delivery.Body))
		}
	}()

	//noStop nunca recibio en un mensaje , entonces nunca va a mandar un mensaje tampoco
	<-noStop

}
