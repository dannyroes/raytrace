package data

import (
	"testing"
)

func TestRayPosition(t *testing.T) {
	ray := Ray(Point(2, 3, 4), Vector(1, 0, 0))

	cases := []struct {
		t        float64
		expected Tuple
	}{
		{t: 0, expected: Point(2, 3, 4)},
		{t: 1, expected: Point(3, 3, 4)},
		{t: -1, expected: Point(1, 3, 4)},
		{t: 2.5, expected: Point(4.5, 3, 4)},
	}

	for _, tc := range cases {
		p := ray.Position(tc.t)

		if !TupleEqual(tc.expected, p) {
			t.Errorf("expected: %+v, received: %+v", tc.expected, p)
		}
	}
}

func TestRayTransform(t *testing.T) {
	cases := []struct {
		r        RayType
		m        Matrix
		expected RayType
	}{
		{
			r:        Ray(Point(1, 2, 3), Vector(0, 1, 0)),
			m:        Translation(3, 4, 5),
			expected: Ray(Point(4, 6, 8), Vector(0, 1, 0)),
		},
		{
			r:        Ray(Point(1, 2, 3), Vector(0, 1, 0)),
			m:        Scaling(2, 3, 4),
			expected: Ray(Point(2, 6, 12), Vector(0, 3, 0)),
		},
	}

	for _, tc := range cases {
		result := tc.r.Transform(tc.m)

		if result != tc.expected {
			t.Errorf("expected: %+v, received: %+v", tc.expected, result)
		}
	}
}
