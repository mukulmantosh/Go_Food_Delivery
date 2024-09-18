package notification

import (
	"Go_Food_Delivery/cmd/api/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
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
	valid, userIdInt := middleware.ValidateToken(token)
	if !valid {
		log.Println("Invalid Token!")
		conn.Close()
		return
	}

	userID := strconv.FormatInt(userIdInt, 10)

	s.clients[userID] = conn
	log.Printf("New client connected::%s", userID)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Client disconnected: %s::%v", userID, err)
			delete(s.clients, userID)
			break
		}
	}
}
