package player

import (
	"gosamples/blot/internal/domain/card"
)

type Player struct {
	name       string
	cards      []card.Card
	trickCards []card.Card
}

func NewPlayer(name string) Player {
	return Player{name, []card.Card{}, []card.Card{}}
}

func (p Player) AddCard(c card.Card) {
	p.cards = append(p.cards, c)
}

func (p Player) AddTrickCards(trick [4]card.Card) {
	p.trickCards = append(p.trickCards, trick[:]...)
}
