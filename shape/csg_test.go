package shape

import (
	"testing"

	"github.com/dannyroes/raytrace/data"
)

func TestCreateCsg(t *testing.T) {
	s1 := Sphere()
	s2 := Cube()

	csg := Csg(CsgUnion, s1, s2)

	if csg.operation != CsgUnion {
		t.Error("Operation is not union")
	}

	if csg.left != s1 {
		t.Error("Left is not sphere")
	}

	if csg.right != s2 {
		t.Error("Right is not cube")
	}

	if s1.Parent != csg {
		t.Error("Sphere parent is not CSG")
	}

	if s2.Parent != csg {
		t.Error("Cube parent is not CSG")
	}
}

func TestIntersectionAllowed(t *testing.T) {
	cases := []struct {
		op       CsgOperation
		lHit     bool
		inL      bool
		inR      bool
		expected bool
	}{
		{CsgUnion, true, true, true, false},
		{CsgUnion, true, true, false, true},
		{CsgUnion, true, false, true, false},
		{CsgUnion, true, false, false, true},
		{CsgUnion, false, true, true, false},
		{CsgUnion, false, true, false, false},
		{CsgUnion, false, false, true, true},
		{CsgUnion, false, false, false, true},
		{CsgIntersection, true, true, true, true},
		{CsgIntersection, true, true, false, false},
		{CsgIntersection, true, false, true, true},
		{CsgIntersection, true, false, false, false},
		{CsgIntersection, false, true, true, true},
		{CsgIntersection, false, true, false, true},
		{CsgIntersection, false, false, true, false},
		{CsgIntersection, false, false, false, false},
		{CsgDifference, true, true, true, false},
		{CsgDifference, true, true, false, true},
		{CsgDifference, true, false, true, false},
		{CsgDifference, true, false, false, true},
		{CsgDifference, false, true, true, true},
		{CsgDifference, false, true, false, true},
		{CsgDifference, false, false, true, false},
		{CsgDifference, false, false, false, false},
	}

	for _, tc := range cases {
		result := intersectionAllowed(tc.op, tc.lHit, tc.inL, tc.inR)

		if result != tc.expected {
			t.Errorf("IntersectionAllowed(%d, %v, %v, %v) expected %v received %v", tc.op, tc.lHit, tc.inL, tc.inR, tc.expected, result)
		}
	}
}

func TestIntersectionFilter(t *testing.T) {
	s1 := Sphere()
	s2 := Cube()

	xs := Intersections(
		Intersection(1, s1),
		Intersection(2, s2),
		Intersection(3, s1),
		Intersection(4, s2),
	)

	cases := []struct {
		op CsgOperation
		x0 int
		x1 int
	}{
		{op: CsgUnion, x0: 0, x1: 3},
		{op: CsgIntersection, x0: 1, x1: 2},
		{op: CsgDifference, x0: 0, x1: 1},
	}

	for _, tc := range cases {
		csg := Csg(tc.op, s1, s2)
		result := csg.filterIntersections(xs)

		if len(result) != 2 {
			t.Errorf("filterIntersections result for %d has %d intersections", tc.op, len(result))
		} else {
			if result[0] != xs[tc.x0] {
				t.Errorf("Intersection 0 mismatch: expected %+v received %+v", xs[tc.x0], result[0])
			}

			if result[1] != xs[tc.x1] {
				t.Errorf("Intersection 1 mismatch: expected %+v received %+v", xs[tc.x1], result[1])
			}
		}
	}
}

func TestCsgIntersect(t *testing.T) {
	c := Csg(CsgUnion, Sphere(), Cube())
	r := data.Ray(data.Point(0, 2, -5), data.Vector(0, 0, -1))
	xs := c.LocalIntersect(r)

	if len(xs) != 0 {
		t.Error("CSG miss len is not 0")
	}

	s1 := Sphere()
	s2 := Sphere()
	s2.SetTransform(data.Translation(0, 0, 0.5))

	c = Csg(CsgUnion, s1, s2)
	r = data.Ray(data.Point(0, 0, -5), data.Vector(0, 0, 1))

	xs = c.LocalIntersect(r)

	if len(xs) != 2 {
		t.Errorf("CSG hit expected 2 received %d", len(xs))
	} else {
		if xs[0].T != 4 || xs[0].Object != s1 {
			t.Errorf("CSG hit 1 mismatch %+v", xs[0])
		}

		if xs[1].T != 6.5 || xs[1].Object != s2 {
			t.Errorf("CSG hit 1 mismatch %+v", xs[1])
		}
	}
}
