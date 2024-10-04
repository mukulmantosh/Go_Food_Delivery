package annoucements

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func (s *AnnouncementHandler) flashNews(c *gin.Context) {
	_, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	events, err := s.service.FlashEvents()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	// Set headers for SSE
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	ticker := time.NewTicker(6 * time.Second)
	defer ticker.Stop()

	eventIndex := 0

	for {
		select {
		case <-ticker.C:
			// Send the current event
			event := (*events)[eventIndex]
			c.SSEvent("message", event.Message)
			c.Writer.Flush()

			// Move to the next event
			eventIndex = (eventIndex + 1) % len(*events)
		case <-c.Request.Context().Done():
			ticker.Stop()
			return
		}
	}

}
