package handler

import (
	"Go_Food_Delivery/pkg/database"
	"Go_Food_Delivery/pkg/storage"
	"github.com/gin-contrib/cors"
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

func NewServer(db database.Database, setLog bool) *Server {

	ginEngine := gin.New()

	// CORS configuration
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},                                       // List of allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // List of allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allow the necessary headers
		AllowCredentials: true,
	}

	// Setting Logger, CORS & MultipartMemory
	if setLog {
		logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
		ginEngine.Use(sloggin.New(logger))
	}
	ginEngine.Use(gin.Recovery())
	ginEngine.Use(cors.New(corsConfig))
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
