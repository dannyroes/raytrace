package main

import "math"

var Epsilon = math.Nextafter(1.0, 2.0) - 1.0

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (t Tuple) IsPoint() bool {
	return FloatEqual(1.0, t.W)
}

func (t Tuple) IsVector() bool {
	return FloatEqual(0.0, t.W)
}

func (t Tuple) Add(b Tuple) Tuple {
	t.X += b.X
	t.Y += b.Y
	t.Z += b.Z
	t.W += b.W

	return t
}

func (t Tuple) Sub(b Tuple) Tuple {
	t.X -= b.X
	t.Y -= b.Y
	t.Z -= b.Z
	t.W -= b.W

	return t
}

func (t Tuple) Neg() Tuple {
	return Vector(0, 0, 0).Sub(t)
}

func FloatEqual(x, y float64) bool {
	return math.Abs(x-y) < Epsilon
}

func Point(x, y, z float64) Tuple {
	return Tuple{
		X: x,
		Y: y,
		Z: z,
		W: 1.0,
	}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{
		X: x,
		Y: y,
		Z: z,
		W: 0.0,
	}
}

func TupleEqual(a, b Tuple) bool {
	return FloatEqual(a.X, b.X) &&
		FloatEqual(a.Y, b.Y) &&
		FloatEqual(a.Z, b.Z) &&
		FloatEqual(a.W, b.W)
}
