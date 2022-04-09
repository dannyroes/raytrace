package shape

import (
	"testing"

	"github.com/dannyroes/raytrace/data"
)

func TestTriangle(t *testing.T) {
	p1 := data.Point(0, 1, 0)
	p2 := data.Point(-1, 0, 0)
	p3 := data.Point(1, 0, 0)

	e1 := data.Vector(-1, -1, 0)
	e2 := data.Vector(1, -1, 0)

	normal := data.Vector(0, 0, -1)

	tri := Triangle(p1, p2, p3)

	if !data.TupleEqual(p1, tri.p1) {
		t.Errorf("P1 mismatch expected %v received %v", p1, tri.p1)
	}

	if !data.TupleEqual(p2, tri.p2) {
		t.Errorf("P2 mismatch expected %v received %v", p2, tri.p2)
	}

	if !data.TupleEqual(p3, tri.p3) {
		t.Errorf("P3 mismatch expected %v received %v", p3, tri.p3)
	}

	if !data.TupleEqual(e1, tri.e1) {
		t.Errorf("E1 mismatch expected %v received %v", e1, tri.e1)
	}

	if !data.TupleEqual(e2, tri.e2) {
		t.Errorf("E2 mismatch expected %v received %v", e2, tri.e2)
	}

	if !data.TupleEqual(normal, tri.normal) {
		t.Errorf("Normal mismatch expected %v received %v", normal, tri.normal)
	}
}

func TestTriangleNormal(t *testing.T) {
	tri := Triangle(data.Point(0, 1, 0), data.Point(-1, 0, 0), data.Point(1, 0, 0))
	cases := []data.Tuple{
		data.Point(0, 0.5, 0),
		data.Point(-0.5, 0.75, 0),
		data.Point(0.5, 0.25, 0),
	}

	for _, p := range cases {
		n := tri.LocalNormalAt(p)

		if !data.TupleEqual(n, tri.normal) {
			t.Errorf("Normal mismatch for point %v expected %v received %v", p, tri.normal, n)
		}
	}
}

func TestTriangleIntersect(t *testing.T) {
	cases := []struct {
		t        *TriangleType
		r        data.RayType
		expected IntersectionList
	}{
		{
			t:        Triangle(data.Point(0, 1, 0), data.Point(-1, 0, 0), data.Point(1, 0, 0)),
			r:        data.Ray(data.Point(0, -1, -2), data.Vector(0, 1, 0)),
			expected: Intersections(),
		},
		{
			t:        Triangle(data.Point(0, 1, 0), data.Point(-1, 0, 0), data.Point(1, 0, 0)),
			r:        data.Ray(data.Point(1, 1, -2), data.Vector(0, 0, 1)),
			expected: Intersections(),
		},
		{
			t:        Triangle(data.Point(0, 1, 0), data.Point(-1, 0, 0), data.Point(1, 0, 0)),
			r:        data.Ray(data.Point(-1, 1, 2), data.Vector(0, 0, 1)),
			expected: Intersections(),
		},
		{
			t:        Triangle(data.Point(0, 1, 0), data.Point(-1, 0, 0), data.Point(1, 0, 0)),
			r:        data.Ray(data.Point(0, -1, -2), data.Vector(0, 0, 1)),
			expected: Intersections(),
		},
		{
			t:        Triangle(data.Point(0, 1, 0), data.Point(-1, 0, 0), data.Point(1, 0, 0)),
			r:        data.Ray(data.Point(0, 0.5, -2), data.Vector(0, 0, 1)),
			expected: Intersections(Intersection(2, &TriangleType{})),
		},
	}

	for _, tc := range cases {
		xs := tc.t.LocalIntersect(tc.r)

		if len(xs) != len(tc.expected) {
			t.Errorf("Intersections len mismatch expected %d received %d", len(tc.expected), len(xs))
		}

		for i, x := range xs {
			if x.T != tc.expected[i].T {
				t.Errorf("Intersection %d t mismatch expected %.2f received %.2f", i, tc.expected[i].T, x.T)
			}
		}
	}
}
