package game

import (
	"gosamples/blot/internal/domain/card"
	"gosamples/blot/internal/domain/team"
)

type Game struct {
	teams     [2]team.Team
	trumpSuit card.Suit
}
