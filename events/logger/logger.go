package consumer

import (
	"task/logger_service/config"
	"task/logger_service/pkg/pubsub"

	"github.com/saidamir98/udevs_pkg/logger"
)

// Consumer ...
type Logger struct {
	cfg config.Config
	log logger.LoggerI
	rmq *pubsub.RMQ
}

// New ...
func New(cfg config.Config, log logger.LoggerI, rmq *pubsub.RMQ) *Logger {
	return &Logger{
		cfg: cfg,
		log: log,
		rmq: rmq,
	}
}

// RegisterConsumers ...
func (s *Logger) RegisterConsumers() {
	s.rmq.AddConsumer(
		s.cfg.ExchangeName+config.DebugConsumer, // consumerName
		s.cfg.ExchangeName,                      // exchangeName
		s.cfg.ExchangeName+config.DebugConsumer, // queueName
		s.cfg.ExchangeName+config.DebugConsumer, // routingKey
		s.debugListener,
	)

	s.rmq.AddConsumer(
		s.cfg.ExchangeName+config.InfoConsumer, // consumerName
		s.cfg.ExchangeName,                     // exchangeName
		s.cfg.ExchangeName+config.InfoConsumer, // queueName
		s.cfg.ExchangeName+config.InfoConsumer, // routingKey
		s.infoListener,
	)

	s.rmq.AddConsumer(
		s.cfg.ExchangeName+config.ErrorConsumer, // consumerName
		s.cfg.ExchangeName,                      // exchangeName
		s.cfg.ExchangeName+config.ErrorConsumer, // queueName
		s.cfg.ExchangeName+config.ErrorConsumer, // routingKey
		s.errorListener,
	)

	s.rmq.AddConsumer(
		s.cfg.ExchangeName+config.AllMessageConsumer, // consumerName
		s.cfg.ExchangeName, // exchangeName
		s.cfg.ExchangeName+config.AllMessageConsumer, // queueName
		s.cfg.ExchangeName+config.AllMessageConsumer, // routingKey
		s.allMessageListener,
	)
}
