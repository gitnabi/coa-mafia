package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	action_pkg "mafia/pkg/proto/action"
	service_pkg "mafia/pkg/proto/mafia_service"
	notification_pkg "mafia/pkg/proto/notification"
	player_info_pkg "mafia/pkg/proto/player_info"
	"os"

	"google.golang.org/grpc"
)

type ActionType int32

const (
	SEND_MESSAGE ActionType = 0
	VOTE         ActionType = 1
	GO_TO_SLEEP  ActionType = 3
	CHECK_ROLE   ActionType = 4
	PUBLISH_ROLE ActionType = 5
	KILL_PLAYER  ActionType = 6
)

type client struct {
	conn           *grpc.ClientConn
	client         service_pkg.MafiaClient
	stream         service_pkg.Mafia_ExecuteActionClient
	player_info    *player_info_pkg.PlayerInfo
	ctx            context.Context
	checked_player *player_info_pkg.PlayerInfo
}

func (c *client) InitAndRun(port string) {
	var err error

	c.conn, err = grpc.Dial(":"+port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed connect: %v", err)
	}

	c.client = service_pkg.NewMafiaClient(c.conn)
	fmt.Println("Клиент запущен...")

	c.ctx = context.Background()
	c.stream, err = c.client.ExecuteAction(c.ctx)
	if err != nil {
		log.Fatalf("err creating stream: %s", err.Error())
	}
	fmt.Println("Стрим успешно создан...")

	fmt.Print("Введите свое имя: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	c.player_info = &player_info_pkg.PlayerInfo{
		Name: name[:len(name)-1],
	}
}

func (c *client) Close() {
	c.conn.Close()
	c.stream.CloseSend()
}

func (c *client) sendАction(action *action_pkg.Action) {
	err := c.stream.Send(action)
	if err != nil {
		log.Fatalf("err send: %s", err.Error())
	}
}

func (c *client) receiveNotification() *notification_pkg.Notification {
	notification, err := c.stream.Recv()
	if err != nil {
		return nil
	}
	return notification
}

func (c *client) PreparingForGame() {
	fmt.Println("Ожидайте начала игры!")

	action := action_pkg.Action{
		Type: action_pkg.ActionType_INIT,
		Actions: &action_pkg.Action_InitAction{
			InitAction: &action_pkg.InitAction{
				Player: c.player_info,
			},
		},
	}
	c.sendАction(&action)

	var notification *notification_pkg.Notification
	for {
		notification = c.receiveNotification()
		switch notification.Type {
		case notification_pkg.NotificationType_CONNECTION_NEW_PLAYER:
			fmt.Printf("Подключился новый игрок! %s\n", notification.GetConnectionNewPlayerNotification().Player.ToString())
		case notification_pkg.NotificationType_START_GAME:
			fmt.Println("------------------------")
			fmt.Println("\t\tИгра начинается")
			c.player_info = notification.GetStartGameNotification().GetPlayer()
			fmt.Printf("Ваша роль: %s\n", c.player_info.Role.ToString())
			return
		default:
			log.Fatalf("Получено неизвестное уведомление!")
		}
	}
}

func (c *client) selectAnAction(available_actions *[]ActionType) int {
	c.printAvailableActions(*available_actions)

	var result int
	_, err := fmt.Scanf("%d", &result)
	for err != nil || result < 0 || result >= len(*available_actions) {
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Неправильно введен номер действия. Повторите попытку!")
		_, err = fmt.Scanf("%d", &result)
	}

	return result
}

func (c *client) selectPlayer(elector *player_info_pkg.PlayerInfo, candidates []*player_info_pkg.PlayerInfo) *player_info_pkg.PlayerInfo {
	fmt.Println("Выберите одного игрока из списка(введите его номер): ")
	for idx, player_info := range candidates {
		if player_info.Uuid != elector.Uuid {
			fmt.Printf("%d. %s\n", idx, player_info.ToString())
		}
	}
	var idx_player int
	_, err := fmt.Scanf("%d", &idx_player)
	for err != nil || idx_player < 0 || idx_player >= len(candidates) || candidates[idx_player].Uuid == elector.Uuid {
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Неправильно введен номер игрока. Повторите попытку!")
		_, err = fmt.Scanf("%d", &idx_player)
	}
	return candidates[idx_player]
}

func (c *client) printAvailableActions(action_types []ActionType) {
	fmt.Println("Выберите одно из действий(введите его номер): ")
	for idx, action_type := range action_types {
		switch action_type {
		case SEND_MESSAGE:
			fmt.Printf("%d. Отправить сообщение остальным участникам игры.\n", idx)
		case VOTE:
			fmt.Printf("%d. Проголосовать за мафию.\n", idx)
		case GO_TO_SLEEP:
			fmt.Printf("%d. Завершить день.\n", idx)
		case CHECK_ROLE:
			fmt.Printf("%d. Проверить роль игрока.\n", idx)
		case PUBLISH_ROLE:
			fmt.Printf("%d. Раскрыть роль игрока, проверенного ночью.\n", idx)
		case KILL_PLAYER:
			fmt.Printf("%d. Убить игрока.\n", idx)
		default:
			log.Fatalf("Неизвестное действие!")
		}
	}
}

func (c *client) ListenNotifications(ch chan *notification_pkg.Notification) {
	var notification *notification_pkg.Notification
	for {
		notification = c.receiveNotification()
		if notification == nil {
			break
		}
		switch notification.Type {
		case notification_pkg.NotificationType_RECEIVE_MESSAGE:
			receive_message_notification := notification.GetReceiveMessageNotification()
			fmt.Printf("От %s получено следующее сообщение: %s", receive_message_notification.Sender.ToString(), receive_message_notification.Message)
		case notification_pkg.NotificationType_VOTING:
			voter := notification.GetVotingNotification().GetVoter()
			candidate := notification.GetVotingNotification().GetCandidate()
			fmt.Printf("Игрок %s проголосовал за игрока %s\n", voter.ToString(), candidate.ToString())
		case notification_pkg.NotificationType_ROLE_PUBLICATION:
			fmt.Println("Комиссар проверил одного из игроков и решил раскрыть его роль!")
			player := notification.GetRolePublicationNotification().Player
			fmt.Printf("У игрока %s роль %s\n", player.ToString(), player.Role.ToString())
		case notification_pkg.NotificationType_WAKE:
			ch <- notification
		case notification_pkg.NotificationType_FINISH:
			fmt.Println("------------------------")
			fmt.Println("Игра окончена!")
			if notification.GetFinishNotification().GetHasMafiaWon() {
				fmt.Println("Мафия победила, т.к. осталось два игрока!")
			} else {
				fmt.Println("Мафия проиграла, т.к. ее убили!")
			}
			ch <- notification
		case notification_pkg.NotificationType_VOTING_RESULT:
			executed_player := notification.GetVotingResultNotification().GetExecutedPlayer()
			if executed_player != nil {
				if executed_player.Uuid != c.player_info.Uuid {
					fmt.Printf("По результатам голосования казнили %s.\n", executed_player.ToString())
				} else {
					executed_player.Role = player_info_pkg.Role_KILLED
					fmt.Println("------------------------")
					fmt.Println("По результатам голосования вас казнили.")
					fmt.Println("------------------------")
				}
			} else {
				fmt.Println("У нескольких игроков равное количество голосов, поэтому никого не казнили.")
			}
		default:
			log.Fatalf("Получено неизвестное уведомление! %s", notification.Type.String())
		}
	}
}

func (c *client) GameDay(available_actions []ActionType, wake_notification *notification_pkg.WakeNotification) {
	fmt.Println("------------------------")
	fmt.Print("***\tДень\t***\n\n")

	if wake_notification != nil && wake_notification.KilledPlayer != nil {
		if wake_notification.KilledPlayer.Uuid == c.player_info.Uuid {
			fmt.Println("------------------------")
			fmt.Println("Ночью вас убила мафия!")
			fmt.Println("------------------------")
			return
		} else {
			fmt.Printf("Ночью мафия убила %s!\n", wake_notification.KilledPlayer.ToString())
		}
	}

	if wake_notification != nil && len(wake_notification.Remaining) <= 2 {
		return
	}

	for {
		idx_action := c.selectAnAction(&available_actions)
		switch available_actions[idx_action] {
		case SEND_MESSAGE:
			fmt.Println("Введите сообщение, которое хотите отправить:")
			reader := bufio.NewReader(os.Stdin)
			message, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("err read string: %s", err)
			}

			action := action_pkg.Action{
				Type: action_pkg.ActionType_SEND_MESSAGE,
				Actions: &action_pkg.Action_SendMessageAction{
					SendMessageAction: &action_pkg.SendMessageAction{
						Message: message,
					},
				},
			}
			c.sendАction(&action)

		case VOTE:
			selected := c.selectPlayer(c.player_info, wake_notification.Remaining)
			action := action_pkg.Action{
				Type: action_pkg.ActionType_VOTE,
				Actions: &action_pkg.Action_VoteAction{
					VoteAction: &action_pkg.VoteAction{
						Player: selected,
					},
				},
			}
			c.sendАction(&action)
			available_actions[idx_action] = available_actions[len(available_actions)-1]
			available_actions = available_actions[:len(available_actions)-1]
		case GO_TO_SLEEP:
			action := action_pkg.Action{
				Type: action_pkg.ActionType_GO_TO_SLEEP,
				Actions: &action_pkg.Action_GoToSleepAction{
					GoToSleepAction: &action_pkg.GoToSleepAction{},
				},
			}
			c.sendАction(&action)
			fmt.Println("Вы заснули! Ожидайте пробуждения!")
			return
		case PUBLISH_ROLE:
			action := action_pkg.Action{
				Type: action_pkg.ActionType_PUBLISH_ROLE,
				Actions: &action_pkg.Action_PublishRoleAction{
					PublishRoleAction: &action_pkg.PublishRoleAction{
						Player: c.checked_player,
					},
				},
			}
			c.sendАction(&action)
			available_actions[idx_action] = available_actions[len(available_actions)-1]
			available_actions = available_actions[:len(available_actions)-1]
		default:
			log.Fatalf("Недоступное или неизвестное действие!")
		}
	}
}

func (c *client) GameNight(wake_notification *notification_pkg.WakeNotification) {
	fmt.Println("------------------------")
	fmt.Print("***\tНочь\t***\n\n")

	switch c.player_info.Role {
	case player_info_pkg.Role_MAFIA:
		fmt.Println("Мафии необходимо убить одного игрока.")
		selected := c.selectPlayer(c.player_info, wake_notification.Remaining)
		action := action_pkg.Action{
			Type: action_pkg.ActionType_KILL,
			Actions: &action_pkg.Action_KillAction{
				KillAction: &action_pkg.KillAction{
					Player: selected,
				},
			},
		}
		c.sendАction(&action)
	case player_info_pkg.Role_COMMISSIONER:
		fmt.Println("Комиссару необходимо проверить одного игрока.")
		selected := c.selectPlayer(c.player_info, wake_notification.Remaining)
		action := action_pkg.Action{
			Type: action_pkg.ActionType_CHECK_ROLE,
			Actions: &action_pkg.Action_CheckAction{
				CheckAction: &action_pkg.CheckRoleAction{
					Player: selected,
				},
			},
		}
		c.sendАction(&action)
		c.checked_player = selected
		fmt.Printf("У игрока %s роль %s\n", selected.ToString(), selected.Role.ToString())
	default:
		log.Fatalln("Проснулся игрок с невалидной ролью!")
	}
	fmt.Println("Вы заснули! Ожидайте пробуждения!")
}

func (c *client) WaitNotification(ch chan *notification_pkg.Notification) *notification_pkg.WakeNotification {
	notification := <-ch
	if notification.Type == notification_pkg.NotificationType_FINISH {
		return nil
	} else if notification.Type != notification_pkg.NotificationType_WAKE {
		log.Fatalf("Получено неожиданное уведомление! %s\n", notification.Type.String())
	}
	return notification.GetWakeNotification()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	var c client
	c.InitAndRun("9000")
	defer c.Close()

	c.PreparingForGame()

	ch := make(chan *notification_pkg.Notification, 1)
	go c.ListenNotifications(ch)

	daily_available_actions := []ActionType{SEND_MESSAGE, GO_TO_SLEEP}
	c.GameDay(daily_available_actions, nil)
	wake_notification := c.WaitNotification(ch)

	is_night_role := c.player_info.Role == player_info_pkg.Role_COMMISSIONER || c.player_info.Role == player_info_pkg.Role_MAFIA
	if is_night_role {
		c.GameNight(wake_notification)
		wake_notification = c.WaitNotification(ch)
	}

	daily_available_actions = append(daily_available_actions, VOTE)
	if c.player_info.Role == player_info_pkg.Role_COMMISSIONER {
		daily_available_actions = append(daily_available_actions, PUBLISH_ROLE)
	}

	for wake_notification != nil {
		if wake_notification.KilledPlayer.Uuid == c.player_info.Uuid {
			c.player_info = wake_notification.KilledPlayer
			daily_available_actions = []ActionType{}
			is_night_role = false
		}

		c.GameDay(daily_available_actions, wake_notification)

		wake_notification = c.WaitNotification(ch)
		if wake_notification == nil {
			break
		}

		if is_night_role {
			c.GameNight(wake_notification)
			wake_notification = c.WaitNotification(ch)
		}
	}
}
