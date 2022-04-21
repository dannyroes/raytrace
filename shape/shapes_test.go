package shape

import (
	"math"
	"testing"

	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
)

type MockShape struct{}

var savedRay data.RayType
var transform data.Matrix
var parent *GroupType

func (s *MockShape) GetMaterial() material.MaterialType {
	return material.Material()
}

func (s *MockShape) SetMaterial(m material.MaterialType) {

}

func (s *MockShape) GetTransform() data.Matrix {
	return transform
}

func (s *MockShape) SetTransform(m data.Matrix) {
	transform = m
}

func (s *MockShape) LocalIntersect(r data.RayType) IntersectionList {
	savedRay = r
	return IntersectionList{}
}

func (s *MockShape) LocalNormalAt(t data.Tuple, i IntersectionType) data.Tuple {
	return data.Vector(t.X, t.Y, t.Z)
}

func (s *MockShape) GetParent() *GroupType {
	return parent
}

func (s *MockShape) SetParent(p *GroupType) {
	parent = p
}

func (s *MockShape) Bounds() Bounds {
	return Bounds{
		Min: data.Point(-1, -1, -1),
		Max: data.Point(1, 1, 1),
	}
}

func TestIntersect(t *testing.T) {
	cases := []struct {
		t        data.Matrix
		expected data.RayType
	}{
		{
			t:        data.Scaling(2, 2, 2),
			expected: data.Ray(data.Point(0, 0, -2.5), data.Vector(0, 0, 0.5)),
		},
		{
			t:        data.Translation(5, 0, 0),
			expected: data.Ray(data.Point(-5, 0, -5), data.Vector(0, 0, 1)),
		},
	}

	for _, tc := range cases {
		r := data.Ray(data.Point(0, 0, -5), data.Vector(0, 0, 1))

		s := &MockShape{}
		s.SetTransform(tc.t)

		Intersects(s, r)

		if !data.TupleEqual(tc.expected.Origin, savedRay.Origin) {
			t.Errorf("Transformed ray origin mismatch expected %v received %v", tc.expected.Origin, savedRay.Origin)
		}

		if !data.TupleEqual(tc.expected.Direction, savedRay.Direction) {
			t.Errorf("Transformed ray direction mismatch expected %v received %v", tc.expected.Direction, savedRay.Direction)
		}
	}
}

func TestNormalAt(t *testing.T) {
	cases := []struct {
		t        data.Matrix
		p        data.Tuple
		expected data.Tuple
	}{
		{
			t:        data.Translation(0, 1, 0),
			p:        data.Point(0, 1.70711, -0.70711),
			expected: data.Vector(0, 0.70711, -0.70711),
		},
		{
			t:        data.IdentityMatrix().RotateZ(math.Pi/5).Scale(1, 0.5, 1),
			p:        data.Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2*-1),
			expected: data.Vector(0, 0.97014, -0.24254),
		},
	}

	for _, tc := range cases {
		s := &MockShape{}
		s.SetTransform(tc.t)

		result := NormalAt(s, tc.p, IntersectionType{})

		if !data.TupleEqual(tc.expected, result) {
			t.Errorf("Normal at mismatch expected %v received %v", tc.expected, result)
		}
	}
}

func TestPatternTransform(t *testing.T) {
	cases := []struct {
		o        Shape
		p        material.Pattern
		point    data.Tuple
		expected material.ColourTuple
	}{
		{
			o: func() *SphereType {
				s := Sphere()
				s.SetTransform(data.Scaling(2, 2, 2))

				return s
			}(),
			p:        material.StripePattern(material.White, material.Black),
			point:    data.Point(1.5, 0, 0),
			expected: material.White,
		},
		{
			o: Sphere(),
			p: func() material.Pattern {
				p := material.StripePattern(material.White, material.Black)
				p.SetTransform(data.Scaling(2, 2, 2))
				return p
			}(),
			point:    data.Point(1.5, 0, 0),
			expected: material.White,
		},
		{
			o: func() *SphereType {
				s := Sphere()
				s.SetTransform(data.Scaling(2, 2, 2))

				return s
			}(),
			p: func() material.Pattern {
				p := material.StripePattern(material.White, material.Black)
				p.SetTransform(data.Translation(0.5, 0, 0))
				return p
			}(),
			point:    data.Point(2.5, 0, 0),
			expected: material.White,
		},
	}

	for _, tc := range cases {
		at := PatternAtObject(tc.p, tc.o, tc.point)
		if !material.ColourEqual(at, tc.expected) {
			t.Errorf("Pattern mismatch at %+v expected %+v received %+v", tc.point, tc.expected, at)
		}
	}
}

func TestParent(t *testing.T) {
	parent = nil
	s := MockShape{}

	if s.GetParent() != nil {
		t.Errorf("Mock shape has a parent")
	}
}
