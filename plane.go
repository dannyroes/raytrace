package main

import "math"

type PlaneType struct {
	ShapeType
}

func Plane() *PlaneType {
	return &PlaneType{ShapeType{Transform: IdentityMatrix(), Material: Material()}}
}

func (p *PlaneType) SetMaterial(m MaterialType) {
	p.Material = m
}

func (p *PlaneType) GetMaterial() MaterialType {
	return p.Material
}

func (p *PlaneType) GetTransform() Matrix {
	return p.Transform
}

func (p *PlaneType) SetTransform(m Matrix) {
	p.Transform = m
}

func (p *PlaneType) LocalIntersect(r RayType) IntersectionList {
	if math.Abs(r.Direction.Y) < Epsilon {
		return IntersectionList{}
	}

	t := (-1 * r.Origin.Y) / r.Direction.Y
	return IntersectionList{Intersection(t, p)}
}

func (p *PlaneType) LocalNormalAt(point Tuple) Tuple {
	return Vector(0, 1, 0)
}
