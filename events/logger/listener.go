package consumer

import (
	"github.com/saidamir98/udevs_pkg/logger"
	"github.com/streadway/amqp"
)

// debugListener ...
func (s *Logger) debugListener(delivery amqp.Delivery) {
	s.log.Debug("debugListener, new message: ", logger.String("Data", string(delivery.Body)), logger.Any("Message", delivery.Headers["message"]))
}

// errorListener ...
func (s *Logger) errorListener(delivery amqp.Delivery) {
	s.log.Info("errorListener, new message: ", logger.String("Data", string(delivery.Body)), logger.Any("Message", delivery.Headers["message"]))
}

// infoListener ...
func (s *Logger) infoListener(delivery amqp.Delivery) {
	s.log.Info("infoListener, new message: ", logger.String("Data", string(delivery.Body)), logger.Any("Message", delivery.Headers["message"]))
}

// allMessageListener ...
func (s *Logger) allMessageListener(delivery amqp.Delivery) {
	s.log.Info("allMessageListener, new message: ", logger.String("Data", string(delivery.Body)), logger.Any("Message", delivery.Headers["message"]))
}
