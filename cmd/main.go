package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"task/logger_service/config"
	"task/logger_service/events"

	"github.com/saidamir98/udevs_pkg/logger"
)

func main() {
	var loggerLevel string
	cfg := config.Load()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)

	pubsubServer, err := events.New(cfg, log)
	if err != nil {
		log.Panic("error on the event server", logger.Error(err))
	}

	go func() {
		pubsubServer.Run(ctx)
	}()

	log.Info("service is running")

	shutdownChan := make(chan os.Signal, 1)
	defer close(shutdownChan)
	signal.Notify(shutdownChan, syscall.SIGTERM, syscall.SIGINT)
	sig := <-shutdownChan

	cancel()

	log.Info("received os signal", logger.Any("signal", sig))

	log.Info("server shutdown successfully")
}
