package player_info

import (
	"fmt"
)

func (role *Role) ToString() string {
	switch *role {
	case Role_MAFIA:
		return "Мафия"
	case Role_COMMISSIONER:
		return "Комиссар"
	case Role_TOWNIE:
		return "Горожанин"
	case Role_KILLED:
		return "Дух"
	}
	return role.String()
}

func (player_info *PlayerInfo) ToString() string {
	return fmt.Sprintf("[name=\"%s\"; uuid=%s; session_id=%d]", player_info.Name, player_info.Uuid, player_info.SessionId)
}
