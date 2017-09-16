package card_deck

import (
	"testing"
	"fmt"
)

func TestDeck_Shuffle(t *testing.T) {
	notShuffleDeck := NewDeck()
	deck := NewDeck()
	deck.Shuffle()
	order := make(chan int, 38)
	defer close(order)
	for n, c := range deck.cards {
		if c.suit == notShuffleDeck.cards[n].suit && c.value == notShuffleDeck.cards[n].value {
			order <- 1
		}
	}
	if len(order) == 38 {
		t.Fatalf("Not shuffled deck")
	}
}

func TestDeck_SortWithParameter(t *testing.T) {
	deck := NewDeck()
	deck.Shuffle()
	deck.SortWithParameter(DIAMONDS, [4]int{CLUBS, HEARTS, DIAMONDS, SPADES})
	if deck.cards[0].value != BJOCKER {
		t.Fatalf("Second card in the deck should be Black Jocker.")
	}
	if deck.cards[1].value != WJOCKER {
		t.Fatalf("First card in the deck should be White Jocker.")
	}
	p := 2
	if deck.cards[p].value != ACE || deck.cards[p].suit != DIAMONDS {
		t.Fatalf("%dth card in the deck should be ACE %s, actual %s%s",
			p, SINGS[DIAMONDS], SINGS[deck.cards[p].value], SINGS[deck.cards[p].suit])
	}
	p = 11
	if deck.cards[p].value != ACE || deck.cards[p].suit != CLUBS {
		t.Fatalf("%dth card in the deck should be ACE %s, actual %s%s",
			p, SINGS[CLUBS], SINGS[deck.cards[p].value], SINGS[deck.cards[p].suit])
	}
	p = 20
	if deck.cards[p].value != ACE || deck.cards[p].suit != HEARTS {
		t.Fatalf("%dth card in the deck should be ACE %s, actual %s%s",
			p, SINGS[HEARTS], SINGS[deck.cards[p].value], SINGS[deck.cards[p].suit])
	}
	p = 29
	if deck.cards[p].value != ACE || deck.cards[p].suit != SPADES {
		t.Fatalf("%dth card in the deck should be ACE %s, actual %s%s",
			p, SINGS[SPADES], SINGS[deck.cards[p].value], SINGS[deck.cards[p].suit])
	}
	print(deck)
}

func print(deck Deck) {
	for _, c := range deck.cards {
		fmt.Printf("%s%s\n", SINGS[c.value], SINGS[c.suit])
	}
}
