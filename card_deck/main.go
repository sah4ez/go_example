package card_deck

import (
	"math/rand"
	"sort"
)

const (
	CLUBS    = 1
	DIAMONDS = 2
	HEARTS   = 3
	SPADES   = 4
	SIX      = 6
	SEVEN    = 7
	EIGHT    = 8
	NINE     = 9
	TEN      = 10
	JACK     = 11
	QUEEN    = 12
	KING     = 13
	ACE      = 14
	WJOCKER  = 15
	BJOCKER  = 16
)

var (
	cards   = [9]int{ACE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING}
	jockers = [2]int{WJOCKER, BJOCKER}
	SINGS   = map[int]string{
		0:  "",
		1:  "♣",
		2:  "♦",
		3:  "♥",
		4:  "♠",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "10",
		11: "J",
		12: "Q",
		13: "K",
		14: "A",
		15: "WJ",
		16: "BJ",
	}
)

type CardDeck interface {
	SortWithParameter(trump int, order [4]int) *Deck
	Shuffle() *Deck
}

type Deck struct {
	cards []Card
	suits [4]int
}

type Card struct {
	suit int
	value int
}

func NewDeck() Deck {
	suits := [4]int{CLUBS, DIAMONDS, HEARTS, SPADES}
	deck := Deck{}
	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(cards); j++ {
			deck.cards = append(deck.cards, Card{value: cards[j], suit: suits[i]})
		}
	}
	deck.cards = append(deck.cards, Card{value: WJOCKER, suit: 5})
	deck.cards = append(deck.cards, Card{value: BJOCKER, suit: 5})
	return deck
}

func (d Deck) SortWithParameter(trump int, order [4]int) {
	d.suits = order
	var suit int
	pos := 0
	for n, s := range d.suits {
		if s == trump {
			pos = n
			suit = s
			break
		}
	}
	for ; pos != 0; pos-- {
		d.suits[pos] = d.suits[pos-1]
	}
	d.suits[0] = suit
	d.Sort()
}

func (d Deck) Sort()         { sort.Sort(d) }
func (d Deck) Len() int      { return len(d.cards) }
func (d Deck) Swap(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] }
func (d Deck) Less(i, j int) bool {
	if d.cards[i].suit == d.cards[j].suit {
		return d.cards[i].value > d.cards[j].value
	}
	var s1 int
	var s2 int
	for n, s := range d.suits {
		if d.cards[i].suit == s {
			s1 = n
		}
		if d.cards[j].suit == s {
			s2 = n
		}
		if d.cards[i].suit == 5 {
			s1 = -1
		}
		if d.cards[j].suit == 5 {
			s2 = -1
		}
	}
	return s1 < s2
}

func (d Deck) Shuffle() Deck {
	for i := 1; i < len(d.cards); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			d.cards[r], d.cards[i] = d.cards[i], d.cards[r]
		}
	}
	return d
}
