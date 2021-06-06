package main

import "math"

var Epsilon = 0.00001

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

func (t Tuple) Mul(m float64) Tuple {
	t.X *= m
	t.Y *= m
	t.Z *= m
	t.W *= m

	return t
}

func (t Tuple) Div(m float64) Tuple {
	t.X /= m
	t.Y /= m
	t.Z /= m
	t.W /= m

	return t
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2))
}

func (t Tuple) Normalize() Tuple {
	m := t.Magnitude()

	return t.Div(m)
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

func Dot(a, b Tuple) float64 {
	return a.X*b.X +
		a.Y*b.Y +
		a.Z*b.Z +
		a.W*b.W
}

func Cross(a, b Tuple) Tuple {
	return Vector(
		a.Y*b.Z-a.Z*b.Y,
		a.Z*b.X-a.X*b.Z,
		a.X*b.Y-a.Y*b.X,
	)
}

func (v Tuple) Reflect(n Tuple) Tuple {
	return v.Sub(n.Mul(2).Mul(Dot(v, n)))
}
