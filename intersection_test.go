package main

import (
	"testing"
)

func TestIntersectionType(t *testing.T) {
	cases := []struct {
		s *SphereType
		t float64
	}{
		{
			s: Sphere(1),
			t: 3.5,
		},
	}

	for _, tc := range cases {
		i := Intersection(tc.t, tc.s)

		if !FloatEqual(tc.t, i.T) {
			t.Errorf("T not equal expected %f received %F", tc.t, i.T)
		}

		if tc.s.Id != i.Object.GetId() {
			t.Errorf("Object ID not equal expected %d received %d", tc.s.Id, i.Object.GetId())
		}
	}
}

func TestIntersectionList(t *testing.T) {
	s := Sphere(1)
	i1 := Intersection(1, s)
	i2 := Intersection(2, s)

	l := Intersections(i1, i2)

	if len(l) != 2 {
		t.Errorf("List length incorrect expected %d received %d", 2, len(l))
	}

	if !FloatEqual(l[0].T, 1) {
		t.Errorf("Intersection 1 incorrect expected %f received %f", 1.0, l[0].T)
	}

	if !FloatEqual(l[1].T, 2) {
		t.Errorf("Intersection 2 incorrect expected %f received %f", 2.0, l[1].T)
	}
}

func TestHit(t *testing.T) {
	s := Sphere(1)
	cases := []struct {
		l        IntersectionList
		expected IntersectionType
	}{
		{
			l: Intersections(
				Intersection(1, s),
				Intersection(2, s),
			),
			expected: Intersection(1, s),
		},
		{
			l: Intersections(
				Intersection(-1, s),
				Intersection(1, s),
			),
			expected: Intersection(1, s),
		},
		{
			l: Intersections(
				Intersection(-2, s),
				Intersection(-1, s),
			),
			expected: IntersectionType{T: -1},
		},
		{
			l: Intersections(
				Intersection(5, s),
				Intersection(7, s),
				Intersection(-3, s),
				Intersection(2, s),
			),
			expected: Intersection(2, s),
		},
	}

	for _, tc := range cases {
		hit := tc.l.Hit()
		if hit.Object == nil {
			if tc.expected.Object != nil {
				t.Errorf("Hit does not match expected %+v received %+v", tc.expected, hit)
			}
			continue
		}
		if tc.expected.Object.GetId() != hit.Object.GetId() || tc.expected.T != hit.T {
			t.Errorf("Hit does not match expected %+v received %+v", tc.expected, hit)
		}
	}
}

func TestPrepareComputations(t *testing.T) {
	defaultSphere := Sphere(1)
	tests := []struct {
		r        RayType
		o        Object
		i        IntersectionType
		expected Computations
	}{
		{
			r: Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			o: defaultSphere,
			i: Intersection(4, defaultSphere),
			expected: Computations{
				T:       4,
				Object:  defaultSphere,
				Point:   Point(0, 0, -1),
				EyeV:    Vector(0, 0, -1),
				NormalV: Vector(0, 0, -1),
			},
		},
		{
			r: Ray(Point(0, 0, 0), Vector(0, 0, 1)),
			o: defaultSphere,
			i: Intersection(1, defaultSphere),
			expected: Computations{
				T:       1,
				Object:  defaultSphere,
				Point:   Point(0, 0, 1),
				EyeV:    Vector(0, 0, -1),
				NormalV: Vector(0, 0, -1),
				Inside:  true,
			},
		},
	}

	for _, tc := range tests {
		comp := tc.i.PrepareComputations(tc.r)
		if comp.T != tc.expected.T {
			t.Errorf("T value doesn't match expected: %f received: %f", tc.expected.T, comp.T)
		}

		if comp.Object.GetId() != tc.expected.Object.GetId() {
			t.Errorf("Object doesn't match expected: %+v received: %+v", tc.expected.Object, comp.Object)
		}

		if !TupleEqual(comp.Point, tc.expected.Point) {
			t.Errorf("Point doesn't match expected: %+v received: %+v", tc.expected.Point, comp.Point)
		}

		if !TupleEqual(comp.EyeV, tc.expected.EyeV) {
			t.Errorf("EyeV value doesn't match expected: %+v received: %+v", tc.expected.EyeV, comp.EyeV)
		}

		if !TupleEqual(comp.NormalV, tc.expected.NormalV) {
			t.Errorf("NormalV value doesn't match expected: %+v received: %+v", tc.expected.NormalV, comp.NormalV)
		}

		if comp.Inside != tc.expected.Inside {
			t.Errorf("Inside value doesn't match expected: %v received: %v", tc.expected.Inside, comp.Inside)
		}
	}
}

func TestOffsetHit(t *testing.T) {
	s := Sphere(1)
	s.Transform = Translation(0, 0, 1)
	cases := []struct {
		r      RayType
		s      *SphereType
		i      IntersectionType
		offset float64
	}{
		{
			r:      Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			s:      s,
			i:      Intersection(5, s),
			offset: -1 * (Epsilon / 2.0),
		},
	}

	for _, tc := range cases {
		c := tc.i.PrepareComputations(tc.r)
		if c.OverPoint.Z >= tc.offset {
			t.Errorf("Offset not applied expected <: %v received: %v", tc.offset, c.OverPoint.Z)
		}

		if c.Point.Z <= c.OverPoint.Z {
			t.Errorf("Offset not applied Point <: %v OverPoint: %v", c.Point.Z, c.OverPoint.Z)
		}
	}
}
