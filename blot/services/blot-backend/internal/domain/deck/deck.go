package deck

import (
	"math/rand"

	"gosamples/blot/internal/domain/card"
	"gosamples/blot/internal/domain/player"
)

type Deck struct {
	cards [32]card.Card
}

func NewDeck() Deck {
	deck := Deck{}
	i := 0
	for _, suit := range card.Suits {
		for _, value := range card.Values {
			deck.cards[i] = card.NewCard(value, suit)
			i++
		}
	}
	return deck
}

func (d *Deck) Shuffle() {
	rand.Shuffle(32, func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *Deck) DealCards(players []player.Player) {
	// TODO может ли он менять чужой агрегат?
	for i := 0; i < 8; i++ {
		for j := 0; j < 4; j++ {
			players[j].AddCard(d.cards[i*4+j])
		}
	}
}
