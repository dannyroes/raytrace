package shape

import (
	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
)

type ShapeType struct {
	Transform data.Matrix
	Material  material.MaterialType
}

type Shape interface {
	GetMaterial() material.MaterialType
	SetMaterial(material.MaterialType)
	GetTransform() data.Matrix
	SetTransform(data.Matrix)
	LocalIntersect(data.RayType) IntersectionList
	LocalNormalAt(data.Tuple) data.Tuple
}

func Intersects(s Shape, r data.RayType) IntersectionList {
	r = transformRay(s, r)
	return s.LocalIntersect(r)
}

func NormalAt(s Shape, p data.Tuple) data.Tuple {
	objectPoint := toObjectSpace(s, p)

	objectNormal := s.LocalNormalAt(objectPoint)

	return toWorldNormal(s, objectNormal)
}

func PatternAtObject(p material.Pattern, o Shape, point data.Tuple) material.ColourTuple {
	objectPoint := o.GetTransform().Invert().MultiplyTuple(point)
	patternPoint := p.GetTransform().Invert().MultiplyTuple(objectPoint)

	return p.At(patternPoint)
}

func transformRay(o Shape, r data.RayType) data.RayType {
	inverse := o.GetTransform().Invert()
	return r.Transform(inverse)
}

func toObjectSpace(o Shape, point data.Tuple) data.Tuple {
	return o.GetTransform().Invert().MultiplyTuple(point)
}

func toWorldNormal(o Shape, vector data.Tuple) data.Tuple {
	worldNormal := o.GetTransform().Invert().Transpose().MultiplyTuple(vector)
	worldNormal.W = 0

	return worldNormal.Normalize()
}
