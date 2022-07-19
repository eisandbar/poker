package typing

type Card struct {
	Value int
	Suit  int
}

func (card Card) ToString() string {

}

func (card Card) ToInt() int {

}

func (card *Card) FromString(s string) Card {

}

func (card *Card) FromInt(input int) {

}
