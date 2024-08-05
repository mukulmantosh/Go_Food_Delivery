package handler

import (
	"Go_Food_Delivery/pkg/database"
	"github.com/gin-gonic/gin"
)

type Server struct {
	gin *gin.Engine
	db  database.Database
}

func (server *Server) Engine() database.Database {
	return server.db
}

func (server *Server) Gin() *gin.Engine {
	return server.gin
}

func NewServer(db database.Database) *Server {
	ginEngine := gin.Default()
	ginEngine.MaxMultipartMemory = 8 << 20 // 8 MB

	return &Server{
		gin: ginEngine,
		db:  db,
	}
}

func (server *Server) Run() error {
	return server.gin.Run(":8080")
}
