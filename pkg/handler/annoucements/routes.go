package annoucements

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *AnnouncementHandler) registerMiddlewareGroup(middleware ...gin.HandlerFunc) gin.IRoutes {
	return s.serve.Gin.Group(s.group).Use(middleware...)
}

func (s *AnnouncementHandler) registerGroup() gin.IRoutes {
	return s.serve.Gin.Group(s.group)
}

func (s *AnnouncementHandler) regularRoutes() http.Handler {
	s.router.GET("/events", s.flashNews)

	return s.serve.Gin
}

func (s *AnnouncementHandler) middlewareRoutes() http.Handler {
	return s.serve.Gin
}
