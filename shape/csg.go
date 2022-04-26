package shape

import (
	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
)

type CsgOperation int

const (
	CsgUnion CsgOperation = iota
	CsgIntersection
	CsgDifference
)

type CsgType struct {
	ShapeType
	operation CsgOperation
	left      Shape
	right     Shape
}

func Csg(op CsgOperation, left, right Shape) *CsgType {
	csg := &CsgType{
		ShapeType: ShapeType{Transform: data.IdentityMatrix()},
		operation: op,
		left:      left,
		right:     right,
	}

	left.SetParent(csg)
	right.SetParent(csg)
	return csg
}

func (c *CsgType) GetMaterial() material.MaterialType {
	return material.Material()
}

func (c *CsgType) SetMaterial(m material.MaterialType) {

}

func (c *CsgType) GetTransform() data.Matrix {
	return c.Transform
}

func (c *CsgType) SetTransform(t data.Matrix) {
	c.Transform = t
}

func (c *CsgType) LocalIntersect(r data.RayType) IntersectionList {
	leftXs := Intersects(c.left, r)
	rightXs := Intersects(c.right, r)

	xs := append(leftXs, rightXs...)
	xs.Sort()

	return c.filterIntersections(xs)
}

func (c *CsgType) LocalNormalAt(p data.Tuple, i IntersectionType) data.Tuple {
	return data.Vector(0, 0, 0)
}

func (c *CsgType) GetParent() Shape {
	return c.Parent
}

func (c *CsgType) SetParent(p Shape) {
	c.Parent = p
}

func (c *CsgType) Bounds() Bounds {
	return EmptyBounds()
}

func (c *CsgType) CastsShadow() bool {
	return !c.DisableShadow
}

func intersectionAllowed(op CsgOperation, lHit, inL, inR bool) bool {
	switch op {
	case CsgUnion:
		return (lHit && !inR) || (!lHit && !inL)
	case CsgIntersection:
		return (lHit && inR) || (!lHit && inL)
	case CsgDifference:
		return (lHit && !inR) || (!lHit && inL)
	}
	return false
}

func (c *CsgType) filterIntersections(xs IntersectionList) IntersectionList {
	inL := false
	inR := false

	res := Intersections()

	for _, i := range xs {
		lHit := includes(c.left, i.Object)

		if intersectionAllowed(c.operation, lHit, inL, inR) {
			res = append(res, i)
		}

		if lHit {
			inL = !inL
		} else {
			inR = !inR
		}
	}

	return res
}

func includes(a, b Shape) bool {
	switch v := a.(type) {
	case *GroupType:
		for _, c := range v.Children {
			if includes(c, b) {
				return true
			}
		}
		return false
	case *CsgType:
		return includes(v.left, b) || includes(v.right, b)
	default:
		return a == b
	}
}
