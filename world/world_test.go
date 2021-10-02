package world

import (
	"math"
	"testing"

	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
	"github.com/dannyroes/raytrace/shape"
)

func TestWorld(t *testing.T) {
	w := World()

	if len(w.Objects) != 0 {
		t.Error("Empty world contains objects")
	}

	if len(w.Lights) != 0 {
		t.Error("Empty world contains light")
	}

	l := PointLight(data.Point(-10, 10, -10), material.Colour(1, 1, 1))

	mat := material.Material()
	mat.Colour = material.Colour(0.8, 1.0, 0.6)
	mat.Diffuse = 0.7
	mat.Specular = 0.2
	s1 := shape.Sphere()
	s1.Material = mat

	scale := data.IdentityMatrix().Scale(0.5, 0.5, 0.5)
	s2 := shape.Sphere()
	s2.Transform = scale

	w = DefaultWorld()

	if w.Lights[0] != l {
		t.Errorf("Default world contains wrong light expected: %+v received: %+v", l, w.Lights[0])
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
	r := data.Ray(data.Point(0, 0, -5), data.Vector(0, 0, 1))

	xs := w.Intersect(r)

	expected := []float64{4, 4.5, 5.5, 6}

	if len(xs) != 4 {
		t.Errorf("Wrong number of intersections expected: %d received: %d", 4, len(xs))
	}

	for i, v := range xs {
		if !data.FloatEqual(expected[i], v.T) {
			t.Errorf("Wrong intersection T expected %f received %f", expected[i], v.T)
		}
	}

}

func TestShadeHit(t *testing.T) {
	w := DefaultWorld()
	s2 := shape.Sphere()
	s2.Transform = data.Translation(0, 0, 10)
	tests := []struct {
		w        WorldType
		l        Light
		r        data.RayType
		i        shape.IntersectionType
		expected material.ColourTuple
	}{
		{
			w:        w,
			l:        w.Lights[0],
			r:        data.Ray(data.Point(0, 0, -5), data.Vector(0, 0, 1)),
			i:        shape.Intersection(4, w.Objects[0]),
			expected: material.Colour(0.38066, 0.47583, 0.2855),
		},
		{
			w:        w,
			l:        PointLight(data.Point(0, 0.25, 0), material.Colour(1, 1, 1)),
			r:        data.Ray(data.Point(0, 0, 0), data.Vector(0, 0, 1)),
			i:        shape.Intersection(0.5, w.Objects[1]),
			expected: material.Colour(0.90498, 0.90498, 0.90498),
		},
		{
			w: func() WorldType {
				w := World()

				s1 := shape.Sphere()
				w.Objects = []shape.Shape{s1, s2}

				return w
			}(),
			l:        PointLight(data.Point(0, 0, -10), material.Colour(1, 1, 1)),
			r:        data.Ray(data.Point(0, 0, 5), data.Vector(0, 0, 1)),
			i:        shape.Intersection(4, s2),
			expected: material.Colour(0.1, 0.1, 0.1),
		},
	}

	for _, tc := range tests {
		tc.w.Lights = []Light{tc.l}
		comp := tc.i.PrepareComputations(tc.r)
		result := tc.w.ShadeHit(comp, 1)

		if !material.ColourEqual(result, tc.expected) {
			t.Errorf("Colour mismatch expected: %+v received %+v", tc.expected, result)
		}
	}
}

func TestColourAt(t *testing.T) {
	tests := []struct {
		w        WorldType
		r        data.RayType
		expected material.ColourTuple
	}{
		{
			w:        DefaultWorld(),
			r:        data.Ray(data.Point(0, 0, -5), data.Vector(0, 1, 0)),
			expected: material.Colour(0, 0, 0),
		},
		{
			w:        DefaultWorld(),
			r:        data.Ray(data.Point(0, 0, -5), data.Vector(0, 0, 1)),
			expected: material.Colour(0.38066, 0.47583, 0.2855),
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
			r:        data.Ray(data.Point(0, 0, 0.75), data.Vector(0, 0, -1)),
			expected: DefaultWorld().Objects[1].GetMaterial().Colour,
		},
	}

	for _, tc := range tests {
		result := tc.w.ColourAt(tc.r, 1)

		if !material.ColourEqual(result, tc.expected) {
			t.Errorf("Colour mismatch expected: %+v received %+v", tc.expected, result)
		}
	}
}

func TestIsShadowed(t *testing.T) {
	tests := []struct {
		w        WorldType
		p        data.Tuple
		expected bool
	}{
		{
			w:        DefaultWorld(),
			p:        data.Point(0, 10, 0),
			expected: false,
		},
		{
			w:        DefaultWorld(),
			p:        data.Point(10, -10, 10),
			expected: true,
		},
		{
			w:        DefaultWorld(),
			p:        data.Point(-20, 20, -20),
			expected: false,
		},
		{
			w:        DefaultWorld(),
			p:        data.Point(-2, 2, -2),
			expected: false,
		},
	}

	for _, tc := range tests {
		result := tc.w.IsShadowed(tc.p, 0)

		if tc.expected != result {
			t.Errorf("Shadow mismatch expected: %+v received %+v", tc.expected, result)
		}
	}
}

func containsObject(w WorldType, obj shape.Shape) bool {
	for _, o := range w.Objects {
		if obj.GetTransform().Equals(o.GetTransform()) && obj.GetMaterial().Equals(o.GetMaterial()) {
			return true
		}
	}
	return false
}

func TestMaterialPattern(t *testing.T) {
	m := material.Material()
	m.Ambient = 1
	m.Diffuse = 0
	m.Specular = 0
	m.Pattern = material.StripePattern(material.White, material.Black)

	eyeV := data.Vector(0, 0, -1)
	normalV := data.Vector(0, 0, -1)
	light := PointLight(data.Point(0, 0, -1), material.White)

	c1 := Lighting(m, shape.Sphere(), light, data.Point(0.9, 0, 0), eyeV, normalV, false)
	if !material.ColourEqual(c1, material.White) {
		t.Errorf("C1 mismatch expected %+v received %+v", material.White, c1)
	}

	c2 := Lighting(m, shape.Sphere(), light, data.Point(1.1, 0, 0), eyeV, normalV, false)
	if !material.ColourEqual(c2, material.Black) {
		t.Errorf("C2 mismatch expected %+v received %+v", material.Black, c2)
	}
}

func TestReflectedColour(t *testing.T) {
	cases := []struct {
		r         data.RayType
		worldFunc func() (WorldType, shape.IntersectionType)
		reflect   material.ColourTuple
		remain    int
		shade     material.ColourTuple
	}{
		{
			r: data.Ray(data.Point(0, 0, 0), data.Vector(0, 0, 1)),
			worldFunc: func() (WorldType, shape.IntersectionType) {
				w := DefaultWorld()

				mat := w.Objects[1].GetMaterial()
				mat.Ambient = 1
				w.Objects[1].SetMaterial(mat)

				i := shape.Intersection(1, w.Objects[1])
				return w, i
			},
			reflect: material.Black,
			remain:  1,
			shade:   material.White,
		},
		{
			r: data.Ray(data.Point(0, 0, -3), data.Vector(0, -1*(math.Sqrt(2)/2), math.Sqrt(2)/2)),
			worldFunc: func() (WorldType, shape.IntersectionType) {
				w := DefaultWorld()

				s := shape.Plane()

				mat := s.GetMaterial()
				mat.Reflective = 0.5
				s.SetMaterial(mat)
				s.SetTransform(data.Translation(0, -1, 0))

				w.Objects = append(w.Objects, s)

				i := shape.Intersection(math.Sqrt(2), s)
				return w, i
			},
			reflect: material.Black,
			shade:   material.Colour(0.87675, 0.92434, 0.82918),
		},
	}

	for _, tc := range cases {
		w, i := tc.worldFunc()
		c := i.PrepareComputations(tc.r)

		colour := w.ReflectedColour(c, tc.remain)
		if !material.ColourEqual(tc.reflect, colour) {
			t.Errorf("Bad reflect colour expected: %v received: %v", tc.reflect, colour)
		}

		colour = w.ShadeHit(c, 1)
		if !material.ColourEqual(tc.shade, colour) {
			t.Errorf("Bad shade colour expected: %v received: %v", tc.shade, colour)
		}
	}
}

func TestInfiniteRecursion(t *testing.T) {
	w := World()
	w.Lights = []Light{PointLight(data.Point(0, 0, 0), material.Colour(1, 1, 1))}

	lower := shape.Plane()
	m := material.Material()
	m.Reflective = 1
	lower.SetMaterial(m)
	lower.Transform = data.Translation(0, -1, 0)

	upper := shape.Plane()
	upper.SetMaterial(m)
	upper.Transform = data.Translation(0, 1, 0)

	w.Objects = []shape.Shape{lower, upper}

	r := data.Ray(data.Point(0, 0, 0), data.Vector(0, 1, 0))
	_ = w.ColourAt(r, MaxReflect)
}

func TestRefractedColour(t *testing.T) {
	cases := []struct {
		r         data.RayType
		worldFunc func() (WorldType, shape.IntersectionList)
		refract   material.ColourTuple
		remain    int
		hitIndex  int
	}{
		{
			r: data.Ray(data.Point(0, 0, 0), data.Vector(0, 0, 1)),
			worldFunc: func() (WorldType, shape.IntersectionList) {
				w := DefaultWorld()

				l := shape.Intersections(shape.Intersection(4, w.Objects[0]), shape.Intersection(6, w.Objects[0]))
				return w, l
			},
			refract: material.Black,
			remain:  5,
		},
		{
			r: data.Ray(data.Point(0, 0, 0), data.Vector(0, 0, 1)),
			worldFunc: func() (WorldType, shape.IntersectionList) {
				w := DefaultWorld()

				m := w.Objects[0].GetMaterial()
				m.Transparency = 1.0
				m.RefractiveIndex = 1.5
				w.Objects[0].SetMaterial(m)

				l := shape.Intersections(shape.Intersection(4, w.Objects[0]), shape.Intersection(6, w.Objects[0]))
				return w, l
			},
			refract: material.Black,
			remain:  0,
		},
		{
			r: data.Ray(data.Point(0, 0, math.Sqrt2/2), data.Vector(0, 1, 0)),
			worldFunc: func() (WorldType, shape.IntersectionList) {
				w := DefaultWorld()

				m := w.Objects[0].GetMaterial()
				m.Transparency = 1.0
				m.RefractiveIndex = 1.5
				w.Objects[0].SetMaterial(m)

				l := shape.Intersections(shape.Intersection(-math.Sqrt2/2, w.Objects[0]), shape.Intersection(math.Sqrt2/2, w.Objects[0]))
				return w, l
			},
			refract: material.Black,
			remain:  0,
		},
		{
			r: data.Ray(data.Point(0, 0, 0.1), data.Vector(0, 1, 0)),
			worldFunc: func() (WorldType, shape.IntersectionList) {
				w := DefaultWorld()

				m := w.Objects[0].GetMaterial()
				m.Ambient = 1.0
				m.Pattern = material.TestPattern()
				w.Objects[0].SetMaterial(m)

				m = w.Objects[1].GetMaterial()
				m.Transparency = 1.0
				m.RefractiveIndex = 1.5
				w.Objects[1].SetMaterial(m)

				l := shape.Intersections(
					shape.Intersection(-0.9899, w.Objects[0]),
					shape.Intersection(-0.4899, w.Objects[1]),
					shape.Intersection(0.4899, w.Objects[1]),
					shape.Intersection(0.9899, w.Objects[0]),
				)
				return w, l
			},
			refract:  material.Colour(0, 0.99887, 0.04721),
			remain:   5,
			hitIndex: 2,
		},
	}

	for _, tc := range cases {
		w, l := tc.worldFunc()
		c := l[tc.hitIndex].PrepareComputations(tc.r, l...)

		colour := w.RefractedColour(c, tc.remain)
		if !material.ColourEqual(tc.refract, colour) {
			t.Errorf("Bad refract colour expected: %v received: %v", tc.refract, colour)
		}
	}
}

func TestShadeHitRefraction(t *testing.T) {
	w := DefaultWorld()
	floor := shape.Plane()
	floor.SetTransform(data.Translation(0, -1, 0))
	floor.Material.Transparency = 0.5
	floor.Material.RefractiveIndex = 1.5

	ball := shape.Sphere()
	ball.Material.Colour = material.Colour(1, 0, 0)
	ball.Material.Ambient = 0.5
	ball.SetTransform(data.Translation(0, -3.5, -0.5))

	w.Objects = append(w.Objects, floor, ball)

	r := data.Ray(data.Point(0, 0, -3), data.Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	xs := shape.IntersectionList{shape.Intersection(math.Sqrt(2), floor)}

	comps := xs.Hit().PrepareComputations(r, xs...)
	colour := w.ShadeHit(comps, 5)

	if !material.ColourEqual(colour, material.Colour(0.93642, 0.68642, 0.68642)) {
		t.Errorf("Colour mismatch expected %v received %v", material.Colour(0.93642, 0.68642, 0.68642), colour)
	}
}

func TestShadeHitTransparentReflection(t *testing.T) {
	w := DefaultWorld()
	floor := shape.Plane()
	floor.SetTransform(data.Translation(0, -1, 0))
	floor.Material.Reflective = 0.5
	floor.Material.Transparency = 0.5
	floor.Material.RefractiveIndex = 1.5

	ball := shape.Sphere()
	ball.Material.Colour = material.Colour(1, 0, 0)
	ball.Material.Ambient = 0.5
	ball.SetTransform(data.Translation(0, -3.5, -0.5))

	w.Objects = append(w.Objects, floor, ball)

	r := data.Ray(data.Point(0, 0, -3), data.Vector(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	xs := shape.IntersectionList{shape.Intersection(math.Sqrt(2), floor)}

	comps := xs.Hit().PrepareComputations(r, xs...)
	colour := w.ShadeHit(comps, 5)

	if !material.ColourEqual(colour, material.Colour(0.93391, 0.69643, 0.69243)) {
		t.Errorf("Colour mismatch expected %v received %v", material.Colour(0.93642, 0.68642, 0.68642), colour)
	}
}
