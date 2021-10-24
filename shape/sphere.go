package shape

import (
	"math"

	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
)

type SphereType struct {
	ShapeType
}

func Sphere() *SphereType {
	return &SphereType{ShapeType{Transform: data.IdentityMatrix(), Material: material.Material()}}
}

func GlassSphere() *SphereType {
	s := Sphere()
	m := s.GetMaterial()
	m.Transparency = 1.0
	m.RefractiveIndex = 1.5
	s.SetMaterial(m)

	return s
}

func (s *SphereType) GetMaterial() material.MaterialType {
	return s.Material
}

func (s *SphereType) SetMaterial(m material.MaterialType) {
	s.Material = m
}

func (s *SphereType) LocalIntersect(r data.RayType) IntersectionList {
	sphereRayVector := r.Origin.Sub(data.Point(0, 0, 0))

	dotA := data.Dot(r.Direction, r.Direction)
	dotB := 2 * data.Dot(r.Direction, sphereRayVector)
	dotC := data.Dot(sphereRayVector, sphereRayVector) - 1

	discriminant := math.Pow(dotB, 2) - 4*dotA*dotC
	if discriminant < 0 {
		return IntersectionList{}
	}

	t1 := ((dotB * -1) - math.Sqrt(discriminant)) / (2 * dotA)
	t2 := ((dotB * -1) + math.Sqrt(discriminant)) / (2 * dotA)

	return IntersectionList{Intersection(t1, s), Intersection(t2, s)}
}

func (s *SphereType) SetTransform(m data.Matrix) {
	s.Transform = m
}

func (s *SphereType) GetTransform() data.Matrix {
	return s.Transform
}

func (s *SphereType) LocalNormalAt(objectPoint data.Tuple) data.Tuple {
	return objectPoint.Sub(data.Point(0, 0, 0))
}

func (s *SphereType) GetParent() *GroupType {
	return s.Parent
}

func (s *SphereType) SetParent(p *GroupType) {
	s.Parent = p
}
