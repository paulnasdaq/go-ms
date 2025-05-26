package common

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Connector struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func (c *Connector) Close() error {
	return c.conn.Close()
}
func (c *Connector) PublishEvent(topic string, event []byte) error {
	return c.ch.Publish("amq.topic", topic, false, false, amqp.Publishing{Timestamp: time.Time{}, Body: event})
}

func NewConnector() (*Connector, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println("Failed to connect to AMQP broker: ", err)
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to create channel: ", err)
		conn.Close()
		return nil, err
	}
	return &Connector{conn: conn, ch: ch}, nil
}

func main() {

}
