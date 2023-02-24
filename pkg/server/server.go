package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/madnaaaaas/crud/pkg/config"
	"github.com/madnaaaaas/crud/pkg/database"
	"github.com/madnaaaaas/crud/pkg/logger"
	"github.com/madnaaaaas/crud/pkg/refrigerator/repo"
	"github.com/madnaaaaas/crud/pkg/refrigerator/rest"
	"github.com/madnaaaaas/crud/pkg/refrigerator/service"
)

type server struct {
	httpServer *http.Server
	log        *zap.Logger
}

func NewServer(cfg *config.Config, log *zap.Logger) (*server, error) {
	db, err := database.NewDatabaseConnection(cfg)
	if err != nil {
		return nil, err
	}

	refrigeratorRepo := repo.NewRepo(db)
	refrigeratorService := service.NewService(refrigeratorRepo)
	refrigeratorRest := rest.NewRest(refrigeratorService)

	logMiddleware := logger.NewLogMiddleware(log)

	router := gin.New()
	api := router.Group("/api/v1")
	api.Use(logMiddleware.Logging)

	refrigeratorRest.Register(api)

	return &server{
		httpServer: &http.Server{
			Addr:    ":" + cfg.ServerPort,
			Handler: router,
		},
		log: log,
	}, nil
}

func (s *server) Start() error {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			s.log.Error(err.Error())
		}
	}()

	s.log.Debug("server started")
	return nil
}

func (s *server) Shutdown(ctx context.Context) error {
	if s.httpServer != nil {
		return s.httpServer.Shutdown(ctx)
	}

	s.log.Debug("server stopped")
	return nil
}
