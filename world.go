package main

type WorldType struct {
	Objects []Object
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
	s1 := Sphere(1)
	s1.Material = mat

	scale := IdentityMatrix().Scale(0.5, 0.5, 0.5)
	s2 := Sphere(2)
	s2.Transform = scale

	return WorldType{
		Objects: []Object{s1, s2},
		Light:   l,
	}
}

func (w WorldType) Intersect(r RayType) IntersectionList {
	list := IntersectionList{}
	for _, obj := range w.Objects {
		l := obj.Intersects(r)
		list = append(list, l...)
	}
	return list.Sort()
}

func (w WorldType) ShadeHit(c Computations) ColourTuple {
	return Lighting(
		c.Object.GetMaterial(),
		w.Light,
		c.Point,
		c.EyeV,
		c.NormalV,
	)
}

func (w WorldType) ColourAt(r RayType) ColourTuple {
	i := w.Intersect(r)
	h := i.Hit()

	if h.T == -1 {
		return Colour(0, 0, 0)
	}

	c := h.PrepareComputations(r)
	return w.ShadeHit(c)
}