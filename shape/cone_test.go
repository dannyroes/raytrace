package shape

import (
	"math"
	"testing"

	"github.com/dannyroes/raytrace/data"
)

func TestHitCone(t *testing.T) {
	cases := []struct {
		origin    data.Tuple
		direction data.Tuple
		t0        float64
		t1        float64
	}{
		{
			origin:    data.Point(0, 0, -5),
			direction: data.Vector(0, 0, 1),
			t0:        5,
			t1:        5,
		},
		{
			origin:    data.Point(0, 0, -5),
			direction: data.Vector(1, 1, 1),
			t0:        8.66025,
			t1:        8.66025,
		},
		{
			origin:    data.Point(1, 1, -5),
			direction: data.Vector(-0.5, -1, 1),
			t0:        4.55006,
			t1:        49.44994,
		},
	}

	for _, tc := range cases {
		c := Cone()
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

func TestHitConeParallel(t *testing.T) {
	cases := []struct {
		origin    data.Tuple
		direction data.Tuple
		t0        float64
	}{
		{
			origin:    data.Point(0, 0, -1),
			direction: data.Vector(0, 1, 1).Normalize(),
			t0:        0.35355,
		},
	}

	for _, tc := range cases {
		c := Cone()
		d := tc.direction.Normalize()
		r := data.Ray(tc.origin, d)

		ints := c.LocalIntersect(r)
		if len(ints) != 1 {
			t.Errorf("Unexpected intersections length %d", len(ints))
		} else {
			if !data.FloatEqual(ints[0].T, tc.t0) {
				t.Errorf("Hit mismatch expected %f received %f", tc.t0, ints[0].T)
			}
		}
	}
}

func TestClosedCone(t *testing.T) {
	cases := []struct {
		point     data.Tuple
		direction data.Tuple
		count     int
	}{
		{
			point:     data.Point(0, 0, -5),
			direction: data.Vector(0, 1, 0),
			count:     0,
		},
		{
			point:     data.Point(0, 0, -0.25),
			direction: data.Vector(0, 1, 1),
			count:     2,
		},
		{
			point:     data.Point(0, 0, -0.25),
			direction: data.Vector(0, 1, 0),
			count:     4,
		},
	}

	for _, tc := range cases {
		c := Cone()
		c.Minimum = -0.5
		c.Maximum = 0.5
		c.Closed = true

		d := tc.direction.Normalize()
		r := data.Ray(tc.point, d)

		ints := c.LocalIntersect(r)
		if len(ints) != tc.count {
			t.Errorf("Unexpected hit count expected %d received %d", tc.count, len(ints))
		}
	}
}

func TestConeNormal(t *testing.T) {
	cases := []struct {
		point  data.Tuple
		normal data.Tuple
	}{
		{
			point:  data.Point(0, 0, 0),
			normal: data.Vector(0, 0, 0),
		},
		{
			point:  data.Point(1, 1, 1),
			normal: data.Vector(1, -math.Sqrt2, 1),
		},
		{
			point:  data.Point(-1, -1, 0),
			normal: data.Vector(-1, 1, 0),
		},
	}

	for _, tc := range cases {
		c := Cone()
		n := c.LocalNormalAt(tc.point, IntersectionType{})

		if !data.TupleEqual(n, tc.normal) {
			t.Errorf("Normal mismatch expected %+v received %+v", tc.normal, n)
		}
	}
}
