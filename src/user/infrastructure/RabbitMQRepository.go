package repositories

import (
	"encoding/json"
	"publisher/src/core"
	"publisher/src/user/domain"
)

type RabbitMQRepository struct {
	Rabbit *core.RabbitMQ
}

func (r *RabbitMQRepository) PublishUser(message domain.Message) error {
	userData, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return r.Rabbit.PublishMessage(userData)
}

func (r *RabbitMQRepository) PublishLogin(message domain.Message) error {
	messageData, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return r.Rabbit.PublishMessage(messageData)
}