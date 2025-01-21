package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ognick/job-interview-playground/internal/yggdrasil/config"
	"github.com/ognick/job-interview-playground/internal/yggdrasil/internal/services/wisdom"
	wisdomapiv1 "github.com/ognick/job-interview-playground/internal/yggdrasil/internal/services/wisdom/api/v1"
	wisdomrepo "github.com/ognick/job-interview-playground/internal/yggdrasil/internal/services/wisdom/repository"
	"github.com/ognick/job-interview-playground/pkg/httpsrv"
	"github.com/ognick/job-interview-playground/pkg/logger/zap"
	"github.com/ognick/job-interview-playground/pkg/shutdown"
)

func Run() {
	logger := zap.NewLogger()
	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatalf("failed to load config: %v", err)
	}
	// Init logger
	logger.InitLogger(cfg.Logger)

	// Init router
	router := gin.Default()
	router.Use(gin.Recovery())

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// Init wisdom service
	wisdomInmemoryRepo := wisdomrepo.NewInmemoryRepository()
	wisdomUsecase := wisdom.NewUsecase(logger, wisdomInmemoryRepo)
	wisdomV1 := wisdomapiv1.NewHandler(logger, wisdomUsecase)
	wisdomV1.Register(router)

	// Start HTTP server
	runner, gracefulCtx := shutdown.CreateRunnerWithGracefulContext()
	srv := httpsrv.NewServer(logger, cfg.HTTPAddress, router)
	runner.Go(func() error {
		return srv.Run(gracefulCtx)
	})

	// Awaiting graceful shutdown
	logger.Infof("Application started")
	if err := runner.Wait(); err != nil {
		logger.Fatalf("%v", err)
	}
	logger.Infof("Application gracefully finished")
}
