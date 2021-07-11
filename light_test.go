package main

import (
	"math"
	"testing"
)

func TestLight(t *testing.T) {
	p := Point(0, 0, 0)
	i := Colour(1, 1, 1)

	l := PointLight(p, i)

	if !TupleEqual(p, l.Position) {
		t.Errorf("bad position expected: %v received %v", p, l.Position)
	}

	if !TupleEqual(i.Tuple, l.Intensity.Tuple) {
		t.Errorf("bad intensity expected: %v received %v", p, l.Position)
	}
}

func TestLighting(t *testing.T) {
	pos := Point(0, 0, 0)
	m := Material()

	cases := []struct {
		eyeV     Tuple
		normalV  Tuple
		light    Light
		inShadow bool
		expected ColourTuple
	}{
		{
			eyeV:     Vector(0, 0, -1),
			normalV:  Vector(0, 0, -1),
			light:    PointLight(Point(0, 0, -10), Colour(1, 1, 1)),
			expected: Colour(1.9, 1.9, 1.9),
		},
		{
			eyeV:     Vector(0, math.Sqrt(2)/2, -1*math.Sqrt(2)/2),
			normalV:  Vector(0, 0, -1),
			light:    PointLight(Point(0, 0, -10), Colour(1, 1, 1)),
			expected: Colour(1.0, 1.0, 1.0),
		},
		{
			eyeV:     Vector(0, 0, -1),
			normalV:  Vector(0, 0, -1),
			light:    PointLight(Point(0, 10, -10), Colour(1, 1, 1)),
			expected: Colour(0.7364, 0.7364, 0.7364),
		},
		{
			eyeV:     Vector(0, 0, -1),
			normalV:  Vector(0, 0, -1),
			light:    PointLight(Point(0, 0, 10), Colour(1, 1, 1)),
			expected: Colour(0.1, 0.1, 0.1),
		},
		{
			eyeV:     Vector(0, 0, -1),
			normalV:  Vector(0, 0, -1),
			light:    PointLight(Point(0, 0, -10), Colour(1, 1, 1)),
			inShadow: true,
			expected: Colour(0.1, 0.1, 0.1),
		},
	}

	for _, tc := range cases {
		result := Lighting(m, tc.light, pos, tc.eyeV, tc.normalV, tc.inShadow)

		if !ColourEqual(tc.expected, result) {
			t.Errorf("colour mismatch expected: %v received %v", tc.expected, result)
		}
	}

}
