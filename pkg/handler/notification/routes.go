package notification

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *NotifyHandler) registerMiddlewareGroup(middleware ...gin.HandlerFunc) gin.IRoutes {
	return s.serve.Gin.Group(s.group).Use(middleware...)
}

func (s *NotifyHandler) registerGroup() gin.IRoutes {
	return s.serve.Gin.Group(s.group)
}

func (s *NotifyHandler) regularRoutes() http.Handler {
	s.router.GET("/ws", s.notifyOrders)
	return s.serve.Gin
}

func (s *NotifyHandler) middlewareRoutes() http.Handler {

	return s.serve.Gin
}
