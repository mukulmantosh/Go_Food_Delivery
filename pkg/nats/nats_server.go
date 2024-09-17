package nats

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"
	"log"
)

type NATS struct {
	Conn *nats.Conn
}

func NewNATS() (*NATS, error) {
	nc, err := nats.Connect(nats.DefaultURL, nats.Name("food-delivery-nats"))
	if err != nil {
		log.Fatalf("Error connecting to NATS:: %s", err)
	}
	return &NATS{Conn: nc}, err
}

func (n *NATS) Pub(topic string, message []byte) error {
	err := n.Conn.Publish(topic, message)
	if err != nil {
		return err
	}
	return nil
}

func (n *NATS) Sub(topic string, clients map[*websocket.Conn]bool) error {

	_, err := n.Conn.Subscribe(topic, func(msg *nats.Msg) {
		message := string(msg.Data)
		log.Println("Received message from NATS:::", message)
		fmt.Println("CLIENTS::", clients)
		for client := range clients {
			fmt.Println("Sending message to client")
			fmt.Println(client)
			err := client.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("Error sending message to client:", err)
				client.Close()
				delete(clients, client)
			}
		}

	})
	if err != nil {
		return err
	}
	return nil
}
