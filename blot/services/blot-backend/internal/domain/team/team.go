package team

import "gosamples/blot/internal/domain/player"

type Team struct {
	name    string
	players [2]player.Player
}

func NewTeam(name string, player1 player.Player, player2 player.Player) Team {
	return Team{name, [2]player.Player{player1, player2}}
}
