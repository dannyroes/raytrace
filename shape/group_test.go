package shape

import (
	"math"
	"testing"

	"github.com/dannyroes/raytrace/data"
)

func TestGroup(t *testing.T) {
	g := Group()

	if !g.Transform.Equals(data.IdentityMatrix()) {
		t.Errorf("Default transform is not identity matrix")
	}

	if len(g.Children) != 0 {
		t.Errorf("Group contains unexpected children")
	}
}

func TestAddChild(t *testing.T) {
	g := Group()
	s := &MockShape{}

	g.AddChild(s)

	if s.GetParent() != g {
		t.Errorf("Shape parent mismatch expected %v received %v", g, s.GetParent())
	}

	if len(g.Children) < 1 {
		t.Errorf("Group doesn't contain enough items, received %d", len(g.Children))
	} else {
		found := false
		for _, shape := range g.Children {
			if shape == s {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Shape not found in group children")
		}
	}
}

func TestGroupLocalIntersect(t *testing.T) {
	s1 := Sphere()

	s2 := Sphere()
	s2.SetTransform(data.Translation(0, 0, -3))

	s3 := Sphere()
	s3.SetTransform(data.Translation(5, 0, 0))

	cases := []struct {
		g        *GroupType
		r        data.RayType
		expected IntersectionList
	}{
		{
			g:        Group(),
			r:        data.Ray(data.Point(0, 0, 0), data.Vector(0, 0, 1)),
			expected: Intersections(),
		},
		{
			g: func() *GroupType {
				g := Group()
				g.AddChild(s1, s2, s3)

				return g
			}(),
			r:        data.Ray(data.Point(0, 0, -5), data.Vector(0, 0, 1)),
			expected: Intersections(Intersection(0, s2), Intersection(1, s2), Intersection(2, s1), Intersection(2, s1)),
		},
	}

	for _, c := range cases {
		xs := c.g.LocalIntersect(c.r)
		if len(c.expected) != len(xs) {
			t.Errorf("Wrong number of intersections expected %d received %d", len(c.expected), len(xs))
		} else {
			for i, x := range xs {
				if c.expected[i].Object != x.Object {
					t.Errorf("Wrong shape and index %d expected %+v received %+v", i, c.expected[i].Object, x.Object)
				}
			}
		}
	}
}

func TestGroupIntersect(t *testing.T) {
	s3 := Sphere()
	s3.SetTransform(data.Translation(5, 0, 0))

	cases := []struct {
		g        *GroupType
		r        data.RayType
		expected IntersectionList
	}{
		{
			g: func() *GroupType {
				g := Group()
				g.SetTransform(data.Scaling(2, 2, 2))
				g.AddChild(s3)

				return g
			}(),
			r:        data.Ray(data.Point(10, 0, -10), data.Vector(0, 0, 1)),
			expected: Intersections(Intersection(0, s3), Intersection(1, s3)),
		},
	}

	for _, c := range cases {
		xs := Intersects(c.g, c.r)
		if len(c.expected) != len(xs) {
			t.Errorf("Wrong number of intersections expected %d received %d", len(c.expected), len(xs))
		} else {
			for i, x := range xs {
				if c.expected[i].Object != x.Object {
					t.Errorf("Wrong shape and index %d expected %+v received %+v", i, c.expected[i].Object, x.Object)
				}
			}
		}
	}
}

func TestWorldToObject(t *testing.T) {
	g1 := Group()
	g1.SetTransform(data.RotateY(math.Pi / 2))

	g2 := Group()
	g2.SetTransform(data.Scaling(2, 2, 2))
	g1.AddChild(g2)

	s := Sphere()
	s.SetTransform(data.Translation(5, 0, 0))
	g2.AddChild(s)

	point := worldToObject(s, data.Point(-2, 0, -10))
	if !data.TupleEqual(point, data.Point(0, 0, -1)) {
		t.Errorf("Bad object point expected %v received %v", data.Point(0, 0, -1), point)
	}
}

func TestNormalToWorld(t *testing.T) {
	g1 := Group()
	g1.SetTransform(data.RotateY(math.Pi / 2))

	g2 := Group()
	g2.SetTransform(data.Scaling(1, 2, 3))
	g1.AddChild(g2)

	s := Sphere()
	s.SetTransform(data.Translation(5, 0, 0))
	g2.AddChild(s)

	sqrt3 := math.Sqrt(3) / 3

	expected := data.Vector(0.28571, 0.42857, -0.85714)

	point := normalToWorld(s, data.Vector(sqrt3, sqrt3, sqrt3))
	if !data.TupleEqual(point, expected) {
		t.Errorf("Bad normal expected %v received %v", expected, point)
	}
}

func TestGroupNormal(t *testing.T) {
	g1 := Group()
	g1.SetTransform(data.RotateY(math.Pi / 2))

	g2 := Group()
	g2.SetTransform(data.Scaling(1, 2, 3))
	g1.AddChild(g2)

	s := Sphere()
	s.SetTransform(data.Translation(5, 0, 0))
	g2.AddChild(s)

	expected := data.Vector(0.28570, 0.42854, -0.85716)

	point := NormalAt(s, data.Point(1.7321, 1.1547, -5.5774))
	if !data.TupleEqual(point, expected) {
		t.Errorf("Bad normal expected %v received %v", expected, point)
	}
}
