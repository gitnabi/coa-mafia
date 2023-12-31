package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	action_pkg "mafia/pkg/proto/action"
	service_pkg "mafia/pkg/proto/mafia_service"
	notification_pkg "mafia/pkg/proto/notification"
	"mafia/pkg/proto/player_info"
	player_info_pkg "mafia/pkg/proto/player_info"

	"google.golang.org/grpc"
)

type MappingRoles struct {
	roles     [4]player_info_pkg.Role
	available [4]bool
}

type PlayerMetaInfo struct {
	player_info  *player_info_pkg.PlayerInfo
	number_votes int
	stream       service_pkg.Mafia_ExecuteActionServer
}

type SessionInfo struct {
	meta_info_by_uuid       map[string]PlayerMetaInfo
	number_sleeping_players int
	number_living_players   int
	is_session_over         bool
	commissioner_has_worked bool
	mafia_has_worked        bool
	is_game_started         bool
}

type server struct {
	number_of_users uint64
	mapping_roles   *MappingRoles
	sessions        map[string]SessionInfo

	service_pkg.UnimplementedMafiaServer
}

func (s *server) InitPlayer(stream *service_pkg.Mafia_ExecuteActionServer, player_info *player_info_pkg.PlayerInfo) {
	player_info.Uuid = strconv.FormatUint(s.number_of_users, 10)
	player_info.SessionId = strconv.FormatUint(s.number_of_users/4, 10)
	if _, ok := s.sessions[player_info.SessionId]; !ok {
		s.sessions[player_info.SessionId] = SessionInfo{
			meta_info_by_uuid:       map[string]PlayerMetaInfo{},
			number_sleeping_players: 0,
			number_living_players:   4,
			is_session_over:         false,
			commissioner_has_worked: false,
			mafia_has_worked:        false,
			is_game_started:         false,
		}
		s.mapping_roles = &MappingRoles{
			roles: [4]player_info_pkg.Role{
				player_info_pkg.Role_MAFIA,
				player_info_pkg.Role_COMMISSIONER,
				player_info_pkg.Role_TOWNIE,
				player_info_pkg.Role_TOWNIE,
			},
			available: [4]bool{true, true, true, true},
		}
	}

	s.sessions[player_info.SessionId].meta_info_by_uuid[player_info.Uuid] = PlayerMetaInfo{
		player_info:  player_info,
		number_votes: 0,
		stream:       *stream,
	}

	seed := time.Now().UnixMilli()
	for i := 0; i < len(s.mapping_roles.roles); i++ {
		idx := (seed + int64(i)) % int64(len(s.mapping_roles.roles))
		if s.mapping_roles.available[idx] == true {
			player_info.Role = s.mapping_roles.roles[idx]
			s.mapping_roles.available[idx] = false
			break
		}
	}

	s.number_of_users += 1
}

func (s *server) NotifyEveryoneAboutJoinedPlayer(joined_player *player_info_pkg.PlayerInfo) {
	var notification = notification_pkg.Notification{
		Type: notification_pkg.NotificationType_CONNECTION_NEW_PLAYER,
		Notifications: &notification_pkg.Notification_ConnectionNewPlayerNotification{
			ConnectionNewPlayerNotification: &notification_pkg.ConnectionNewPlayerNotification{
				Player: joined_player,
			},
		},
	}
	s.SendNotifications(&notification, joined_player.SessionId, joined_player)
}

func (s *server) SendNotifications(notification *notification_pkg.Notification, sessionId string, sender *player_info_pkg.PlayerInfo) {
	for _, player_meta_info := range s.sessions[sessionId].meta_info_by_uuid {
		if sender == nil || player_meta_info.player_info.Uuid != sender.Uuid {
			err := player_meta_info.stream.Send(notification)
			if err != nil {
				return
			}
		}
	}
}

func (s *server) FinishGame(sessionId string, has_mafia_won bool) {
	notification := notification_pkg.Notification{
		Type: notification_pkg.NotificationType_FINISH,
		Notifications: &notification_pkg.Notification_FinishNotification{
			FinishNotification: &notification_pkg.FinishNotification{
				HasMafiaWon: has_mafia_won,
			},
		},
	}
	s.SendNotifications(&notification, sessionId, nil)
	if session_info, ok := s.sessions[sessionId]; ok {
		session_info.is_session_over = true
		s.sessions[sessionId] = session_info
	}
}

func (s *server) ExecutePlayer(sessionId string) {
	curr_max := -1
	number_max := 0
	var candidate *player_info_pkg.PlayerInfo
	for _, meta_player_info := range s.sessions[sessionId].meta_info_by_uuid {
		if meta_player_info.number_votes == curr_max {
			number_max += 1
		} else if meta_player_info.number_votes > curr_max {
			curr_max = meta_player_info.number_votes
			number_max = 1
			candidate = meta_player_info.player_info
		}
	}
	if number_max > 1 {
		candidate = nil
	}

	if candidate != nil {
		if candidate.Role == player_info_pkg.Role_MAFIA {
			s.FinishGame(sessionId, false)

		} else {
			candidate.Role = player_info_pkg.Role_KILLED

			if session_info, ok := s.sessions[sessionId]; ok {
				session_info.number_living_players -= 1
				s.sessions[sessionId] = session_info
			}
			if s.sessions[sessionId].number_living_players <= 2 && s.sessions[sessionId].is_game_started {
				s.FinishGame(sessionId, true)
			}
		}

	}

	notification := notification_pkg.Notification{
		Type: notification_pkg.NotificationType_VOTING_RESULT,
		Notifications: &notification_pkg.Notification_VotingResultNotification{
			VotingResultNotification: &notification_pkg.VotingResultNotification{
				ExecutedPlayer: candidate,
			},
		},
	}
	s.SendNotifications(&notification, sessionId, nil)
}

func (s *server) GetListLivingPlayers(sessionId string) []*player_info_pkg.PlayerInfo {
	players_info := []*player_info_pkg.PlayerInfo{}
	for _, meta_player_info := range s.sessions[sessionId].meta_info_by_uuid {
		if meta_player_info.player_info.Role != player_info_pkg.Role_KILLED {
			players_info = append(players_info, meta_player_info.player_info)
		}
	}
	return players_info
}

func (s *server) WakeUpMafia(sessionId string) {
	notification := notification_pkg.Notification{
		Type: notification_pkg.NotificationType_WAKE,
		Notifications: &notification_pkg.Notification_WakeNotification{
			WakeNotification: &notification_pkg.WakeNotification{
				Remaining: s.GetListLivingPlayers(sessionId),
			},
		},
	}

	is_mafia_dead := true
	for _, player_meta_info := range s.sessions[sessionId].meta_info_by_uuid {
		if player_meta_info.player_info.Role == player_info_pkg.Role_MAFIA {
			is_mafia_dead = false
			player_meta_info.player_info.Asleep = true
			err := player_meta_info.stream.Send(&notification)
			if err != nil {
				log.Fatalf("err send: %v", err)
			}
			break
		}
	}
	if is_mafia_dead {
		if session_info, ok := s.sessions[sessionId]; ok {
			session_info.mafia_has_worked = true
			s.sessions[sessionId] = session_info
		}
	}
}

func (s *server) WakeUpCommissioner(sessionId string) {
	notification := notification_pkg.Notification{
		Type: notification_pkg.NotificationType_WAKE,
		Notifications: &notification_pkg.Notification_WakeNotification{
			WakeNotification: &notification_pkg.WakeNotification{
				Remaining: s.GetListLivingPlayers(sessionId),
			},
		},
	}

	is_commissioner_dead := true
	for _, player_meta_info := range s.sessions[sessionId].meta_info_by_uuid {
		if player_meta_info.player_info.Role == player_info_pkg.Role_COMMISSIONER {
			is_commissioner_dead = false
			err := player_meta_info.stream.Send(&notification)
			if err != nil {
				log.Fatalf("err send: %v", err)
			}
			break
		}
	}

	if is_commissioner_dead {
		if session_info, ok := s.sessions[sessionId]; ok {
			session_info.commissioner_has_worked = true
			s.sessions[sessionId] = session_info
		}
	}
}

func (s *server) WakeUpAll(sessionId string, killed_player *player_info_pkg.PlayerInfo) {
	for s.sessions[sessionId].commissioner_has_worked == false || s.sessions[sessionId].mafia_has_worked == false {
		time.Sleep(100 * time.Millisecond)
	}

	notification := notification_pkg.Notification{
		Type: notification_pkg.NotificationType_WAKE,
		Notifications: &notification_pkg.Notification_WakeNotification{
			WakeNotification: &notification_pkg.WakeNotification{
				Remaining:    s.GetListLivingPlayers(sessionId),
				KilledPlayer: killed_player,
			},
		},
	}

	for _, player_meta_info := range s.sessions[sessionId].meta_info_by_uuid {
		err := player_meta_info.stream.Send(&notification)
		if err != nil {
			log.Fatalf("err send: %v", err)
		}
	}
}

func (s *server) Disconnect(sender *player_info.PlayerInfo) {
	for i, role := range s.mapping_roles.roles {
		if role == sender.Role {
			s.mapping_roles.available[i] = true
			break
		}
	}
	s.number_of_users -= 1
	if session_info, ok := s.sessions[sender.SessionId]; ok {
		session_info.number_living_players -= 1
		if sender.Asleep {
			session_info.number_sleeping_players -= 1
		}
		if sender.Role == player_info.Role_MAFIA {
			s.FinishGame(sender.SessionId, false)
			os.Exit(0)
		} else if session_info.number_living_players <= 2 {
			s.FinishGame(sender.SessionId, true)
			os.Exit(0)
		}
		delete(session_info.meta_info_by_uuid, sender.Uuid)
		s.sessions[sender.SessionId] = session_info
	}
}

func (s *server) ExecuteAction(stream service_pkg.Mafia_ExecuteActionServer) error {
	if s.number_of_users == 0 {
		s.sessions = map[string]SessionInfo{}
	}

	action, err := stream.Recv()
	if err != nil {
		return err
	}
	if action.Type != action_pkg.ActionType_INIT {
		log.Fatalf("Первое действие не инициализирующее!")
	}

	player := action.GetInitAction().GetPlayer()
	s.InitPlayer(&stream, player)
	fmt.Printf("recv init action: %s\n", player.ToString())
	s.NotifyEveryoneAboutJoinedPlayer(player)

	if s.number_of_users%4 == 0 && !s.sessions[player.SessionId].is_game_started {
		if session_info, ok := s.sessions[player.SessionId]; ok {
			session_info.is_game_started = true
			s.sessions[player.SessionId] = session_info
		}
		for _, player_meta_info := range s.sessions[player.SessionId].meta_info_by_uuid {
			notification := notification_pkg.Notification{
				Type: notification_pkg.NotificationType_START_GAME,
				Notifications: &notification_pkg.Notification_StartGameNotification{
					StartGameNotification: &notification_pkg.StartGameNotification{
						Player: player_meta_info.player_info,
					},
				},
			}
			err := player_meta_info.stream.Send(&notification)
			if err != nil {
				log.Fatalf("err send: %v", err)
			}
		}
	}

	for s.sessions[player.SessionId].is_session_over == false {
		action, err := stream.Recv()
		if err != nil {
			return err
		}

		switch action.Type {
		case action_pkg.ActionType_CHECK_ROLE:
			fmt.Println("recv check role action")
			checked_player := action.GetCheckAction().GetPlayer()
			fmt.Printf("Проверили %s, %s\n", checked_player.ToString(), checked_player.Role.ToString())
			if session_info, ok := s.sessions[player.SessionId]; ok {
				session_info.commissioner_has_worked = true
				s.sessions[player.SessionId] = session_info
			}

		case action_pkg.ActionType_VOTE:
			fmt.Println("recv vote action")
			candidate := action.GetVoteAction().GetPlayer()
			if meta_info, ok := s.sessions[candidate.SessionId].meta_info_by_uuid[candidate.Uuid]; ok {
				meta_info.number_votes += 1
				s.sessions[player.SessionId].meta_info_by_uuid[candidate.Uuid] = meta_info
			}
			fmt.Printf("voter - %s, candidate - %s\n", player.ToString(), candidate.ToString())
			var notification = notification_pkg.Notification{
				Type: notification_pkg.NotificationType_VOTING,
				Notifications: &notification_pkg.Notification_VotingNotification{
					VotingNotification: &notification_pkg.VotingNotification{
						Voter:     player,
						Candidate: candidate,
					},
				},
			}
			s.SendNotifications(&notification, player.SessionId, player)

		case action_pkg.ActionType_KILL:
			fmt.Println("recv kill action")
			killed_player := action.GetKillAction().GetPlayer()
			fmt.Printf("Убивают %s, %s\n", killed_player.ToString(), killed_player.Role.ToString())
			s.sessions[killed_player.GetSessionId()].meta_info_by_uuid[killed_player.GetUuid()].player_info.Role = player_info_pkg.Role_KILLED

			if session_info, ok := s.sessions[killed_player.GetSessionId()]; ok {
				session_info.number_living_players -= 1
				session_info.mafia_has_worked = true
				s.sessions[killed_player.GetSessionId()] = session_info
			}
			s.WakeUpAll(player.SessionId, killed_player)
			if s.sessions[killed_player.GetSessionId()].number_living_players <= 2 {
				s.FinishGame(player.GetSessionId(), true)
			}

		case action_pkg.ActionType_PUBLISH_ROLE:
			fmt.Println("recv publish role action")
			published_player := action.GetPublishRoleAction().GetPlayer()
			fmt.Printf("Раскрыли %s, %s\n", published_player.ToString(), published_player.Role.ToString())
			var notification = notification_pkg.Notification{
				Type: notification_pkg.NotificationType_ROLE_PUBLICATION,
				Notifications: &notification_pkg.Notification_RolePublicationNotification{
					RolePublicationNotification: &notification_pkg.RolePublicationNotification{
						Player: published_player,
					},
				},
			}
			s.SendNotifications(&notification, player.SessionId, player)

		case action_pkg.ActionType_GO_TO_SLEEP:
			fmt.Println("recv go to sleep")
			fmt.Printf("заснул %s\n", player.ToString())
			player.Asleep = true
			if session_info, ok := s.sessions[player.SessionId]; ok {

				session_info.number_sleeping_players += 1
				if session_info.number_sleeping_players == session_info.number_living_players {
					player.Asleep = false
					session_info.number_sleeping_players = 0
					session_info.commissioner_has_worked = false
					session_info.mafia_has_worked = false
					s.sessions[player.SessionId] = session_info

					if session_info.number_living_players != 4 {
						s.ExecutePlayer(player.SessionId)
					}
					if s.sessions[player.GetSessionId()].is_session_over == true {
						return nil
					}
					s.WakeUpMafia(player.SessionId)
					s.WakeUpCommissioner(player.SessionId)
					session_info = s.sessions[player.SessionId]
				}

				s.sessions[player.SessionId] = session_info
			}
		case action_pkg.ActionType_DISCONNECT:
			fmt.Printf("recv disconnect: %s\n", player.ToString())
			s.Disconnect(player)
		}

	}
	return nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	service_pkg.RegisterMafiaServer(s, &server{number_of_users: 0})
	fmt.Println("Сервер запущен...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
