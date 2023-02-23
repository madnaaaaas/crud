package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/madnaaaaas/crud/pkg/config"
	"github.com/madnaaaaas/crud/pkg/server"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	srv, err := server.NewServer(cfg)
	if err != nil {
		panic(err)
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	if err = srv.Start(); err != nil {
		log.Fatal(err.Error())
	}

	<-signalChan
	if err = srv.Shutdown(context.Background()); err != nil {
		log.Fatal(err.Error())
	}
}
