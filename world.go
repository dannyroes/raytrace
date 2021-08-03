package main

type WorldType struct {
	Objects []Shape
	Light   Light
}

func World() WorldType {
	return WorldType{}
}

func DefaultWorld() WorldType {
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

	return WorldType{
		Objects: []Shape{s1, s2},
		Light:   l,
	}
}

func (w WorldType) Intersect(r RayType) IntersectionList {
	list := IntersectionList{}
	for _, obj := range w.Objects {
		l := Intersects(obj, r)
		list = append(list, l...)
	}
	return list.Sort()
}

func (w WorldType) ShadeHit(c Computations, remain int) ColourTuple {
	surface := Lighting(
		c.Object.GetMaterial(),
		c.Object,
		w.Light,
		c.OverPoint,
		c.EyeV,
		c.NormalV,
		w.IsShadowed(c.OverPoint),
	)

	reflect := w.ReflectedColour(c, remain)

	return surface.Add(reflect)
}

func (w WorldType) IsShadowed(p Tuple) bool {
	v := w.Light.Position.Sub(p)
	distance := v.Magnitude()
	direction := v.Normalize()
	r := Ray(p, direction)
	intersections := w.Intersect(r)

	h := intersections.Hit()
	if h.Object != nil && h.T < distance {
		return true
	}

	return false
}

func (w WorldType) ColourAt(r RayType, remain int) ColourTuple {
	i := w.Intersect(r)
	h := i.Hit()

	if h.T == -1 {
		return Colour(0, 0, 0)
	}

	c := h.PrepareComputations(r)
	return w.ShadeHit(c, remain)
}

func (w WorldType) ReflectedColour(c Computations, remain int) ColourTuple {
	if c.Object.GetMaterial().Reflective == 0 || remain == 0 {
		return Black
	}

	ray := Ray(c.OverPoint, c.ReflectV)
	colour := w.ColourAt(ray, remain-1)

	return colour.Mul(c.Object.GetMaterial().Reflective)
}
