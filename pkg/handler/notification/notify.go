package notification

import (
	"github.com/gin-gonic/gin"
	"log"
)

func (s *NotifyHandler) notifyOrders(c *gin.Context) {
	conn, err := s.ws.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error while upgrading connection:", err)
		return
	}
	defer conn.Close()

	s.clients[conn] = true
	log.Println("New client connected")

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("Client disconnected:", err)
			delete(s.clients, conn)
			break
		}
	}
}
