package main

import "testing"

func TestCubeIntersect(t *testing.T) {
	c := Cube()
	testCases := []struct {
		o        Tuple
		d        Tuple
		expected []float64
	}{
		{
			o:        Point(5, 0.5, 0),
			d:        Vector(-1, 0, 0),
			expected: []float64{4, 6},
		},
		{
			o:        Point(-5, 0.5, 0),
			d:        Vector(1, 0, 0),
			expected: []float64{4, 6},
		},
		{
			o:        Point(0.5, 5, 0),
			d:        Vector(0, -1, 0),
			expected: []float64{4, 6},
		},
		{
			o:        Point(0.5, -5, 0),
			d:        Vector(0, 1, 0),
			expected: []float64{4, 6},
		},
		{
			o:        Point(0.5, 0, 5),
			d:        Vector(0, 0, -1),
			expected: []float64{4, 6},
		},
		{
			o:        Point(0.5, 0, -5),
			d:        Vector(0, 0, 1),
			expected: []float64{4, 6},
		},
		{
			o:        Point(0, 0.5, 0),
			d:        Vector(0, 0, 1),
			expected: []float64{-1, 1},
		},
		{
			o:        Point(-2, 0, 0),
			d:        Vector(0.2673, 0.5345, 0.8018),
			expected: []float64{},
		},
		{
			o:        Point(0, -2, 0),
			d:        Vector(0.8018, 0.2673, 0.5345),
			expected: []float64{},
		},
		{
			o:        Point(0, 0, -2),
			d:        Vector(0.5345, 0.8018, 0.2673),
			expected: []float64{},
		},
		{
			o:        Point(2, 0, 2),
			d:        Vector(0, 0, -1),
			expected: []float64{},
		},
		{
			o:        Point(0, 2, 2),
			d:        Vector(0, -1, 0),
			expected: []float64{},
		},
		{
			o:        Point(2, 2, 0),
			d:        Vector(-1, 0, 0),
			expected: []float64{},
		},
	}

	for _, tc := range testCases {
		xs := c.LocalIntersect(Ray(tc.o, tc.d))

		if len(xs) != len(tc.expected) {
			t.Errorf("Expected %d intersects received: %d", len(tc.expected), len(xs))
			continue
		}

		if len(tc.expected) == 2 {
			if !FloatEqual(tc.expected[0], xs[0].T) {
				t.Errorf("First intersect mismatch expected: %f received: %f", tc.expected[0], xs[0].T)
			}

			if !FloatEqual(tc.expected[1], xs[1].T) {
				t.Errorf("Second intersect mismatch expected: %f received: %f", tc.expected[1], xs[1].T)
			}
		}
	}
}

func TestCubeNormal(t *testing.T) {
	c := Cube()
	testCases := []struct {
		p        Tuple
		expected Tuple
	}{
		{
			p:        Point(1, 0.5, -0.8),
			expected: Vector(1, 0, 0),
		},
		{
			p:        Point(-1, -0.2, 0.9),
			expected: Vector(-1, 0, 0),
		},
		{
			p:        Point(-0.4, 1, -0.1),
			expected: Vector(0, 1, 0),
		},
		{
			p:        Point(0.3, -1, -0.7),
			expected: Vector(0, -1, 0),
		},
		{
			p:        Point(-0.6, 0.3, 1),
			expected: Vector(0, 0, 1),
		},
		{
			p:        Point(0.4, 0.4, -1),
			expected: Vector(0, 0, -1),
		},
		{
			p:        Point(1, 1, 1),
			expected: Vector(1, 0, 0),
		},
		{
			p:        Point(-1, -1, -1),
			expected: Vector(-1, 0, 0),
		},
	}

	for _, tc := range testCases {
		result := c.LocalNormalAt(tc.p)
		if !TupleEqual(result, tc.expected) {
			t.Errorf("Normal mismatch expected %v received %v", tc.expected, result)
		}
	}
}
