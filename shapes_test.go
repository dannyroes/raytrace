package main

import "testing"

func TestSphereTransform(t *testing.T) {
	cases := []struct {
		s        SphereType
		expected Matrix
	}{
		{
			s:        Sphere(1),
			expected: IdentityMatrix(),
		},
		{
			s:        Sphere(1).SetTransform(Translation(2, 3, 4)),
			expected: Translation(2, 3, 4),
		},
	}

	for _, tc := range cases {
		if !tc.expected.Equals(tc.s.Transform) {
			t.Errorf("transform mismatch expected %+v received %+v", tc.expected, tc.s.Transform)
		}
	}
}

func TestSphereIntersection(t *testing.T) {
	cases := []struct {
		s        SphereType
		r        RayType
		expected IntersectionList
	}{
		{
			s: Sphere(1).SetTransform(Scaling(2, 2, 2)),
			r: Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			expected: Intersections(
				Intersection(3, Sphere(1)),
				Intersection(7, Sphere(1)),
			),
		},
		{
			s:        Sphere(1).SetTransform(Translation(5, 0, 0)),
			r:        Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			expected: Intersections(),
		},
	}

	for _, tc := range cases {
		result := tc.s.Intersects(tc.r)

		if len(tc.expected) != len(result) {
			t.Errorf("intersection length mismatch expected %d received %d", len(tc.expected), len(result))
		}

		for i := range tc.expected {
			if !FloatEqual(tc.expected[i].T, result[i].T) {
				t.Errorf("intersection mismatch expected %f received %f", tc.expected[i].T, result[i].T)
			}
		}
	}
}
