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

	token := c.Query("token")
	if token == "" {
		log.Println("No Token Found!")
		conn.Close()
		return
	}
	userId := "1"
	s.clients[userId] = conn
	log.Printf("New client connected::%s", userId)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Client disconnected: %s::%v", userId, err)
			delete(s.clients, userId)
			break
		}
	}
}
