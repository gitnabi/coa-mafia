package messanger_pkg

import (
	"context"
	"fmt"
	"log"
	"mafia/pkg/proto/player_info"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ChatMessage struct {
	PlayerInfo *player_info.PlayerInfo
	Msg        string
}

type Messanger struct {
	player_info *player_info.PlayerInfo
	conn        *amqp.Connection
	ch          *amqp.Channel
	queue       amqp.Queue
	ctx         context.Context
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func (m *Messanger) ConnectionToBroker() {
	var err error
	m.conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/%2f")
	failOnError(err, "Failed to connect to RabbitMQ")
}

func (m *Messanger) CreateChannel() {
	var err error
	m.ch, err = m.conn.Channel()
	failOnError(err, "Failed to open a channel")
}

func (m *Messanger) QueueDeclare() {
	var err error
	m.queue, err = m.ch.QueueDeclare(
		m.player_info.Uuid, // name
		true,               // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Failed to declare a queue")
}

func (m *Messanger) ExchangeDeclare() {
	err := m.ch.ExchangeDeclare(
		m.player_info.SessionId, // exchange
		"fanout",                // kind (send without binding keys; do not need them)
		true,                    // durable (do not store sent and received messages after session end)
		false,                   // autoDelete
		false,                   // internal
		false,                   // noWait (sync wait declaration)
		nil,                     // args
	)
	failOnError(err, "Failed to declare exchange")
}

func (m *Messanger) QueueBind() {
	err := m.ch.QueueBind(
		m.player_info.Uuid,      // name
		"",                      // routing key (fanout ignores routing key)
		m.player_info.SessionId, // exchange
		false,                   // noWait (sync wait declaration)
		nil,                     // args
	)
	failOnError(err, "Failed queue bind")
}

func (m *Messanger) Init(player_info *player_info.PlayerInfo) {
	m.player_info = player_info
	m.ctx = context.Background()
	m.ConnectionToBroker()
	m.CreateChannel()
	m.ExchangeDeclare()
	m.QueueDeclare()
	m.QueueBind()
}

func (m *Messanger) Close() {
	m.conn.Close()
	m.ch.Close()
}
