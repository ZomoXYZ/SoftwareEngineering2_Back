package gameplay

import "edu/letu/wan/structs"

type Hand int

const (
	SingleFree Hand = iota

	Pair
	PairInverted
	PairFree

	TripleFree
	BigPair // double triangle + double circle / doubletrianglecircle + doubletriangleinverted / doublecircletriangle + doublecircleinverted

	WanMo

	NoHand
)

func (h Hand) Points() int {
	switch h {
	case SingleFree:
	case Pair:
		return 1
	case PairInverted:
	case PairFree:
		return 2
	case TripleFree:
	case BigPair:
		return 3
	case WanMo:
		return 4
	}
	return 0
}

// TODO count points and confirm cards are valid
func calculateHand(hand []structs.Card) Hand {
	if len(hand) == 1 && hand[0].IsFree() {
		return SingleFree
	}
	if len(hand) == 2 {
		
	}
	return NoHand
}