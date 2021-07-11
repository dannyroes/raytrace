package main

import (
	"testing"
)

func TestWorld(t *testing.T) {
	w := World()

	if len(w.Objects) != 0 {
		t.Error("Empty world contains objects")
	}

	emptyLight := Light{}
	if w.Light != emptyLight {
		t.Error("Empty world contains light")
	}

	l := PointLight(Point(-10, 10, -10), Colour(1, 1, 1))

	mat := Material()
	mat.Colour = Colour(0.8, 1.0, 0.6)
	mat.Diffuse = 0.7
	mat.Specular = 0.2
	s1 := Sphere(1)
	s1.Material = mat

	scale := IdentityMatrix().Scale(0.5, 0.5, 0.5)
	s2 := Sphere(2)
	s2.Transform = scale

	w = DefaultWorld()

	if w.Light != l {
		t.Errorf("Default world contains wrong light expected: %+v received: %+v", l, w.Light)
	}

	if !containsObject(w, s1) {
		t.Errorf("Default world missing object expected: %+v", s1)
	}
	if !containsObject(w, s2) {
		t.Errorf("Default world missing object expected: %+v", s2)
	}
}

func TestIntersectWorld(t *testing.T) {
	w := DefaultWorld()
	r := Ray(Point(0, 0, -5), Vector(0, 0, 1))

	xs := w.Intersect(r)

	expected := []float64{4, 4.5, 5.5, 6}

	if len(xs) != 4 {
		t.Errorf("Wrong number of intersections expected: %d received: %d", 4, len(xs))
	}

	for i, v := range xs {
		if !FloatEqual(expected[i], v.T) {
			t.Errorf("Wrong intersection T expected %f received %f", expected[i], v.T)
		}
	}

}

func TestShadeHit(t *testing.T) {
	w := DefaultWorld()
	s2 := Sphere(2)
	s2.Transform = Translation(0, 0, 10)
	tests := []struct {
		w        WorldType
		l        Light
		r        RayType
		i        IntersectionType
		expected ColourTuple
	}{
		{
			w:        w,
			l:        w.Light,
			r:        Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			i:        Intersection(4, w.Objects[0]),
			expected: Colour(0.38066, 0.47583, 0.2855),
		},
		{
			w:        w,
			l:        PointLight(Point(0, 0.25, 0), Colour(1, 1, 1)),
			r:        Ray(Point(0, 0, 0), Vector(0, 0, 1)),
			i:        Intersection(0.5, w.Objects[1]),
			expected: Colour(0.90498, 0.90498, 0.90498),
		},
		{
			w: func() WorldType {
				w := World()

				s1 := Sphere(1)
				w.Objects = []Object{s1, s2}

				return w
			}(),
			l:        PointLight(Point(0, 0, -10), Colour(1, 1, 1)),
			r:        Ray(Point(0, 0, 5), Vector(0, 0, 1)),
			i:        Intersection(4, s2),
			expected: Colour(0.1, 0.1, 0.1),
		},
	}

	for _, tc := range tests {
		tc.w.Light = tc.l
		comp := tc.i.PrepareComputations(tc.r)
		result := tc.w.ShadeHit(comp)

		if !ColourEqual(result, tc.expected) {
			t.Errorf("Colour mismatch expected: %+v received %+v", tc.expected, result)
		}
	}
}

func TestColourAt(t *testing.T) {
	tests := []struct {
		w        WorldType
		r        RayType
		expected ColourTuple
	}{
		{
			w:        DefaultWorld(),
			r:        Ray(Point(0, 0, -5), Vector(0, 1, 0)),
			expected: Colour(0, 0, 0),
		},
		{
			w:        DefaultWorld(),
			r:        Ray(Point(0, 0, -5), Vector(0, 0, 1)),
			expected: Colour(0.38066, 0.47583, 0.2855),
		},
		{
			w: func() WorldType {
				world := DefaultWorld()
				for i := range world.Objects {
					m := world.Objects[i].GetMaterial()
					m.Ambient = 1
					world.Objects[i].SetMaterial(m)
				}

				return world
			}(),
			r:        Ray(Point(0, 0, 0.75), Vector(0, 0, -1)),
			expected: DefaultWorld().Objects[1].GetMaterial().Colour,
		},
	}

	for _, tc := range tests {
		result := tc.w.ColourAt(tc.r)

		if !ColourEqual(result, tc.expected) {
			t.Errorf("Colour mismatch expected: %+v received %+v", tc.expected, result)
		}
	}
}

func TestIsShadowed(t *testing.T) {
	tests := []struct {
		w        WorldType
		p        Tuple
		expected bool
	}{
		{
			w:        DefaultWorld(),
			p:        Point(0, 10, 0),
			expected: false,
		},
		{
			w:        DefaultWorld(),
			p:        Point(10, -10, 10),
			expected: true,
		},
		{
			w:        DefaultWorld(),
			p:        Point(-20, 20, -20),
			expected: false,
		},
		{
			w:        DefaultWorld(),
			p:        Point(-2, 2, -2),
			expected: false,
		},
	}

	for _, tc := range tests {
		result := tc.w.IsShadowed(tc.p)

		if tc.expected != result {
			t.Errorf("Shadow mismatch expected: %+v received %+v", tc.expected, result)
		}
	}
}

func containsObject(w WorldType, obj Object) bool {
	for _, o := range w.Objects {
		if o.GetId() == obj.GetId() {
			return true
		}
	}
	return false
}
