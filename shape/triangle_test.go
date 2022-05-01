package shape

import (
	"fmt"
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

	n1 := data.Vector(0, 1, 0)
	n2 := data.Vector(-1, 0, 0)
	n3 := data.Vector(1, 0, 0)
	st := SmoothTriangle(p1, p2, p3, n1, n2, n3)

	if !data.TupleEqual(p1, st.p1) {
		t.Errorf("P1 mismatch expected %v received %v", p1, st.p1)
	}

	if !data.TupleEqual(p2, st.p2) {
		t.Errorf("P2 mismatch expected %v received %v", p2, st.p2)
	}

	if !data.TupleEqual(p3, st.p3) {
		t.Errorf("P3 mismatch expected %v received %v", p3, st.p3)
	}

	if !data.TupleEqual(n1, st.n1) {
		t.Errorf("N1 mismatch expected %v received %v", n1, st.n1)
	}

	if !data.TupleEqual(n2, st.n2) {
		t.Errorf("N2 mismatch expected %v received %v", n2, st.n2)
	}

	if !data.TupleEqual(n3, st.n3) {
		t.Errorf("N3 mismatch expected %v received %v", n3, st.n3)
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
		n := tri.LocalNormalAt(p, IntersectionType{})

		if !data.TupleEqual(n, tri.normal) {
			t.Errorf("Normal mismatch for point %v expected %v received %v", p, tri.normal, n)
		}
	}
}

func TestSmoothTriangleNormal(t *testing.T) {
	tri := SmoothTriangle(data.Point(0, 1, 0), data.Point(-1, 0, 0), data.Point(1, 0, 0),
		data.Vector(0, 1, 0), data.Vector(-1, 0, 0), data.Vector(1, 0, 0))
	p := data.Point(0, 0, 0)

	expected := data.Vector(-0.5547, 0.83205, 0)

	i := IntersectionWithUv(1, tri, 0.45, 0.25)
	n := NormalAt(tri, p, i)

	if !data.TupleEqual(n, expected) {
		t.Errorf("Local normal mismatch for point %v expected %v received %v", p, expected, n)
	}

	r := data.Ray(data.Point(-0.2, 0.3, -2), data.Vector(0, 0, 1))
	comps := i.PrepareComputations(r, Intersections(i)...)

	if !data.TupleEqual(comps.NormalV, expected) {
		t.Errorf("Normal mismatch for point %v expected %v received %v", p, expected, n)
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
		{
			t: SmoothTriangle(data.Point(0, 1, 0), data.Point(-1, 0, 0), data.Point(1, 0, 0),
				data.Vector(0, 1, 0), data.Vector(-1, 0, 0), data.Vector(1, 0, 0)),
			r:        data.Ray(data.Point(-0.2, 0.3, -2), data.Vector(0, 0, 1)),
			expected: Intersections(IntersectionWithUv(2, &TriangleType{}, 0.45, 0.25)),
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

			if !data.FloatEqual(x.U, tc.expected[i].U) {
				fmt.Printf("Failed: %+v", tc.expected[i])
				t.Errorf("Intersection %d u mismatch expected %.2f received %.2f", i, tc.expected[i].U, x.U)
			}

			if !data.FloatEqual(x.V, tc.expected[i].V) {
				t.Errorf("Intersection %d v mismatch expected %.2f received %.2f", i, tc.expected[i].V, x.V)
			}
		}
	}
}
