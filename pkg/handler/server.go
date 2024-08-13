package handler

import (
	"Go_Food_Delivery/pkg/database"
	"Go_Food_Delivery/pkg/storage"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	"log/slog"
	"os"
)

type Server struct {
	Gin     *gin.Engine
	db      database.Database
	Storage storage.ImageStorage
}

func NewServer(db database.Database) *Server {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	ginEngine := gin.New()

	// Setting Logger & MultipartMemory
	ginEngine.Use(sloggin.New(logger))
	ginEngine.Use(gin.Recovery())
	ginEngine.MaxMultipartMemory = 8 << 20 // 8 MB

	localStoragePath := os.Getenv("LOCAL_STORAGE_PATH")
	if len(localStoragePath) > 0 {
		// Set static path
		ginEngine.Static(os.Getenv("STORAGE_DIRECTORY"), localStoragePath)
	}

	return &Server{
		Gin:     ginEngine,
		db:      db,
		Storage: storage.CreateImageStorage(os.Getenv("STORAGE_TYPE")),
	}
}

func (server *Server) Run() error {
	return server.Gin.Run(":8080")
}
