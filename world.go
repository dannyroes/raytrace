package main

import (
	"math"
)

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
	refract := w.RefractedColour(c, remain)

	material := c.Object.GetMaterial()
	if material.Reflective > 0 && material.Transparency > 0 {
		reflectance := c.Schlick()
		return surface.Add(reflect.Mul(reflectance)).Add(refract.Mul(1 - reflectance))
	}

	return surface.Add(reflect).Add(refract)
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

	c := h.PrepareComputations(r, i...)
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

func (w WorldType) RefractedColour(c Computations, remain int) ColourTuple {
	if c.Object.GetMaterial().Transparency == 0 || remain == 0 {
		return Black
	}

	nRatio := c.N1 / c.N2
	cosi := Dot(c.EyeV, c.NormalV)

	sin2t := math.Pow(nRatio, 2) * (1 - math.Pow(cosi, 2))

	if sin2t > 1 {
		return Black
	}

	cost := math.Sqrt(1.0 - sin2t)
	dir := c.NormalV.Mul((nRatio * cosi) - cost).Sub(c.EyeV.Mul(nRatio))

	refractRay := Ray(c.UnderPoint, dir)
	colour := w.ColourAt(refractRay, remain-1).Mul(c.Object.GetMaterial().Transparency)

	return colour
}
