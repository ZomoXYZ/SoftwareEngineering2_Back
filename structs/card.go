package structs

type Card int

const (
	Circle1 Card = iota
	Circle2
	CircleInverted1
	CircleInverted2

	CircleTriangle1
	CircleTriangle2

	CircleFree

	Triangle1
	Triangle2
	TriangleInverted1
	TriangleInverted2

	TriangleCircle1
	TriangleCircle2

	TriangleFree

	Free
)