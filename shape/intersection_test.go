package shape

import (
	"math"
	"testing"

	"github.com/dannyroes/raytrace/data"
)

func TestIntersectionType(t *testing.T) {
	cases := []struct {
		s *SphereType
		t float64
	}{
		{
			s: Sphere(),
			t: 3.5,
		},
	}

	for _, tc := range cases {
		i := Intersection(tc.t, tc.s)

		if !data.FloatEqual(tc.t, i.T) {
			t.Errorf("T not equal expected %f received %F", tc.t, i.T)
		}

		if tc.s != i.Object.(*SphereType) {
			t.Errorf("Object ID not equal expected %p received %p", tc.s, i.Object)
		}
	}
}

func TestIntersectionWithUv(t *testing.T) {
	cases := []struct {
		s Shape
		t float64
		u float64
		v float64
	}{
		{
			s: Triangle(data.Point(0, 1, 0), data.Point(-1, 0, 0), data.Point(1, 0, 0)),
			t: 3.5,
			u: 0.2,
			v: 0.4,
		},
	}

	for _, tc := range cases {
		i := IntersectionWithUv(tc.t, tc.s, tc.u, tc.v)

		if !data.FloatEqual(tc.t, i.T) {
			t.Errorf("T not equal expected %f received %F", tc.t, i.T)
		}

		if tc.s != i.Object {
			t.Errorf("Object ID not equal expected %p received %p", tc.s, i.Object)
		}

		if tc.u != i.U {
			t.Errorf("U not equal expected %f received %f", tc.u, i.U)
		}

		if tc.v != i.V {
			t.Errorf("V ID not equal expected %f received %f", tc.v, i.V)
		}
	}
}

func TestIntersectionList(t *testing.T) {
	s := Sphere()
	i1 := Intersection(1, s)
	i2 := Intersection(2, s)

	l := Intersections(i1, i2)

	if len(l) != 2 {
		t.Errorf("List length incorrect expected %d received %d", 2, len(l))
	}

	if !data.FloatEqual(l[0].T, 1) {
		t.Errorf("Intersection 1 incorrect expected %f received %f", 1.0, l[0].T)
	}

	if !data.FloatEqual(l[1].T, 2) {
		t.Errorf("Intersection 2 incorrect expected %f received %f", 2.0, l[1].T)
	}
}

func TestHit(t *testing.T) {
	s := Sphere()
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
		if tc.expected.Object != hit.Object || tc.expected.T != hit.T {
			t.Errorf("Hit does not match expected %+v received %+v", tc.expected, hit)
		}
	}
}

func TestPrepareComputations(t *testing.T) {
	defaultSphere := Sphere()
	tests := []struct {
		r        data.RayType
		o        Shape
		i        IntersectionType
		expected Computations
	}{
		{
			r: data.Ray(data.Point(0, 0, -5), data.Vector(0, 0, 1)),
			o: defaultSphere,
			i: Intersection(4, defaultSphere),
			expected: Computations{
				T:       4,
				Object:  defaultSphere,
				Point:   data.Point(0, 0, -1),
				EyeV:    data.Vector(0, 0, -1),
				NormalV: data.Vector(0, 0, -1),
			},
		},
		{
			r: data.Ray(data.Point(0, 0, 0), data.Vector(0, 0, 1)),
			o: defaultSphere,
			i: Intersection(1, defaultSphere),
			expected: Computations{
				T:       1,
				Object:  defaultSphere,
				Point:   data.Point(0, 0, 1),
				EyeV:    data.Vector(0, 0, -1),
				NormalV: data.Vector(0, 0, -1),
				Inside:  true,
			},
		},
	}

	for _, tc := range tests {
		comp := tc.i.PrepareComputations(tc.r)
		if comp.T != tc.expected.T {
			t.Errorf("T value doesn't match expected: %f received: %f", tc.expected.T, comp.T)
		}

		if comp.Object != tc.expected.Object {
			t.Errorf("Object doesn't match expected: %+v received: %+v", tc.expected.Object, comp.Object)
		}

		if !data.TupleEqual(comp.Point, tc.expected.Point) {
			t.Errorf("data.Point doesn't match expected: %+v received: %+v", tc.expected.Point, comp.Point)
		}

		if !data.TupleEqual(comp.EyeV, tc.expected.EyeV) {
			t.Errorf("EyeV value doesn't match expected: %+v received: %+v", tc.expected.EyeV, comp.EyeV)
		}

		if !data.TupleEqual(comp.NormalV, tc.expected.NormalV) {
			t.Errorf("NormalV value doesn't match expected: %+v received: %+v", tc.expected.NormalV, comp.NormalV)
		}

		if comp.Inside != tc.expected.Inside {
			t.Errorf("Inside value doesn't match expected: %v received: %v", tc.expected.Inside, comp.Inside)
		}
	}
}

func TestOffsetHit(t *testing.T) {
	s := GlassSphere()
	s.Transform = data.Translation(0, 0, 1)
	cases := []struct {
		r      data.RayType
		s      *SphereType
		i      IntersectionType
		offset float64
	}{
		{
			r:      data.Ray(data.Point(0, 0, -5), data.Vector(0, 0, 1)),
			s:      s,
			i:      Intersection(5, s),
			offset: -1 * (data.Epsilon / 2.0),
		},
	}

	for _, tc := range cases {
		c := tc.i.PrepareComputations(tc.r, tc.i)
		if c.OverPoint.Z >= tc.offset {
			t.Errorf("Offset not applied expected <: %v received: %v", tc.offset, c.OverPoint.Z)
		}

		if c.Point.Z <= c.OverPoint.Z {
			t.Errorf("Offset not applied Point <: %v OverPoint: %v", c.Point.Z, c.OverPoint.Z)
		}

		if c.UnderPoint.Z <= -1*tc.offset {
			t.Errorf("Offset not applied expected <: %v received: %v", -1*tc.offset, c.UnderPoint.Z)
		}

		if c.Point.Z >= c.UnderPoint.Z {
			t.Errorf("Offset not applied Point <: %v UnderPoint: %v", c.Point.Z, c.UnderPoint.Z)
		}
	}
}

func TestReflectV(t *testing.T) {
	plane := Plane()
	cases := []struct {
		r        data.RayType
		s        Shape
		i        IntersectionType
		expected data.Tuple
	}{
		{
			r:        data.Ray(data.Point(0, 1, -1), data.Vector(0, math.Sqrt(2)/2*-1, math.Sqrt(2)/2)),
			s:        plane,
			i:        Intersection(math.Sqrt(2), plane),
			expected: data.Vector(0, math.Sqrt(2)/2, math.Sqrt(2)/2),
		},
	}

	for _, tc := range cases {
		c := tc.i.PrepareComputations(tc.r)
		if !data.TupleEqual(tc.expected, c.ReflectV) {
			t.Errorf("Bad reflect vector expected: %v received: %v", tc.expected, c.ReflectV)
		}
	}
}

func TestReraction(t *testing.T) {
	a := GlassSphere()
	a.SetTransform(data.Scaling(2, 2, 2))

	b := GlassSphere()
	b.SetTransform(data.Translation(0, 0, -0.25))
	b.Material.RefractiveIndex = 2.0

	c := GlassSphere()
	c.SetTransform(data.Translation(0, 0, 0.25))
	c.Material.RefractiveIndex = 2.5

	ray := data.Ray(data.Point(0, 0, -4), data.Vector(0, 0, 1))

	xs := Intersections(
		Intersection(2, a),
		Intersection(2.75, b),
		Intersection(3.25, c),
		Intersection(4.75, b),
		Intersection(5.25, c),
		Intersection(6, a),
	)

	results := []Computations{
		{N1: 1.0, N2: 1.5},
		{N1: 1.5, N2: 2.0},
		{N1: 2.0, N2: 2.5},
		{N1: 2.5, N2: 2.5},
		{N1: 2.5, N2: 1.5},
		{N1: 1.5, N2: 1.0},
	}

	for i, x := range xs {
		comps := x.PrepareComputations(ray, xs...)

		if !data.FloatEqual(results[i].N1, comps.N1) {
			t.Errorf("index %d n1 mismatch expected %f received %f", i, results[i].N1, comps.N1)
		}

		if !data.FloatEqual(results[i].N2, comps.N2) {
			t.Errorf("index %d n2 mismatch expected %f received %f", i, results[i].N2, comps.N2)
		}
	}
}

func TestSchlick(t *testing.T) {
	s := GlassSphere()
	cases := []struct {
		s        Shape
		r        data.RayType
		xs       IntersectionList
		i        int
		expected float64
	}{
		{
			s: s,
			r: data.Ray(data.Point(0, 0, math.Sqrt(2)/2), data.Vector(0, 1, 0)),
			xs: IntersectionList{
				{T: -math.Sqrt(2) / 2, Object: s}, {T: math.Sqrt(2) / 2, Object: s},
			},
			i:        1,
			expected: 1.0,
		},
		{
			s: s,
			r: data.Ray(data.Point(0, 0, 0), data.Vector(0, 1, 0)),
			xs: IntersectionList{
				{T: -1, Object: s}, {T: 1, Object: s},
			},
			i:        1,
			expected: 0.04,
		},
		{
			s: s,
			r: data.Ray(data.Point(0, 0.99, -2), data.Vector(0, 0, 1)),
			xs: IntersectionList{
				{T: 1.8589, Object: s},
			},
			i:        0,
			expected: 0.48873,
		},
	}

	for _, tc := range cases {
		comps := tc.xs[tc.i].PrepareComputations(tc.r, tc.xs...)
		reflectance := comps.Schlick()

		if !data.FloatEqual(reflectance, tc.expected) {
			t.Errorf("Reflectance mismatch expected: %f received %f", tc.expected, reflectance)
		}
	}
}
