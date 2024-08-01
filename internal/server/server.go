package server

import (
	"Uber_Food_Delivery/internal/database"
	"github.com/gin-gonic/gin"
)

type Server struct {
	gin *gin.Engine
	db  database.Database
}

func (server *Server) Gin() *gin.Engine {
	return server.gin
}

func NewServer(db database.Database) *Server {
	return &Server{
		gin: gin.Default(),
		db:  db,
	}
}

func (server *Server) Run() error {
	return server.gin.Run(":8080")
}
