package material

import (
	"testing"

	"github.com/dannyroes/raytrace/data"
)

func TestColourBasics(t *testing.T) {
	cases := []struct {
		c ColourTuple
		r float64
		g float64
		b float64
	}{
		{
			c: Colour(
				-0.5,
				0.4,
				1.7,
			),
			r: -0.5,
			g: 0.4,
			b: 1.7,
		},
	}

	for _, c := range cases {
		if !data.FloatEqual(c.c.Red(), c.r) {
			t.Errorf("Red value incorrect. expected: %f; actual: %f", c.r, c.c.Red())
		}

		if !data.FloatEqual(c.c.Green(), c.g) {
			t.Errorf("Green value incorrect. expected: %f; actual: %f", c.g, c.c.Green())
		}

		if !data.FloatEqual(c.c.Blue(), c.b) {
			t.Errorf("Blue value incorrect. expected: %f; actual: %f", c.b, c.c.Blue())
		}
	}
}

func TestAddColour(t *testing.T) {
	a := Colour(0.9, 0.6, 0.75)
	b := Colour(0.7, 0.1, 0.25)

	c := a.Add(b)
	expected := Colour(1.6, 0.7, 1.0)

	if !ColourEqual(c, expected) {
		t.Errorf("Expected %+v, received %+v", expected, c)
	}
}

func TestSubColour(t *testing.T) {
	a := Colour(0.9, 0.6, 0.75)
	b := Colour(0.7, 0.1, 0.25)

	c := a.Add(b)
	expected := Colour(1.6, 0.7, 1.0)

	if !ColourEqual(c, expected) {
		t.Errorf("Expected %+v, received %+v", expected, c)
	}
}

func TestMultiplyColour(t *testing.T) {
	a := Colour(0.2, 0.3, 0.4)

	result := a.Mul(2)
	expected := Colour(0.4, 0.6, 0.8)

	if !ColourEqual(result, expected) {
		t.Errorf("Expected %+v, received %+v", expected, result)
	}
}

func TestMultiplyColours(t *testing.T) {
	a := Colour(1, 0.2, 0.4)
	b := Colour(0.9, 1, 0.1)

	c := MultiplyColours(a, b)
	expected := Colour(0.9, 0.2, 0.04)

	if !ColourEqual(c, expected) {
		t.Errorf("Expected %+v, received %+v", expected, c)
	}
}
