package shape

import (
	"math"
	"testing"

	"github.com/dannyroes/raytrace/data"
)

func TestMissCylinder(t *testing.T) {
	cases := []struct {
		origin    data.Tuple
		direction data.Tuple
	}{
		{
			origin:    data.Point(1, 0, 0),
			direction: data.Vector(0, 1, 0),
		},
		{
			origin:    data.Point(0, 0, 0),
			direction: data.Vector(0, 1, 0),
		},
		{
			origin:    data.Point(0, 0, -5),
			direction: data.Vector(1, 1, 1),
		},
	}

	for _, tc := range cases {
		c := Cylinder()
		d := tc.direction.Normalize()
		r := data.Ray(tc.origin, d)

		ints := c.LocalIntersect(r)
		if len(ints) != 0 {
			t.Errorf("Unexpected hit %+v", ints)
		}
	}
}

func TestHitCylinder(t *testing.T) {
	cases := []struct {
		origin    data.Tuple
		direction data.Tuple
		t0        float64
		t1        float64
	}{
		{
			origin:    data.Point(1, 0, -5),
			direction: data.Vector(0, 0, 1),
			t0:        5,
			t1:        5,
		},
		{
			origin:    data.Point(0, 0, -5),
			direction: data.Vector(0, 0, 1),
			t0:        4,
			t1:        6,
		},
		{
			origin:    data.Point(0.5, 0, -5),
			direction: data.Vector(0.1, 1, 1),
			t0:        6.80798,
			t1:        7.08872,
		},
	}

	for _, tc := range cases {
		c := Cylinder()
		d := tc.direction.Normalize()
		r := data.Ray(tc.origin, d)

		ints := c.LocalIntersect(r)
		if len(ints) != 2 {
			t.Errorf("Unexpected intersections length %d", len(ints))
		} else {
			if !data.FloatEqual(ints[0].T, tc.t0) || !data.FloatEqual(ints[1].T, tc.t1) {
				t.Errorf("Hit mismatch expected %f, %f received %f, %f", tc.t0, tc.t1, ints[0].T, ints[1].T)
			}
		}
	}
}

func TestCylinderNormal(t *testing.T) {
	cases := []struct {
		point  data.Tuple
		normal data.Tuple
	}{
		{
			point:  data.Point(1, 0, 0),
			normal: data.Vector(1, 0, 0),
		},
		{
			point:  data.Point(0, 5, -1),
			normal: data.Vector(0, 0, -1),
		},
		{
			point:  data.Point(0, -2, 1),
			normal: data.Vector(0, 0, 1),
		},
		{
			point:  data.Point(-1, 1, 0),
			normal: data.Vector(-1, 0, 0),
		},
	}

	for _, tc := range cases {
		c := Cylinder()
		n := c.LocalNormalAt(tc.point, IntersectionType{})

		if !data.TupleEqual(n, tc.normal) {
			t.Errorf("Normal mismatch expected %+v received %+v", tc.normal, n)
		}
	}
}

func TestCylinderAttributes(t *testing.T) {
	c := Cylinder()

	if c.Minimum != math.Inf(-1) {
		t.Error("Min is not -inf")
	}

	if c.Maximum != math.Inf(1) {
		t.Error("Max is not inf")
	}

	if c.Closed {
		t.Error("Cylinder is not closed")
	}
}

func TestTruncateCylinder(t *testing.T) {
	cases := []struct {
		point     data.Tuple
		direction data.Tuple
		count     int
	}{
		{
			point:     data.Point(0, 1.5, 0),
			direction: data.Vector(0.1, 1, 0),
			count:     0,
		},
		{
			point:     data.Point(0, 3, -5),
			direction: data.Vector(0, 0, 1),
			count:     0,
		},
		{
			point:     data.Point(0, 0, -5),
			direction: data.Vector(0, 0, 1),
			count:     0,
		},
		{
			point:     data.Point(0, 2, -5),
			direction: data.Vector(0, 0, 1),
			count:     0,
		},
		{
			point:     data.Point(0, 1, -5),
			direction: data.Vector(0, 0, 1),
			count:     0,
		},
		{
			point:     data.Point(0, 1.5, -2),
			direction: data.Vector(0, 0, 1),
			count:     2,
		},
	}

	for _, tc := range cases {
		c := Cylinder()
		c.Minimum = 1
		c.Maximum = 2
		d := tc.direction.Normalize()
		r := data.Ray(tc.point, d)

		ints := c.LocalIntersect(r)
		if len(ints) != tc.count {
			t.Errorf("Unexpected hit count expected %d received %d", tc.count, len(ints))
		}
	}
}

func TestClosedCylinder(t *testing.T) {
	cases := []struct {
		point     data.Tuple
		direction data.Tuple
		count     int
	}{
		{
			point:     data.Point(0, 3, 0),
			direction: data.Vector(0, -1, 0),
			count:     2,
		},
		{
			point:     data.Point(0, 3, -2),
			direction: data.Vector(0, -1, 2),
			count:     2,
		},
		{
			point:     data.Point(0, 4, -2),
			direction: data.Vector(0, -1, 1),
			count:     2,
		},
		{
			point:     data.Point(0, 0, -2),
			direction: data.Vector(0, 1, 2),
			count:     2,
		},
		{
			point:     data.Point(0, 0, -2),
			direction: data.Vector(0, 1, 2),
			count:     2,
		},
		{
			point:     data.Point(0, -1, -2),
			direction: data.Vector(0, 1, 1),
			count:     2,
		},
	}

	for _, tc := range cases {
		c := Cylinder()
		c.Minimum = 1
		c.Maximum = 2
		c.Closed = true

		d := tc.direction.Normalize()
		r := data.Ray(tc.point, d)

		ints := c.LocalIntersect(r)
		if len(ints) != tc.count {
			t.Errorf("Unexpected hit count expected %d received %d", tc.count, len(ints))
		}
	}
}

func TestCylinderEndcapNormal(t *testing.T) {
	cases := []struct {
		point  data.Tuple
		normal data.Tuple
	}{
		{
			point:  data.Point(0, 1, 0),
			normal: data.Vector(0, -1, 0),
		},
		{
			point:  data.Point(0.5, 1, 0),
			normal: data.Vector(0, -1, 0),
		},
		{
			point:  data.Point(0, 1, 0.5),
			normal: data.Vector(0, -1, 0),
		},
		{
			point:  data.Point(0, 2, 0),
			normal: data.Vector(0, 1, 0),
		},
		{
			point:  data.Point(0.5, 2, 0),
			normal: data.Vector(0, 1, 0),
		},
		{
			point:  data.Point(0, 2, 0.5),
			normal: data.Vector(0, 1, 0),
		},
	}

	for _, tc := range cases {
		c := Cylinder()
		c.Minimum = 1
		c.Maximum = 2
		c.Closed = true
		n := c.LocalNormalAt(tc.point, IntersectionType{})

		if !data.TupleEqual(n, tc.normal) {
			t.Errorf("Normal mismatch expected %+v received %+v", tc.normal, n)
		}
	}
}
