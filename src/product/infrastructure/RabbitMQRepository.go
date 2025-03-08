package repositories

import (
	"encoding/json"
	"publisher/src/core"
	"publisher/src/product/domain"
)

type RabbitMQRepository struct {
	Rabbit *core.RabbitMQ
}

func (r *RabbitMQRepository) PublishMessage(message domain.Message) error {
	messageData, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return r.Rabbit.PublishMessage(messageData)
}
