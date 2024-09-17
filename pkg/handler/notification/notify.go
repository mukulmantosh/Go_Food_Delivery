package notification

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"strings"
)

func (s *NotifyHandler) notifyOrders(c *gin.Context) {
	conn, err := s.ws.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error while upgrading connection:", err)
		return
	}
	defer conn.Close()

	userID := c.GetInt64("userID")

	for {
		select {
		case msg := <-*s.message:
			parts := strings.Split(msg, "|")
			if len(parts) != 2 {
				log.Fatalf("Invalid input format")
				return
			}

			// Split the parts to get user ID and message
			userIdPart := strings.Split(parts[0], ":")
			messagePart := strings.Split(parts[1], ":")

			if len(userIdPart) != 2 || len(messagePart) != 2 {
				log.Fatalf("Invalid input format")
				return
			}

			// Extract and convert user ID
			userIdStr := userIdPart[1]
			userId, err := strconv.ParseInt(userIdStr, 10, 64)
			if err != nil {
				log.Fatalf("error converting user ID: %s", err)
				return
			}

			if userId != userID {
				continue
			}

			// Extract message
			message := messagePart[1]

			err = conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("Error while writing message to WebSocket:", err)
				return
			}
		}
	}
}
