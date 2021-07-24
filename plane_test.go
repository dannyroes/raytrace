package main

import "testing"

func TestPlaneNormal(t *testing.T) {
	p := Plane()

	points := []Tuple{
		Point(0, 0, 0),
		Point(10, 0, -10),
		Point(-5, 0, 150),
	}

	for _, point := range points {
		normal := p.LocalNormalAt(point)

		if !TupleEqual(Vector(0, 1, 0), normal) {
			t.Errorf("Normal mismatch at %v expected %v received %v", point, Vector(0, 1, 0), normal)
		}
	}
}

func TestPlaneIntersect(t *testing.T) {
	p := Plane()

	cases := []struct {
		ray      RayType
		expected IntersectionList
	}{
		{
			ray:      Ray(Point(0, 10, 0), Vector(0, 0, 1)),
			expected: IntersectionList{},
		},
		{
			ray:      Ray(Point(0, 0, 0), Vector(0, 0, 1)),
			expected: IntersectionList{},
		},
		{
			ray:      Ray(Point(0, 1, 0), Vector(0, -1, 0)),
			expected: IntersectionList{Intersection(1, p)},
		},
		{
			ray:      Ray(Point(0, -1, 0), Vector(0, 1, 0)),
			expected: IntersectionList{Intersection(1, p)},
		},
	}

	for _, tc := range cases {
		intersects := p.LocalIntersect(tc.ray)

		if len(tc.expected) != len(intersects) {
			t.Errorf("Wrong number of intersects expected %v received %v", len(tc.expected), len(intersects))
		} else {
			for i, intersect := range intersects {
				if !FloatEqual(intersect.T, tc.expected[i].T) || intersect.Object != tc.expected[i].Object {
					t.Errorf("Intersect mismatch expected %v received %v", tc.expected[i], intersects[i])
				}
			}
		}
	}
}
