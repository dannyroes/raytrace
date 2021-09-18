package main

import (
	"math"
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
	s1 := Sphere()
	s1.Material = mat

	scale := IdentityMatrix().Scale(0.5, 0.5, 0.5)
	s2 := Sphere()
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
	s2 := Sphere()
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

				s1 := Sphere()
				w.Objects = []Shape{s1, s2}

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
		result := tc.w.ShadeHit(comp, 1)

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
		result := tc.w.ColourAt(tc.r, 1)

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

func containsObject(w WorldType, obj Shape) bool {
	for _, o := range w.Objects {
		if obj.GetTransform().Equals(o.GetTransform()) && obj.GetMaterial().Equals(o.GetMaterial()) {
			return true
		}
	}
	return false
}

func TestReflectedColour(t *testing.T) {
	cases := []struct {
		r         RayType
		worldFunc func() (WorldType, IntersectionType)
		reflect   ColourTuple
		remain    int
		shade     ColourTuple
	}{
		{
			r: Ray(Point(0, 0, 0), Vector(0, 0, 1)),
			worldFunc: func() (WorldType, IntersectionType) {
				w := DefaultWorld()

				mat := w.Objects[1].GetMaterial()
				mat.Ambient = 1
				w.Objects[1].SetMaterial(mat)

				i := Intersection(1, w.Objects[1])
				return w, i
			},
			reflect: Black,
			remain:  1,
			shade:   White,
		},
		{
			r: Ray(Point(0, 0, -3), Vector(0, -1*(math.Sqrt(2)/2), math.Sqrt(2)/2)),
			worldFunc: func() (WorldType, IntersectionType) {
				w := DefaultWorld()

				s := Plane()

				mat := s.GetMaterial()
				mat.Reflective = 0.5
				s.SetMaterial(mat)
				s.SetTransform(Translation(0, -1, 0))

				w.Objects = append(w.Objects, s)

				i := Intersection(math.Sqrt(2), s)
				return w, i
			},
			reflect: Black,
			shade:   Colour(0.87675, 0.92434, 0.82918),
		},
	}

	for _, tc := range cases {
		w, i := tc.worldFunc()
		c := i.PrepareComputations(tc.r)

		colour := w.ReflectedColour(c, tc.remain)
		if !ColourEqual(tc.reflect, colour) {
			t.Errorf("Bad reflect colour expected: %v received: %v", tc.reflect, colour)
		}

		colour = w.ShadeHit(c, 1)
		if !ColourEqual(tc.shade, colour) {
			t.Errorf("Bad shade colour expected: %v received: %v", tc.shade, colour)
		}
	}
}

func TestInfiniteRecursion(t *testing.T) {
	w := World()
	w.Light = PointLight(Point(0, 0, 0), Colour(1, 1, 1))

	lower := Plane()
	m := Material()
	m.Reflective = 1
	lower.SetMaterial(m)
	lower.Transform = Translation(0, -1, 0)

	upper := Plane()
	upper.SetMaterial(m)
	upper.Transform = Translation(0, 1, 0)

	w.Objects = []Shape{lower, upper}

	r := Ray(Point(0, 0, 0), Vector(0, 1, 0))
	_ = w.ColourAt(r, MaxReflect)
}

func TestRefractedColour(t *testing.T) {
	cases := []struct {
		r         RayType
		worldFunc func() (WorldType, IntersectionList)
		refract   ColourTuple
		remain    int
		hitIndex  int
	}{
		{
			r: Ray(Point(0, 0, 0), Vector(0, 0, 1)),
			worldFunc: func() (WorldType, IntersectionList) {
				w := DefaultWorld()

				l := Intersections(Intersection(4, w.Objects[0]), Intersection(6, w.Objects[0]))
				return w, l
			},
			refract: Black,
			remain:  5,
		},
		{
			r: Ray(Point(0, 0, 0), Vector(0, 0, 1)),
			worldFunc: func() (WorldType, IntersectionList) {
				w := DefaultWorld()

				m := w.Objects[0].GetMaterial()
				m.Transparency = 1.0
				m.RefractiveIndex = 1.5
				w.Objects[0].SetMaterial(m)

				l := Intersections(Intersection(4, w.Objects[0]), Intersection(6, w.Objects[0]))
				return w, l
			},
			refract: Black,
			remain:  0,
		},
		{
			r: Ray(Point(0, 0, math.Sqrt2/2), Vector(0, 1, 0)),
			worldFunc: func() (WorldType, IntersectionList) {
				w := DefaultWorld()

				m := w.Objects[0].GetMaterial()
				m.Transparency = 1.0
				m.RefractiveIndex = 1.5
				w.Objects[0].SetMaterial(m)

				l := Intersections(Intersection(-math.Sqrt2/2, w.Objects[0]), Intersection(math.Sqrt2/2, w.Objects[0]))
				return w, l
			},
			refract: Black,
			remain:  0,
		},
		{
			r: Ray(Point(0, 0, 0.1), Vector(0, 1, 0)),
			worldFunc: func() (WorldType, IntersectionList) {
				w := DefaultWorld()

				m := w.Objects[0].GetMaterial()
				m.Ambient = 1.0
				m.Pattern = TestPattern()
				w.Objects[0].SetMaterial(m)

				m = w.Objects[1].GetMaterial()
				m.Transparency = 1.0
				m.RefractiveIndex = 1.5
				w.Objects[1].SetMaterial(m)

				l := Intersections(
					Intersection(-0.9899, w.Objects[0]),
					Intersection(-0.4899, w.Objects[1]),
					Intersection(0.4899, w.Objects[1]),
					Intersection(0.9899, w.Objects[0]),
				)
				return w, l
			},
			refract:  Colour(0, 0.99887, 0.04721),
			remain:   5,
			hitIndex: 2,
		},
	}

	for _, tc := range cases {
		w, l := tc.worldFunc()
		c := l[tc.hitIndex].PrepareComputations(tc.r, l...)

		colour := w.RefractedColour(c, tc.remain)
		if !ColourEqual(tc.refract, colour) {
			t.Errorf("Bad refract colour expected: %v received: %v", tc.refract, colour)
		}
	}
}

func TestShadeHitRefraction(t *testing.T) {
	w := DefaultWorld()
	floor := Plane()
	floor.SetTransform(Translation(0, -1, 0))
	floor.Material.Transparency = 0.5
	floor.Material.RefractiveIndex = 1.5

	ball := Sphere()
	ball.Material.Colour = Colour(1, 0, 0)
	ball.Material.Ambient = 0.5
	ball.SetTransform(Translation(0, -3.5, -0.5))

	w.Objects = append(w.Objects, floor, ball)

	r := Ray(Point(0, 0, -3), Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	xs := IntersectionList{Intersection(math.Sqrt(2), floor)}

	comps := xs.Hit().PrepareComputations(r, xs...)
	colour := w.ShadeHit(comps, 5)

	if !ColourEqual(colour, Colour(0.93642, 0.68642, 0.68642)) {
		t.Errorf("Colour mismatch expected %v received %v", Colour(0.93642, 0.68642, 0.68642), colour)
	}
}

func TestShadeHitTransparentReflection(t *testing.T) {
	w := DefaultWorld()
	floor := Plane()
	floor.SetTransform(Translation(0, -1, 0))
	floor.Material.Reflective = 0.5
	floor.Material.Transparency = 0.5
	floor.Material.RefractiveIndex = 1.5

	ball := Sphere()
	ball.Material.Colour = Colour(1, 0, 0)
	ball.Material.Ambient = 0.5
	ball.SetTransform(Translation(0, -3.5, -0.5))

	w.Objects = append(w.Objects, floor, ball)

	r := Ray(Point(0, 0, -3), Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	xs := IntersectionList{Intersection(math.Sqrt(2), floor)}

	comps := xs.Hit().PrepareComputations(r, xs...)
	colour := w.ShadeHit(comps, 5)

	if !ColourEqual(colour, Colour(0.93391, 0.69643, 0.69243)) {
		t.Errorf("Colour mismatch expected %v received %v", Colour(0.93642, 0.68642, 0.68642), colour)
	}
}
