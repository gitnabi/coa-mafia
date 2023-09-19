package messanger_pkg

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (m *Messanger) SendMsg(chatMst *ChatMessage) {
	msgBytes, err := json.Marshal(chatMst)
	if err != nil {
		failOnError(err, "Failed marshal msg")
	}
	err = m.ch.PublishWithContext(
		m.ctx,
		m.player_info.SessionId, // exchange
		"",                      // routing key
		false,                   // mandatory
		false,                   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msgBytes,
		})
	failOnError(err, "Failed to publish a message")
}
