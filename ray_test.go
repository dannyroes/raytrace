package main

import "testing"

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

func TestSphereIntersect(t *testing.T) {
	cases := []struct {
		s        SphereType
		r        RayType
		expected []float64
	}{
		{
			s:        Sphere(1),
			r:        Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			expected: []float64{4, 6},
		},
		{
			s:        Sphere(2),
			r:        Ray(Point(0, 1, -5), Vector(0, 0, 1)),
			expected: []float64{5, 5},
		},
		{
			s:        Sphere(3),
			r:        Ray(Point(0, 2, -5), Vector(0, 0, 1)),
			expected: []float64{},
		},
		{
			s:        Sphere(4),
			r:        Ray(Point(0, 0, 0), Vector(0, 0, 1)),
			expected: []float64{-1, 1},
		},
		{
			s:        Sphere(5),
			r:        Ray(Point(0, 0, 5), Vector(0, 0, 1)),
			expected: []float64{-6, -4},
		},
	}

	for _, tc := range cases {
		result := tc.s.Intersects(tc.r)

		if len(result) != len(tc.expected) {
			t.Errorf("Result length mismatch expected %d, received %d", len(tc.expected), len(result))
		}

		for i := range result {
			if tc.expected[i] != result[i].T {
				t.Errorf("Result index %d mismatch expected %d, received %d", i, len(tc.expected), len(result))
			}

			if result[i].Object.Id != tc.s.Id {
				t.Errorf("Result index %d id mismatch expected %d, received %d", i, tc.s.Id, result[i].Object.Id)
			}
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
