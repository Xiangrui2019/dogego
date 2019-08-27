package modules

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

var AMQPModule *amqp.Connection

func InitAMQPModule() {
	var err error

	AMQPModule, err = amqp.Dial(os.Getenv("AMQP_ADDR"))

	if err != nil {
		log.Fatal(err)
	}
}
