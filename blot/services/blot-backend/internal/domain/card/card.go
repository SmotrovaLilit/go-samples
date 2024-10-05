package card

type Card struct {
	value Value
	suit  Suit
}

func NewCard(value Value, suit Suit) Card {
	return Card{value, suit}
}
