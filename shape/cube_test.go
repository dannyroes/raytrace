package shape

import (
	"testing"

	"github.com/dannyroes/raytrace/data"
)

func TestCubeIntersect(t *testing.T) {
	c := Cube()
	testCases := []struct {
		o        data.Tuple
		d        data.Tuple
		expected []float64
	}{
		{
			o:        data.Point(5, 0.5, 0),
			d:        data.Vector(-1, 0, 0),
			expected: []float64{4, 6},
		},
		{
			o:        data.Point(-5, 0.5, 0),
			d:        data.Vector(1, 0, 0),
			expected: []float64{4, 6},
		},
		{
			o:        data.Point(0.5, 5, 0),
			d:        data.Vector(0, -1, 0),
			expected: []float64{4, 6},
		},
		{
			o:        data.Point(0.5, -5, 0),
			d:        data.Vector(0, 1, 0),
			expected: []float64{4, 6},
		},
		{
			o:        data.Point(0.5, 0, 5),
			d:        data.Vector(0, 0, -1),
			expected: []float64{4, 6},
		},
		{
			o:        data.Point(0.5, 0, -5),
			d:        data.Vector(0, 0, 1),
			expected: []float64{4, 6},
		},
		{
			o:        data.Point(0, 0.5, 0),
			d:        data.Vector(0, 0, 1),
			expected: []float64{-1, 1},
		},
		{
			o:        data.Point(-2, 0, 0),
			d:        data.Vector(0.2673, 0.5345, 0.8018),
			expected: []float64{},
		},
		{
			o:        data.Point(0, -2, 0),
			d:        data.Vector(0.8018, 0.2673, 0.5345),
			expected: []float64{},
		},
		{
			o:        data.Point(0, 0, -2),
			d:        data.Vector(0.5345, 0.8018, 0.2673),
			expected: []float64{},
		},
		{
			o:        data.Point(2, 0, 2),
			d:        data.Vector(0, 0, -1),
			expected: []float64{},
		},
		{
			o:        data.Point(0, 2, 2),
			d:        data.Vector(0, -1, 0),
			expected: []float64{},
		},
		{
			o:        data.Point(2, 2, 0),
			d:        data.Vector(-1, 0, 0),
			expected: []float64{},
		},
	}

	for _, tc := range testCases {
		xs := c.LocalIntersect(data.Ray(tc.o, tc.d))

		if len(xs) != len(tc.expected) {
			t.Errorf("Expected %d intersects received: %d", len(tc.expected), len(xs))
			continue
		}

		if len(tc.expected) == 2 {
			if !data.FloatEqual(tc.expected[0], xs[0].T) {
				t.Errorf("First intersect mismatch expected: %f received: %f", tc.expected[0], xs[0].T)
			}

			if !data.FloatEqual(tc.expected[1], xs[1].T) {
				t.Errorf("Second intersect mismatch expected: %f received: %f", tc.expected[1], xs[1].T)
			}
		}
	}
}

func TestCubeNormal(t *testing.T) {
	c := Cube()
	testCases := []struct {
		p        data.Tuple
		expected data.Tuple
	}{
		{
			p:        data.Point(1, 0.5, -0.8),
			expected: data.Vector(1, 0, 0),
		},
		{
			p:        data.Point(-1, -0.2, 0.9),
			expected: data.Vector(-1, 0, 0),
		},
		{
			p:        data.Point(-0.4, 1, -0.1),
			expected: data.Vector(0, 1, 0),
		},
		{
			p:        data.Point(0.3, -1, -0.7),
			expected: data.Vector(0, -1, 0),
		},
		{
			p:        data.Point(-0.6, 0.3, 1),
			expected: data.Vector(0, 0, 1),
		},
		{
			p:        data.Point(0.4, 0.4, -1),
			expected: data.Vector(0, 0, -1),
		},
		{
			p:        data.Point(1, 1, 1),
			expected: data.Vector(1, 0, 0),
		},
		{
			p:        data.Point(-1, -1, -1),
			expected: data.Vector(-1, 0, 0),
		},
	}

	for _, tc := range testCases {
		result := c.LocalNormalAt(tc.p)
		if !data.TupleEqual(result, tc.expected) {
			t.Errorf("Normal mismatch expected %v received %v", tc.expected, result)
		}
	}
}
