package material

import (
	"math"

	"github.com/dannyroes/raytrace/data"
)

var (
	Black = Colour(0, 0, 0)
	White = Colour(1, 1, 1)
)

type ColourTuple struct {
	data.Tuple
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

func (c ColourTuple) Sub(b ColourTuple) ColourTuple {
	return ColourTuple{c.Tuple.Sub(b.Tuple)}
}

func (c ColourTuple) Mul(b float64) ColourTuple {
	return ColourTuple{c.Tuple.Mul(b)}
}

func (c ColourTuple) Div(b float64) ColourTuple {
	return ColourTuple{c.Tuple.Div(b)}
}

func MultiplyColours(a, b ColourTuple) ColourTuple {
	return Colour(
		a.Red()*b.Red(),
		a.Green()*b.Green(),
		a.Blue()*b.Blue(),
	)
}

func ColourEqual(a, b ColourTuple) bool {
	return data.TupleEqual(a.Tuple, b.Tuple)
}

func Colour(r, g, b float64) ColourTuple {
	return ColourTuple{data.Tuple{X: r, Y: g, Z: b}}
}

func (c ColourTuple) RGBA() (r, g, b, a uint32) {
	r = getPreciseColour(GetCappedColour(c.Red(), 255))
	g = getPreciseColour(GetCappedColour(c.Green(), 255))
	b = getPreciseColour(GetCappedColour(c.Blue(), 255))
	a = getPreciseColour(255)

	return
}

func GetCappedColour(ratio float64, cap int) int {
	colour := int(math.Ceil(ratio * float64(cap)))

	if colour < 0 {
		return 0
	}

	if colour > cap {
		return cap
	}

	return colour
}

func getPreciseColour(c int) uint32 {
	f := float64(c) / 255.0 * 65535.0
	return uint32(f)
}
