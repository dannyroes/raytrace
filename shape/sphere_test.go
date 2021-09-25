package shape

import (
	"math"
	"testing"

	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
)

func TestSphereTransform(t *testing.T) {
	cases := []struct {
		s        *SphereType
		expected data.Matrix
	}{
		{
			s:        Sphere(),
			expected: data.IdentityMatrix(),
		},
		{
			s:        sphereWithTransform(data.Translation(2, 3, 4)),
			expected: data.Translation(2, 3, 4),
		},
	}

	for _, tc := range cases {
		if !tc.expected.Equals(tc.s.Transform) {
			t.Errorf("transform mismatch expected %+v received %+v", tc.expected, tc.s.Transform)
		}
	}
}

func TestSphereIntersect(t *testing.T) {
	cases := []struct {
		s        *SphereType
		r        data.RayType
		expected []float64
	}{
		{
			s:        Sphere(),
			r:        data.Ray(data.Point(0, 0, -5), data.Vector(0, 0, 1)),
			expected: []float64{4, 6},
		},
		{
			s:        Sphere(),
			r:        data.Ray(data.Point(0, 1, -5), data.Vector(0, 0, 1)),
			expected: []float64{5, 5},
		},
		{
			s:        Sphere(),
			r:        data.Ray(data.Point(0, 2, -5), data.Vector(0, 0, 1)),
			expected: []float64{},
		},
		{
			s:        Sphere(),
			r:        data.Ray(data.Point(0, 0, 0), data.Vector(0, 0, 1)),
			expected: []float64{-1, 1},
		},
		{
			s:        Sphere(),
			r:        data.Ray(data.Point(0, 0, 5), data.Vector(0, 0, 1)),
			expected: []float64{-6, -4},
		},
	}

	for _, tc := range cases {
		result := Intersects(tc.s, tc.r)

		if len(result) != len(tc.expected) {
			t.Errorf("Result length mismatch expected %d, received %d", len(tc.expected), len(result))
		}

		for i := range result {
			if tc.expected[i] != result[i].T {
				t.Errorf("Result index %d mismatch expected %d, received %d", i, len(tc.expected), len(result))
			}

			if result[i].Object != tc.s {
				t.Errorf("Result index %d id mismatch expected %p, received %p", i, tc.s, result[i].Object)
			}
		}
	}
}

func TestSphereIntersection(t *testing.T) {
	cases := []struct {
		s        *SphereType
		r        data.RayType
		expected IntersectionList
	}{
		{
			s: Sphere(),
			r: data.Ray(data.Point(0, 0, -2.5), data.Vector(0, 0, 0.5)),
			expected: Intersections(
				Intersection(3, Sphere()),
				Intersection(7, Sphere()),
			),
		},
		{
			s:        Sphere(),
			r:        data.Ray(data.Point(-5, 0, -5), data.Vector(0, 0, 1)),
			expected: Intersections(),
		},
	}

	for _, tc := range cases {
		result := tc.s.LocalIntersect(tc.r)

		if len(tc.expected) != len(result) {
			t.Errorf("intersection length mismatch expected %d received %d", len(tc.expected), len(result))
		}

		for i := range tc.expected {
			if !data.FloatEqual(tc.expected[i].T, result[i].T) {
				t.Errorf("intersection mismatch expected %f received %f", tc.expected[i].T, result[i].T)
			}
		}
	}
}

func TestSphereNormal(t *testing.T) {
	angle := math.Sqrt(3) / 3.0

	cases := []struct {
		s        *SphereType
		p        data.Tuple
		expected data.Tuple
	}{
		{
			s:        Sphere(),
			p:        data.Point(1, 0, 0),
			expected: data.Vector(1, 0, 0),
		},
		{
			s:        Sphere(),
			p:        data.Point(0, 1, 0),
			expected: data.Vector(0, 1, 0),
		},
		{
			s:        Sphere(),
			p:        data.Point(0, 0, 1),
			expected: data.Vector(0, 0, 1),
		},
		{
			s:        Sphere(),
			p:        data.Point(angle, angle, angle),
			expected: data.Vector(angle, angle, angle),
		},
	}

	for _, tc := range cases {
		result := tc.s.LocalNormalAt(tc.p)

		if !data.TupleEqual(result, result.Normalize()) {
			t.Errorf("normal is not normalized expected %v received %v", result.Normalize(), result)
		}

		if !data.TupleEqual(tc.expected, result) {
			t.Errorf("wrong normal expected %v received %v", tc.expected, result)
		}
	}
}

func TestSphereMaterial(t *testing.T) {
	s := Sphere()
	m := material.Material()

	if s.Material != m {
		t.Errorf("Material mismatch expected: %+v received: %+v", m, s.Material)
	}

	m.Ambient = 1
	s.Material = m

	if s.Material != m {
		t.Errorf("Material mismatch expected: %+v received: %+v", m, s.Material)
	}
}

func sphereWithTransform(t data.Matrix) *SphereType {
	s := Sphere()
	s.SetTransform(t)
	return s
}
