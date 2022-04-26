package world

import (
	"math"

	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
	"github.com/dannyroes/raytrace/shape"
)

type WorldType struct {
	Objects []shape.Shape
	Lights  []Light
}

func World() WorldType {
	return WorldType{}
}

func DefaultWorld() WorldType {
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

	return WorldType{
		Objects: []shape.Shape{s1, s2},
		Lights:  []Light{l},
	}
}

func (w WorldType) Intersect(r data.RayType) shape.IntersectionList {
	list := shape.IntersectionList{}
	for _, obj := range w.Objects {
		l := shape.Intersects(obj, r)
		list = append(list, l...)
	}
	return list.Sort()
}

func (w WorldType) ShadeHit(c shape.Computations, remain int) material.ColourTuple {
	surface := material.Colour(0, 0, 0)

	for i, l := range w.Lights {
		surface = surface.Add(Lighting(
			c.Object.GetMaterial(),
			c.Object,
			l,
			c.OverPoint,
			c.EyeV,
			c.NormalV,
			w.IsShadowed(c.OverPoint, i),
		))
	}

	// surface := Lighting(
	// 	c.Object.GetMaterial(),
	// 	c.Object,
	// 	w.Light,
	// 	c.OverPoint,
	// 	c.EyeV,
	// 	c.NormalV,
	// 	w.IsShadowed(c.OverPoint),
	// )

	reflect := w.ReflectedColour(c, remain)
	refract := w.RefractedColour(c, remain)

	material := c.Object.GetMaterial()
	if material.Reflective > 0 && material.Transparency > 0 {
		reflectance := c.Schlick()
		return surface.Add(reflect.Mul(reflectance)).Add(refract.Mul(1 - reflectance))
	}

	return surface.Add(reflect).Add(refract)
}

func (w WorldType) IsShadowed(p data.Tuple, lightIndex int) bool {
	v := w.Lights[lightIndex].Position.Sub(p)
	distance := v.Magnitude()
	direction := v.Normalize()
	r := data.Ray(p, direction)
	intersections := w.Intersect(r)

	h := intersections.Hit()
	if h.Object != nil && h.T < distance && h.Object.CastsShadow() {
		return true
	}

	return false
}

func (w WorldType) ColourAt(r data.RayType, remain int) material.ColourTuple {
	i := w.Intersect(r)
	h := i.Hit()

	if h.T == -1 {
		return material.Colour(0, 0, 0)
	}

	c := h.PrepareComputations(r, i...)
	return w.ShadeHit(c, remain)
}

func (w WorldType) ReflectedColour(c shape.Computations, remain int) material.ColourTuple {
	if c.Object.GetMaterial().Reflective == 0 || remain == 0 {
		return material.Black
	}

	ray := data.Ray(c.OverPoint, c.ReflectV)
	colour := w.ColourAt(ray, remain-1)

	return colour.Mul(c.Object.GetMaterial().Reflective)
}

func (w WorldType) RefractedColour(c shape.Computations, remain int) material.ColourTuple {
	if c.Object.GetMaterial().Transparency == 0 || remain == 0 {
		return material.Black
	}

	nRatio := c.N1 / c.N2
	cosi := data.Dot(c.EyeV, c.NormalV)

	sin2t := math.Pow(nRatio, 2) * (1 - math.Pow(cosi, 2))

	if sin2t > 1 {
		return material.Black
	}

	cost := math.Sqrt(1.0 - sin2t)
	dir := c.NormalV.Mul((nRatio * cosi) - cost).Sub(c.EyeV.Mul(nRatio))

	refractRay := data.Ray(c.UnderPoint, dir)
	colour := w.ColourAt(refractRay, remain-1).Mul(c.Object.GetMaterial().Transparency)

	return colour
}
