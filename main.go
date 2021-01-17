package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	//run with docker
	// docker run --rm -it --hostname my-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
	//user guest , pass guest
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	fmt.Println("hello word")
}
