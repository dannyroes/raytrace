package shape

import (
	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
)

type ShapeType struct {
	Transform     data.Matrix
	Material      material.MaterialType
	Parent        Shape
	DisableShadow bool
}

type Shape interface {
	GetMaterial() material.MaterialType
	SetMaterial(material.MaterialType)
	GetTransform() data.Matrix
	SetTransform(data.Matrix)
	LocalIntersect(data.RayType) IntersectionList
	LocalNormalAt(data.Tuple, IntersectionType) data.Tuple
	GetParent() Shape
	SetParent(Shape)
	Bounds() Bounds
	CastsShadow() bool
}

func Intersects(s Shape, r data.RayType) IntersectionList {
	r = transformRay(s, r)
	return s.LocalIntersect(r)
}

func NormalAt(s Shape, p data.Tuple, i IntersectionType) data.Tuple {
	localPoint := worldToObject(s, p)
	objectNormal := s.LocalNormalAt(localPoint, i)

	return normalToWorld(s, objectNormal)
}

func PatternAtObject(p material.Pattern, o Shape, point data.Tuple) material.ColourTuple {
	objectPoint := worldToObject(o, point)
	patternPoint := p.GetTransform().Invert().MultiplyTuple(objectPoint)

	return p.At(patternPoint)
}

func transformRay(o Shape, r data.RayType) data.RayType {
	inverse := o.GetTransform().Invert()
	return r.Transform(inverse)
}

func worldToObject(o Shape, point data.Tuple) data.Tuple {
	if o.GetParent() != nil {
		point = worldToObject(o.GetParent(), point)
	}

	return o.GetTransform().Invert().MultiplyTuple(point)
}

func normalToWorld(o Shape, normal data.Tuple) data.Tuple {
	normal = o.GetTransform().Invert().Transpose().MultiplyTuple(normal)
	normal.W = 0
	normal = normal.Normalize()

	if o.GetParent() != nil {
		normal = normalToWorld(o.GetParent(), normal)
	}

	return normal
}
