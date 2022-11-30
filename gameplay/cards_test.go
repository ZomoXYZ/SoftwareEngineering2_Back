package gameplay

import (
	"edu/letu/wan/structs"
	"testing"
)

/*
SingleFree

Pair
PairInverted
PairFree

TripleFree
DoubleShapePair // double triangle + double circle
BigPair // doubletrianglecircle + doubletriangleinverted / doublecircletriangle + doublecircleinverted

QuadFree
WanMo

NoHand
*/

func TestSingleFree(t *testing.T) {
	
	hand := calculateHand([]structs.Card{
		structs.Free,
	})

	if hand != SingleFree {
		t.Errorf("Free Card should be SingleFree, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.CircleFree,
	})

	if hand != SingleFree {
		t.Errorf("Circle Free Card should be SingleFree, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.TriangleFree,
	})

	if hand != SingleFree {
		t.Errorf("Triangle Free Card should be SingleFree, GOT: %d", hand)
	}

}

func TestPair(t *testing.T) {
	
	hand := calculateHand([]structs.Card{
		structs.Circle1,
		structs.Circle1,
	})

	if hand != Pair {
		t.Errorf("Circle1 should be Pair, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.Circle2,
		structs.Circle2,
	})

	if hand != Pair {
		t.Errorf("Circle2 should be Pair, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.Triangle1,
		structs.Triangle1,
	})

	if hand != Pair {
		t.Errorf("Triangle1 should be Pair, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.Triangle2,
		structs.Triangle2,
	})

	if hand != Pair {
		t.Errorf("Triangle2 should be Pair, GOT: %d", hand)
	}

}

func TestPairInverted(t *testing.T) {
	
	hand := calculateHand([]structs.Card{
		structs.Circle1,
		structs.CircleInverted1,
	})

	if hand != PairInverted {
		t.Errorf("Circle1 should be PairInverted, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.Circle2,
		structs.CircleInverted2,
	})

	if hand != PairInverted {
		t.Errorf("Circle2 should be PairInverted, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.Triangle1,
		structs.TriangleInverted1,
	})

	if hand != PairInverted {
		t.Errorf("Triangle1 should be PairInverted, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.Triangle2,
		structs.TriangleInverted2,
	})

	if hand != PairInverted {
		t.Errorf("Triangle2 should be PairInverted, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.Circle1,
		structs.CircleFree,
	})

	if hand != PairInverted {
		t.Errorf("Circle1+CircleFree should be PairInverted, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.Circle2,
		structs.Free,
	})

	if hand != PairInverted {
		t.Errorf("Circle2+Free should be PairInverted, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.Triangle1,
		structs.TriangleFree,
	})

	if hand != PairInverted {
		t.Errorf("Triangle1+TriangleFree should be PairInverted, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.Triangle2,
		structs.Free,
	})

	if hand != PairInverted {
		t.Errorf("Triangle2+Free should be PairInverted, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.CircleInverted1,
		structs.CircleFree,
	})

	if hand != PairInverted {
		t.Errorf("CircleInverted1+CircleFree should be PairInverted, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.CircleInverted2,
		structs.Free,
	})

	if hand != PairInverted {
		t.Errorf("CircleInverted2+Free should be PairInverted, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.TriangleInverted1,
		structs.TriangleFree,
	})

	if hand != PairInverted {
		t.Errorf("TriangleInverted1+TriangleFree should be PairInverted, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.TriangleInverted2,
		structs.Free,
	})

	if hand != PairInverted {
		t.Errorf("TriangleInverted2+Free should be PairInverted, GOT: %d", hand)
	}

}

func TestDoubleShapePair(t *testing.T) {
	
	hand := calculateHand([]structs.Card{
		structs.Circle2,
		structs.Triangle2,
	})

	if hand != DoubleShapePair {
		t.Errorf("Circle2+Triangle2 should be DoubleShapePair, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.Triangle2,
		structs.Circle2,
	})

	if hand != DoubleShapePair {
		t.Errorf("Triangle2+Circle2 should be DoubleShapePair, GOT: %d", hand)
	}

}

func TestBigPair(t *testing.T) {
	
	hand := calculateHand([]structs.Card{
		structs.TriangleCircle2,
		structs.TriangleInverted2,
	})

	if hand != BigPair {
		t.Errorf("TriangleCircle2+TriangleInverted2 should be BigPair, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.TriangleInverted2,
		structs.TriangleCircle2,
	})

	if hand != BigPair {
		t.Errorf("TriangleInverted2+TriangleCircle2 should be BigPair, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.CircleTriangle2,
		structs.CircleInverted2,
	})

	if hand != BigPair {
		t.Errorf("CircleTriangle2+CircleInverted2 should be BigPair, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.CircleInverted2,
		structs.CircleTriangle2,
	})

	if hand != BigPair {
		t.Errorf("CircleInverted2+CircleTriangle2 should be BigPair, GOT: %d", hand)
	}

}

func TestFree(t *testing.T) {
	
	hand := calculateHand([]structs.Card{
		structs.Free,
	})

	if hand != SingleFree {
		t.Errorf("SingleFree, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.CircleFree,
	})

	if hand != SingleFree {
		t.Errorf("SingleFree, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.TriangleFree,
	})

	if hand != SingleFree {
		t.Errorf("SingleFree, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.Free,
		structs.CircleFree,
	})

	if hand != DoubleFree {
		t.Errorf("DoubleFree, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.Free,
		structs.TriangleFree,
		structs.Free,
	})

	if hand != TripleFree {
		t.Errorf("TripleFree, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.Free,
		structs.TriangleFree,
		structs.CircleFree,
		structs.Free,
	})

	if hand != QuadFree {
		t.Errorf("QuadFree, GOT: %d", hand)
	}

}

func TestNoHand(t *testing.T) {
	
	hand := calculateHand([]structs.Card{
		structs.Circle1,
	})

	if hand != NoHand {
		t.Errorf("NoHand, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.Circle1,
		structs.Circle2,
	})

	if hand != NoHand {
		t.Errorf("NoHand, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.TriangleFree,
		structs.Circle1,
	})

	if hand != NoHand {
		t.Errorf("NoHand, GOT: %d", hand)
	}
	
	hand = calculateHand([]structs.Card{
		structs.TriangleFree,
		structs.Circle1,
	})

	if hand != NoHand {
		t.Errorf("NoHand, GOT: %d", hand)
	}

}

// func TestWanMo(t *testing.T) {
	
// 	hand := calculateHand([]structs.Card{
// 		structs.TriangleCircle2,
// 		structs.TriangleInverted2,
// 		structs.Circle1,
// 		structs.Circle1,
// 	})

// 	if hand != WanMo {
// 		t.Errorf("WanMo 1, GOT: %d", hand)
// 	}
	
// 	hand = calculateHand([]structs.Card{
// 		structs.CircleFree,
// 		structs.TriangleInverted2,
// 		structs.Free,
// 		structs.Circle1,
// 	})

// 	if hand != WanMo {
// 		t.Errorf("WanMo 1 + free, GOT: %d", hand)
// 	}
	
// 	hand = calculateHand([]structs.Card{
// 		structs.Circle1,
// 		structs.TriangleInverted2,
// 		structs.Circle1,
// 		structs.TriangleCircle2,
// 	})

// 	if hand != WanMo {
// 		t.Errorf("WanMo 1 shuffle, GOT: %d", hand)
// 	}
	
// 	hand = calculateHand([]structs.Card{
// 		structs.Free,
// 		structs.TriangleInverted2,
// 		structs.Circle1,
// 		structs.CircleFree,
// 	})

// 	if hand != WanMo {
// 		t.Errorf("WanMo 1 shuffle + free, GOT: %d", hand)
// 	}
	
// 	hand = calculateHand([]structs.Card{
// 		structs.CircleTriangle2,
// 		structs.CircleInverted2,
// 		structs.TriangleInverted1,
// 		structs.TriangleInverted1,
// 	})

// 	if hand != BigPair {
// 		t.Errorf("WanMo 2, GOT: %d", hand)
// 	}
	
// 	hand = calculateHand([]structs.Card{
// 		structs.TriangleInverted1,
// 		structs.CircleTriangle2,
// 		structs.CircleInverted2,
// 		structs.TriangleInverted1,
// 	})

// 	if hand != BigPair {
// 		t.Errorf("WanMo 2 shuffle, GOT: %d", hand)
// 	}

// }
