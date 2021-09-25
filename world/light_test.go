package world

import (
	"math"
	"testing"

	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
	"github.com/dannyroes/raytrace/shape"
)

func TestLight(t *testing.T) {
	p := data.Point(0, 0, 0)
	i := material.Colour(1, 1, 1)

	l := PointLight(p, i)

	if !data.TupleEqual(p, l.Position) {
		t.Errorf("bad position expected: %v received %v", p, l.Position)
	}

	if !data.TupleEqual(i.Tuple, l.Intensity.Tuple) {
		t.Errorf("bad intensity expected: %v received %v", p, l.Position)
	}
}

func TestLighting(t *testing.T) {
	pos := data.Point(0, 0, 0)
	m := material.Material()

	cases := []struct {
		eyeV     data.Tuple
		normalV  data.Tuple
		light    Light
		inShadow bool
		expected material.ColourTuple
	}{
		{
			eyeV:     data.Vector(0, 0, -1),
			normalV:  data.Vector(0, 0, -1),
			light:    PointLight(data.Point(0, 0, -10), material.Colour(1, 1, 1)),
			expected: material.Colour(1.9, 1.9, 1.9),
		},
		{
			eyeV:     data.Vector(0, math.Sqrt(2)/2, -1*math.Sqrt(2)/2),
			normalV:  data.Vector(0, 0, -1),
			light:    PointLight(data.Point(0, 0, -10), material.Colour(1, 1, 1)),
			expected: material.Colour(1.0, 1.0, 1.0),
		},
		{
			eyeV:     data.Vector(0, 0, -1),
			normalV:  data.Vector(0, 0, -1),
			light:    PointLight(data.Point(0, 10, -10), material.Colour(1, 1, 1)),
			expected: material.Colour(0.7364, 0.7364, 0.7364),
		},
		{
			eyeV:     data.Vector(0, 0, -1),
			normalV:  data.Vector(0, 0, -1),
			light:    PointLight(data.Point(0, 0, 10), material.Colour(1, 1, 1)),
			expected: material.Colour(0.1, 0.1, 0.1),
		},
		{
			eyeV:     data.Vector(0, 0, -1),
			normalV:  data.Vector(0, 0, -1),
			light:    PointLight(data.Point(0, 0, -10), material.Colour(1, 1, 1)),
			inShadow: true,
			expected: material.Colour(0.1, 0.1, 0.1),
		},
	}

	for _, tc := range cases {
		result := Lighting(m, shape.Sphere(), tc.light, pos, tc.eyeV, tc.normalV, tc.inShadow)

		if !material.ColourEqual(tc.expected, result) {
			t.Errorf("colour mismatch expected: %v received %v", tc.expected, result)
		}
	}
}
