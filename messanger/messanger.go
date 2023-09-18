package messenger

import (
	"context"
	"fmt"
	"log"
	"time"

	// github.com/streadway/amqp --- deprecated
	service_pkg "mafia/pkg/proto/mafia_service"
	player_info_pkg "mafia/pkg/proto/player_info"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MafiaMessenger struct {
	messenger *Messenger

	client      *service_pkg.MafiaClient
	player_info *player_info_pkg.PlayerInfo
	ctx         context.Context
}

func NewMafiaMessenger(client *service_pkg.MafiaClient, player_info *player_info_pkg.PlayerInfo) *MafiaMessenger {
	return &MafiaMessenger{
		NewMessenger(fmt.Sprintf("%d", player_info.SessionId)),
		client,
		player_info,
		context.Background(),
	}
}

func (m *MafiaMessenger) Send(msg string) error {
	msg = fmt.Sprintf("%s: %s", m.player_info.ToString(), msg)

	m.messenger.Send(msg)
	return nil
}

func (m *MafiaMessenger) Receive() chan string {
	return m.messenger.Receive()
}

func (m *MafiaMessenger) Close() {
	m.messenger.Close()
}

type Messenger struct {
	chat_id     string
	sentMsg     chan string
	receivedMsg chan string
}

func NewMessenger(chat_id string) *Messenger {
	messenger := Messenger{
		chat_id,
		make(chan string),
		make(chan string),
	}

	go messenger.publish()
	go messenger.consume()

	return &messenger
}

func (m *Messenger) Send(msg string) {
	m.sentMsg <- msg
}

func (m *Messenger) Receive() chan string {
	return m.receivedMsg
}

func (m *Messenger) Close() {
	close(m.sentMsg)
	close(m.receivedMsg)
}

func (m *Messenger) publish() {
	const rabbitmq_url = "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(rabbitmq_url)
	for err != nil {
		log.Println(err)
		conn, err = amqp.Dial(rabbitmq_url)
		time.Sleep(200 * time.Millisecond)
	}
	defer conn.Close()

	log.Println("Подключились к RabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}
	defer ch.Close()

	// ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args Table)
	err = ch.ExchangeDeclare(
		m.chat_id, "fanout", true, false, false, false, nil,
	)
	if err != nil {
		log.Println(err)
	}

	for msg := range m.sentMsg {
		ctx := context.Background()

		// PublishWithContext(ctx context.Context, exchange, key string, mandatory, immediate bool, msg Publishing)
		err := ch.PublishWithContext(ctx,
			m.chat_id, "", false, false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg),
			})

		if err != nil {
			log.Println(err)
		}
	}
}

func (m *Messenger) consume() {
	const rabbitmq_url = "amqp://guest:guest@localhost:5672"
	conn, err := amqp.Dial(rabbitmq_url)
	for err != nil {
		log.Println(err)
		conn, err = amqp.Dial(rabbitmq_url)
		time.Sleep(200 * time.Millisecond)
	}
	defer conn.Close()

	log.Println("Подключились к RabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}
	defer ch.Close()

	// ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args Table)
	err = ch.ExchangeDeclare(
		m.chat_id, "fanout", true, false, false, false, nil,
	)
	if err != nil {
		log.Println(err)
	}

	// QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args Table) (Queue, error)
	queue, err := ch.QueueDeclare(
		"", false, false, true, false, nil,
	)
	if err != nil {
		log.Println(err)
	}

	// QueueBind(name, key, exchange string, noWait bool, args Table)
	err = ch.QueueBind(
		queue.Name, "", m.chat_id, false, nil,
	)
	if err != nil {
		log.Println(err)
	}

	msgs, err := ch.Consume(
		queue.Name, "", true, false, false, false, nil,
	)
	if err != nil {
		log.Println(err)
	}

	for msg := range msgs {
		m.receivedMsg <- string(msg.Body)
		log.Printf("Recv: %s", msg.Body)
	}
}
