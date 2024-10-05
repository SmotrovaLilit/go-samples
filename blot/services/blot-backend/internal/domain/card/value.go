package card

import "fmt"

var (
	ValueSeven = Value{"seven"}
	ValueEight = Value{"eight"}
	ValueNine  = Value{"nine"}
	ValueTen   = Value{"ten"}
	ValueJack  = Value{"jack"}
	ValueDame  = Value{"dame"}
	ValueQueen = Value{"queen"}
	ValueKing  = Value{"king"}
	ValueAce   = Value{"ace"}

	Values = []Value{ValueSeven, ValueEight, ValueNine, ValueTen, ValueJack, ValueDame, ValueQueen, ValueKing, ValueAce}
)

type Value struct {
	value string
}

func NewValue(valueString string) Value {
	for _, value := range Values {
		if value.value == valueString {
			return value
		}
	}
	panic(fmt.Sprintf("Invalid value: %s", valueString))
}
