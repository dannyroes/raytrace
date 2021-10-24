package shape

import (
	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
)

type GroupType struct {
	ShapeType
	Children []Shape
}

func Group() *GroupType {
	return &GroupType{ShapeType: ShapeType{Transform: data.IdentityMatrix()}}
}

func (g *GroupType) AddChild(s ...Shape) {
	g.Children = append(g.Children, s...)
	for _, shape := range s {
		shape.SetParent(g)
	}
}

func (g *GroupType) LocalIntersect(r data.RayType) IntersectionList {
	xs := Intersections()

	for _, s := range g.Children {
		xs = append(xs, Intersects(s, r)...)
	}
	xs.Sort()

	return xs
}

func (g *GroupType) LocalNormalAt(objectPoint data.Tuple) data.Tuple {
	panic("called localNormalAt on group")
	// return data.Vector(0, 0, 0)
}

func (g *GroupType) GetMaterial() material.MaterialType {
	return g.Material
}

func (g *GroupType) SetMaterial(m material.MaterialType) {
	g.Material = m
}

func (g *GroupType) GetTransform() data.Matrix {
	return g.Transform
}

func (g *GroupType) SetTransform(m data.Matrix) {
	g.Transform = m
}

func (g *GroupType) GetParent() *GroupType {
	return g.Parent
}

func (g *GroupType) SetParent(p *GroupType) {
	g.Parent = p
}
