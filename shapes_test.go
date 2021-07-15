package main

import (
	"math"
	"testing"
)

type MockShape struct{}

var savedRay RayType
var transform Matrix

func (s *MockShape) GetMaterial() MaterialType {
	return Material()
}

func (s *MockShape) SetMaterial(m MaterialType) {

}

func (s *MockShape) GetTransform() Matrix {
	return transform
}

func (s *MockShape) SetTransform(m Matrix) {
	transform = m
}

func (s *MockShape) LocalIntersect(r RayType) IntersectionList {
	savedRay = r
	return IntersectionList{}
}

func (s *MockShape) LocalNormalAt(t Tuple) Tuple {
	return Vector(t.X, t.Y, t.Z)
}

func TestIntersect(t *testing.T) {
	cases := []struct {
		t        Matrix
		expected RayType
	}{
		{
			t:        Scaling(2, 2, 2),
			expected: Ray(Point(0, 0, -2.5), Vector(0, 0, 0.5)),
		},
		{
			t:        Translation(5, 0, 0),
			expected: Ray(Point(-5, 0, -5), Vector(0, 0, 1)),
		},
	}

	for _, tc := range cases {
		r := Ray(Point(0, 0, -5), Vector(0, 0, 1))

		s := &MockShape{}
		s.SetTransform(tc.t)

		Intersects(s, r)

		if !TupleEqual(tc.expected.Origin, savedRay.Origin) {
			t.Errorf("Transformed ray origin mismatch expected %v received %v", tc.expected.Origin, savedRay.Origin)
		}

		if !TupleEqual(tc.expected.Direction, savedRay.Direction) {
			t.Errorf("Transformed ray direction mismatch expected %v received %v", tc.expected.Direction, savedRay.Direction)
		}
	}
}

func TestNormalAt(t *testing.T) {
	cases := []struct {
		t        Matrix
		p        Tuple
		expected Tuple
	}{
		{
			t:        Translation(0, 1, 0),
			p:        Point(0, 1.70711, -0.70711),
			expected: Vector(0, 0.70711, -0.70711),
		},
		{
			t:        IdentityMatrix().RotateZ(math.Pi/5).Scale(1, 0.5, 1),
			p:        Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2*-1),
			expected: Vector(0, 0.97014, -0.24254),
		},
	}

	for _, tc := range cases {
		s := &MockShape{}
		s.SetTransform(tc.t)

		result := NormalAt(s, tc.p)

		if !TupleEqual(tc.expected, result) {
			t.Errorf("Normal at mismatch expected %v received %v", tc.expected, result)
		}
	}
}
