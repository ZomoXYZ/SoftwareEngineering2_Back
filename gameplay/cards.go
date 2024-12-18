package gameplay

import (
	"edu/letu/wan/structs"
)

type Hand int
const (
	SingleFree Hand = iota

	Pair
	PairInverted
	DoubleFree

	TripleFree
	DoubleShapePair // double triangle + double circle
	BigPair // doubletrianglecircle + doubletriangleinverted / doublecircletriangle + doublecircleinverted

	QuadFree
	WanMo_DoubleShapePair
	WanMo_BigPair

	NoHand
)

func (h Hand) Points() int {
	switch h {
	case SingleFree, Pair:
		return 1
	case PairInverted, DoubleFree:
		return 2
	case TripleFree, DoubleShapePair, BigPair:
		return 3
	case QuadFree, WanMo_DoubleShapePair, WanMo_BigPair:
		return 4
	}
	return 0
}

func calculateHand(hand []structs.Card, wanMoPair []structs.Card) Hand {
	if len(hand) == 1 && hand[0].IsFree() {
		return SingleFree
	}
	if len(hand) == 2 {
		if hand[0].IsFree() && hand[1].IsFree() {
			return DoubleFree
		}
		// cards are the same
		//     don't take free cards into account because they'll be be given PairInverted below (more points)
		if hand[0] == hand[1] {
			return Pair
		}
		// special pairs
		if structs.CardsFollow(hand, structs.Circle2, structs.Triangle2) {
			if len(wanMoPair) == 2 && wanMoPair[0].MatchAll(wanMoPair[1]) {
				return WanMo_DoubleShapePair
			}
			return DoubleShapePair
		}
		if structs.CardsFollow(hand, structs.CircleTriangle2, structs.CircleInverted2) || structs.CardsFollow(hand, structs.TriangleCircle2, structs.TriangleInverted2) {
			if len(wanMoPair) == 2 && wanMoPair[0].MatchAll(wanMoPair[1]) {
				return WanMo_BigPair
			}
			return BigPair
		}
		// cards match shape and count but are inverted
		if hand[0].MatchShapeCount(hand[1]) {
			return PairInverted
		}
	}
	if len(hand) == 3 {
		if hand[0].IsFree() && hand[1].IsFree() && hand[2].IsFree() {
			return TripleFree
		}
	}
	if len(hand) == 4 {
		if hand[0].IsFree() && hand[1].IsFree() && hand[2].IsFree() && hand[3].IsFree() {
			return QuadFree
		}
	}
	return NoHand
}