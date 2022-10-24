package events

import (
	"context"
	"task/logger_service/config"
	consumer "task/logger_service/events/logger"
	"task/logger_service/pkg/pubsub"

	"github.com/saidamir98/udevs_pkg/logger"
)

// PubsubServer ...
type PubsubServer struct {
	cfg config.Config
	log logger.LoggerI
	RMQ *pubsub.RMQ
}

// New ...
func New(cfg config.Config, log logger.LoggerI) (*PubsubServer, error) {
	rmq, err := pubsub.NewRMQ(cfg.RabbitMqURL, log)
	if err != nil {
		return nil, err
	}

	rmq.AddPublisher(cfg.ExchangeName) // one publisher is enough for application service

	return &PubsubServer{
		cfg: cfg,
		log: log,
		RMQ: rmq,
	}, nil
}

// Run ...
func (s *PubsubServer) Run(ctx context.Context) {
	loggerServer := consumer.New(s.cfg, s.log, s.RMQ)
	loggerServer.RegisterConsumers()

	s.RMQ.RunConsumers(ctx)
}
