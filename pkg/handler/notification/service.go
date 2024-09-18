package notification

import (
	"Go_Food_Delivery/pkg/handler"
	"Go_Food_Delivery/pkg/service/notification"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
	"net/http"
)

type NotifyHandler struct {
	serve             *handler.Server
	group             string
	middlewareGuarded gin.IRoutes
	router            gin.IRoutes
	service           *notification.NotificationService
	middleware        []gin.HandlerFunc
	validate          *validator.Validate
	ws                *websocket.Upgrader
	clients           map[string]*websocket.Conn
}

func NewNotifyHandler(s *handler.Server, group string,
	service *notification.NotificationService, middleware []gin.HandlerFunc,
	validate *validator.Validate, clients map[string]*websocket.Conn) {

	// WebSocket
	var ws = &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	cartHandler := &NotifyHandler{
		s,
		group,
		nil,
		nil,
		service,
		middleware,
		validate,
		ws,
		clients,
	}
	cartHandler.middlewareGuarded = cartHandler.registerMiddlewareGroup(middleware...)
	cartHandler.router = cartHandler.registerGroup()
	cartHandler.regularRoutes()
	cartHandler.middlewareRoutes()
	cartHandler.registerValidator()
}

func (s *NotifyHandler) registerValidator() {

}
