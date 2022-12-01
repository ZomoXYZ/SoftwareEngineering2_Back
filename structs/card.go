package structs

import (
	"math/rand"
	"time"
)

type Card int

const (
	Circle1 Card = iota
	Circle2
	CircleInverted1
	CircleInverted2

	CircleTriangle1
	CircleTriangle2

	Triangle1
	Triangle2
	TriangleInverted1
	TriangleInverted2

	TriangleCircle1
	TriangleCircle2

	Free
	CircleFree
	TriangleFree
)

type CardShape int
const (
	Circle CardShape = iota
	Triangle
	Combination
	FreeShape
)

func (c Card) Shape() CardShape {
	switch c {
	case Circle1, Circle2, CircleFree, CircleInverted1, CircleInverted2:
		return Circle
	case Triangle1, Triangle2, TriangleFree, TriangleInverted1, TriangleInverted2:
		return Triangle
	case CircleTriangle1, CircleTriangle2, TriangleCircle1, TriangleCircle2:
		return Combination
	}
	return FreeShape
}

type CardInverted int
const (
	Regular CardInverted = iota
	Inverted
	FreeInvert
)

func (c Card) Inverted() CardInverted {
	switch c {
	case Circle1, Circle2, Triangle1, Triangle2, CircleTriangle1, CircleTriangle2:
		return Regular
	case CircleInverted1, CircleInverted2, TriangleInverted1, TriangleInverted2, TriangleCircle1, TriangleCircle2:
		return Inverted
	}
	return FreeInvert
}

type CardCount int
const (
	Single CardCount = iota
	Double
	FreeCount
)

func (c Card) Count() CardCount {
	switch c {
	case Circle1, Triangle1, CircleInverted1, TriangleInverted1, CircleTriangle1, TriangleCircle1:
		return Single
	case Circle2, Triangle2, CircleInverted2, TriangleInverted2, CircleTriangle2, TriangleCircle2:
		return Double
	}
	return FreeCount
}

func (c Card) IsFree() bool {
	return c == Free || c == CircleFree || c == TriangleFree
}

func (c Card) MatchShape(d Card) bool {
	return c.Shape() == d.Shape() || c.Shape() == FreeShape || d.Shape() == FreeShape
}

func (c Card) MatchInvert(d Card) bool {
	return c.Inverted() == d.Inverted() || c.Inverted() == FreeInvert || d.Inverted() == FreeInvert
}

func (c Card) MatchCount(d Card) bool {
	return c.Count() == d.Count() || c.Count() == FreeCount || d.Count() == FreeCount
}

func (c Card) MatchShapeCount(d Card) bool {
	return c.MatchShape(d) && c.MatchCount(d)
}

func (c Card) MatchAll(d Card) bool {
	return c.MatchShape(d) && c.MatchCount(d) && c.MatchInvert(d)
}

func (card Card) isMoreSpecific(other Card) bool {
	// card isn't free
	if !card.IsFree() {
		return true
	}
	// other isn't free
	if !other.IsFree() {
		return false
	}
	// card isn't any free
	if card.Shape() != FreeShape {
		return true
	}
	// other isn't any free
	if other.Shape() != FreeShape {
		return false
	}
	// both are any free
	return true
}

func removeCard(hand []Card, card Card) (bool, []Card) {
	var bestCard Card = -1
	var bestIndex = -1
	for i, c := range hand {
		if card.MatchAll(c) {
			if bestCard == -1 || c.isMoreSpecific(bestCard) {
				bestCard = c
				bestIndex = i
			}
		}
	}
	if bestIndex != -1 {
		return true, append(hand[:bestIndex], hand[bestIndex+1:]...)
	}
	return false, []Card{}
}

func CardsFollow(hand []Card, pattern ...Card) bool {
	if len(hand) != len(pattern) {
		return false
	}
	handCopy := make([]Card, len(hand))
	copy(handCopy, hand)
	for _, cardPattern := range pattern {
		removed, newHand := removeCard(handCopy, cardPattern)
		if !removed {
			return false
		}
		handCopy = newHand
	}
	return true
}

func RandomCard() Card {
	// every individual card has an equal chance of being selected, except all forms of free card are counted as one
    rand.Seed(time.Now().UnixNano())
	num := rand.Intn(13)

	// is free card
	if num < 12 {
		return Card(num)
	} else {
		num := rand.Intn(4)
		if num < 2 {
			return Free
		} else if num == 2 {
			return CircleFree
		} else {
			return TriangleFree
		}
	}
}
