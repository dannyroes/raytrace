package main

import "math"

type SphereType struct {
	ShapeType
}

func Sphere() *SphereType {
	return &SphereType{ShapeType{Transform: IdentityMatrix(), Material: Material()}}
}

func (s *SphereType) GetMaterial() MaterialType {
	return s.Material
}

func (s *SphereType) SetMaterial(m MaterialType) {
	s.Material = m
}

func (s *SphereType) LocalIntersect(r RayType) IntersectionList {
	sphereRayVector := r.Origin.Sub(Point(0, 0, 0))

	dotA := Dot(r.Direction, r.Direction)
	dotB := 2 * Dot(r.Direction, sphereRayVector)
	dotC := Dot(sphereRayVector, sphereRayVector) - 1

	discriminant := math.Pow(dotB, 2) - 4*dotA*dotC
	if discriminant < 0 {
		return IntersectionList{}
	}

	t1 := ((dotB * -1) - math.Sqrt(discriminant)) / (2 * dotA)
	t2 := ((dotB * -1) + math.Sqrt(discriminant)) / (2 * dotA)

	return IntersectionList{Intersection(t1, s), Intersection(t2, s)}
}

func transformRay(o Shape, r RayType) RayType {
	inverse := o.GetTransform().Invert()
	return r.Transform(inverse)
}

func (s *SphereType) SetTransform(m Matrix) {
	s.Transform = m
}

func (s *SphereType) GetTransform() Matrix {
	return s.Transform
}

func (s *SphereType) LocalNormalAt(objectPoint Tuple) Tuple {
	return objectPoint.Sub(Point(0, 0, 0))
}
