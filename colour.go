package main

type ColourTuple struct {
	Tuple
}

func (c ColourTuple) Red() float64 {
	return c.Tuple.X
}

func (c ColourTuple) Green() float64 {
	return c.Tuple.Y
}

func (c ColourTuple) Blue() float64 {
	return c.Tuple.Z
}

func (c ColourTuple) Add(b ColourTuple) ColourTuple {
	return ColourTuple{c.Tuple.Add(b.Tuple)}
}

func (c ColourTuple) Mul(b float64) ColourTuple {
	return ColourTuple{c.Tuple.Mul(b)}
}

func MultiplyColours(a, b ColourTuple) ColourTuple {
	return Colour(
		a.Red()*b.Red(),
		a.Green()*b.Green(),
		a.Blue()*b.Blue(),
	)
}

func ColourEqual(a, b ColourTuple) bool {
	return TupleEqual(a.Tuple, b.Tuple)
}

func Colour(r, g, b float64) ColourTuple {
	return ColourTuple{Tuple{X: r, Y: g, Z: b}}
}
