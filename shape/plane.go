package shape

import (
	"math"

	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
)

type PlaneType struct {
	ShapeType
}

func Plane() *PlaneType {
	return &PlaneType{ShapeType{Transform: data.IdentityMatrix(), Material: material.Material()}}
}

func (p *PlaneType) SetMaterial(m material.MaterialType) {
	p.Material = m
}

func (p *PlaneType) GetMaterial() material.MaterialType {
	return p.Material
}

func (p *PlaneType) GetTransform() data.Matrix {
	return p.Transform
}

func (p *PlaneType) SetTransform(m data.Matrix) {
	p.Transform = m
}

func (p *PlaneType) LocalIntersect(r data.RayType) IntersectionList {
	if math.Abs(r.Direction.Y) < data.Epsilon {
		return IntersectionList{}
	}

	t := (-1 * r.Origin.Y) / r.Direction.Y
	return IntersectionList{Intersection(t, p)}
}

func (p *PlaneType) LocalNormalAt(point data.Tuple) data.Tuple {
	return data.Vector(0, 1, 0)
}
