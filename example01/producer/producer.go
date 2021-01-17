package main

import (
	"fmt"
	"log"
	"time"

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

	q, err := ch.QueueDeclare("gophers", false, false, false, false, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(q)

	//envía eternamente mensajes cada 2 segundos
	for {
		err := ch.Publish("", q.Name, false, false,
			amqp.Publishing{
				Headers:     nil,
				ContentType: "text/plain",
				Body:        []byte("sent at " + time.Now().String()),
			})

		if err != nil {
			break
		}

		//wait 2 seconds until send another message
		time.Sleep(2 * time.Second)
	}

}
