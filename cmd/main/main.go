package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/madnaaaaas/crud/pkg/config"
	"github.com/madnaaaaas/crud/pkg/logger"
	"github.com/madnaaaaas/crud/pkg/server"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	log, err := logger.NewLogger(cfg.LogLevel)
	if err != nil {
		panic(err)
	}

	srv, err := server.NewServer(cfg, log)
	if err != nil {
		log.Fatal("can't create server", zap.Any("error", err))
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	if err = srv.Start(); err != nil {
		log.Fatal("can't start server", zap.Any("error", err))
	}

	<-signalChan
	if err = srv.Shutdown(context.Background()); err != nil {
		log.Fatal("error shutting server", zap.Any("error", err))
	}
}
