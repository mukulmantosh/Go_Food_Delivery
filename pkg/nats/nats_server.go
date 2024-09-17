package nats

import (
	"github.com/nats-io/nats.go"
	"log"
)

type NATS struct {
	conn *nats.Conn
}

func NewNATS() (*NATS, error) {
	nc, err := nats.Connect(nats.DefaultURL, nats.Name("food-delivery-nats"))
	if err != nil {
		log.Fatalf("Error connecting to NATS:: %s", err)
	}
	return &NATS{conn: nc}, err
}

func (n *NATS) Pub(topic string, message []byte) error {
	err := n.conn.Publish(topic, message)
	if err != nil {
		return err
	}
	return nil
}

func (n *NATS) Sub(topic string, message chan<- string) error {
	_, err := n.conn.Subscribe(topic, func(msg *nats.Msg) {
		message <- string(msg.Data)
	})
	if err != nil {
		return err
	}
	return nil
}
