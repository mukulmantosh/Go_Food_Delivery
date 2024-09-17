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
		log.Fatalf("Error: %s", err)
	}
	return &NATS{conn: nc}, err
}

func (n *NATS) Publish(channel string, message []byte) error {
	err := n.conn.Publish(channel, message)
	if err != nil {
		return err
	}
	return nil
}

func (n *NATS) Subscribe(channel string) error {
	_, err := n.conn.Subscribe(channel, func(msg *nats.Msg) {
		log.Printf("Received a message: %s", string(msg.Data))
	})
	if err != nil {
		return err
	}
	return nil
}
