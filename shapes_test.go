package main

import (
	"math"
	"testing"
)

func TestSphereTransform(t *testing.T) {
	cases := []struct {
		s        *SphereType
		expected Matrix
	}{
		{
			s:        Sphere(),
			expected: IdentityMatrix(),
		},
		{
			s:        sphereWithTransform(Translation(2, 3, 4)),
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
		s        *SphereType
		r        RayType
		expected IntersectionList
	}{
		{
			s: sphereWithTransform(Scaling(2, 2, 2)),
			r: Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			expected: Intersections(
				Intersection(3, Sphere()),
				Intersection(7, Sphere()),
			),
		},
		{
			s:        sphereWithTransform(Translation(5, 0, 0)),
			r:        Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			expected: Intersections(),
		},
	}

	for _, tc := range cases {
		result := Intersects(tc.s, tc.r)

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

func TestSphereNormal(t *testing.T) {
	angle := math.Sqrt(3) / 3.0

	cases := []struct {
		s        *SphereType
		p        Tuple
		expected Tuple
	}{
		{
			s:        Sphere(),
			p:        Point(1, 0, 0),
			expected: Vector(1, 0, 0),
		},
		{
			s:        Sphere(),
			p:        Point(0, 1, 0),
			expected: Vector(0, 1, 0),
		},
		{
			s:        Sphere(),
			p:        Point(0, 0, 1),
			expected: Vector(0, 0, 1),
		},
		{
			s:        Sphere(),
			p:        Point(angle, angle, angle),
			expected: Vector(angle, angle, angle),
		},
		{
			s:        sphereWithTransform(IdentityMatrix().Translate(0, 1, 0)),
			p:        Point(0, 1.70711, -0.70711),
			expected: Vector(0, 0.70711, -0.70711),
		},
		{
			s:        sphereWithTransform(IdentityMatrix().RotateZ(math.Pi/5).Scale(1, 0.5, 1)),
			p:        Point(0, math.Sqrt(2)/2, -1*math.Sqrt(2)/2),
			expected: Vector(0, 0.97014, -0.24254),
		},
	}

	for _, tc := range cases {
		result := NormalAt(tc.s, tc.p)

		if !TupleEqual(result, result.Normalize()) {
			t.Errorf("normal is not normalized expected %v received %v", result.Normalize(), result)
		}

		if !TupleEqual(tc.expected, result) {
			t.Errorf("wrong normal expected %v received %v", tc.expected, result)
		}
	}
}

func TestSphereMaterial(t *testing.T) {
	s := Sphere()
	m := Material()

	if s.Material != m {
		t.Errorf("Material mismatch expected: %+v received: %+v", m, s.Material)
	}

	m.Ambient = 1
	s.Material = m

	if s.Material != m {
		t.Errorf("Material mismatch expected: %+v received: %+v", m, s.Material)
	}
}

// func TestPlaneNormal(t *testing.T) {
// 	p := Plane()

// 	n1 := p.
// }

func sphereWithTransform(t Matrix) *SphereType {
	s := Sphere()
	s.SetTransform(t)
	return s
}
