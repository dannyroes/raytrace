package main

import "testing"

func TestIntersectionType(t *testing.T) {
	cases := []struct {
		s SphereType
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

		if tc.s.Id != i.Object.Id {
			t.Errorf("Object ID not equal expected %d received %d", tc.s.Id, i.Object.Id)
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
			expected: IntersectionType{},
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
		if tc.expected.Object.Id != hit.Object.Id || tc.expected.T != hit.T {
			t.Errorf("Hit does not match expected %+v received %+v", tc.expected, hit)
		}
	}
}
