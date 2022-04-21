package shape

import (
	"math"

	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
)

type CylinderType struct {
	ShapeType
	Minimum float64
	Maximum float64
	Closed  bool
}

func Cylinder() *CylinderType {
	return &CylinderType{
		Minimum: math.Inf(-1),
		Maximum: math.Inf(1),
		ShapeType: ShapeType{
			Material:  material.Material(),
			Transform: data.IdentityMatrix(),
		}}
}

func (cyl *CylinderType) GetMaterial() material.MaterialType {
	return cyl.Material
}

func (cyl *CylinderType) SetMaterial(m material.MaterialType) {
	cyl.Material = m
}

func (cyl *CylinderType) LocalIntersect(r data.RayType) IntersectionList {
	a := math.Pow(r.Direction.X, 2) + math.Pow(r.Direction.Z, 2)

	// if data.FloatEqual(0, a) {
	// 	return Intersections()
	// }

	b := 2*r.Origin.X*r.Direction.X + 2*r.Origin.Z*r.Direction.Z
	c := math.Pow(r.Origin.X, 2) + math.Pow(r.Origin.Z, 2) - 1

	disc := math.Pow(b, 2) - 4*a*c

	if disc < 0 {
		return Intersections()
	}

	t0 := (-b - math.Sqrt(disc)) / (2 * a)
	t1 := (-b + math.Sqrt(disc)) / (2 * a)
	if t0 > t1 {
		t0, t1 = t1, t0
	}

	xs := Intersections()

	y0 := r.Origin.Y + t0*r.Direction.Y
	if cyl.Minimum < y0 && y0 < cyl.Maximum {
		xs = append(xs, Intersection(t0, cyl))
	}

	y1 := r.Origin.Y + t1*r.Direction.Y
	if cyl.Minimum < y1 && y1 < cyl.Maximum {
		xs = append(xs, Intersection(t1, cyl))
	}

	xs = cyl.intersectCaps(r, xs)

	return xs
}

func (cyl *CylinderType) SetTransform(m data.Matrix) {
	cyl.Transform = m
}

func (cyl *CylinderType) GetTransform() data.Matrix {
	return cyl.Transform
}

func (cyl *CylinderType) LocalNormalAt(objectPoint data.Tuple, i IntersectionType) data.Tuple {
	dist := math.Pow(objectPoint.X, 2) + math.Pow(objectPoint.Z, 2)
	if dist < 1 && objectPoint.Y >= cyl.Maximum-data.Epsilon {
		return data.Vector(0, 1, 0)
	}

	if dist < 1 && objectPoint.Y <= cyl.Minimum+data.Epsilon {
		return data.Vector(0, -1, 0)
	}

	return data.Vector(objectPoint.X, 0, objectPoint.Z)
}

func (cyl *CylinderType) GetParent() Shape {
	return cyl.Parent
}

func (cyl *CylinderType) SetParent(p Shape) {
	cyl.Parent = p
}

func (cyl *CylinderType) Bounds() Bounds {
	return Bounds{
		Min: data.Point(-1, cyl.Minimum, -1),
		Max: data.Point(1, cyl.Maximum, 1),
	}
}

func checkCap(r data.RayType, t, y float64) bool {
	x := r.Origin.X + t*r.Direction.X
	z := r.Origin.Z + t*r.Direction.Z

	return math.Pow(x, 2)+math.Pow(z, 2) <= y*y
}

func (cyl *CylinderType) intersectCaps(r data.RayType, xs IntersectionList) IntersectionList {
	if !cyl.Closed || data.FloatEqual(r.Direction.Y, 0) {
		return xs
	}

	t := (cyl.Minimum - r.Origin.Y) / r.Direction.Y
	if checkCap(r, t, 1) {
		xs = append(xs, Intersection(t, cyl))
	}

	t = (cyl.Maximum - r.Origin.Y) / r.Direction.Y
	if checkCap(r, t, 1) {
		xs = append(xs, Intersection(t, cyl))
	}

	return xs
}
