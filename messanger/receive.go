package messanger_pkg

import (
	"encoding/json"
	"log"
)

func (m *Messanger) ReceiveMsgs() {
	msgs, err := m.ch.Consume(
		m.player_info.Uuid, // queue
		"",                 // consumer
		true,               // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		var chatMsg ChatMessage
		for msg := range msgs {
			if err := json.Unmarshal(msg.Body, &chatMsg); err != nil {
				failOnError(err, "Failed unmarshal msg")
				return
			}
			log.Printf("Сообщение от %s: %s", chatMsg.PlayerInfo.ToString(), chatMsg.Msg)
		}
	}()
}
