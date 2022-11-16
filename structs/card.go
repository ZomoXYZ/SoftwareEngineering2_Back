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

func RandomCard() Card {
	// every individual card has an equal chance of being selected, except all forms of free card are counted as one
    rand.Seed(time.Now().UnixNano())
	num := rand.Intn(13)

	// is free card
	if num < 12 {
		return Card(num)
	} else {
		num := rand.Intn(3)
		if num < 2 {
			return Free
		} else if num == 2 {
			return CircleFree
		} else {
			return TriangleFree
		}
	}
}
