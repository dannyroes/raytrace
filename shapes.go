package main

import "math"

type SphereType struct {
	ShapeType
}

type ShapeType struct {
	Transform Matrix
	Material  MaterialType
}

type Shape interface {
	GetMaterial() MaterialType
	SetMaterial(MaterialType)
	GetTransform() Matrix
	LocalIntersect(RayType) IntersectionList
	LocalNormal(Tuple) Tuple
}

func Sphere(id int) *SphereType {
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

func Intersects(s Shape, r RayType) IntersectionList {
	r = transformRay(s, r)
	return s.LocalIntersect(r)
}

func NormalAt(s Shape, p Tuple) Tuple {
	objectPoint := toObjectSpace(s, p)

	objectNormal := s.LocalNormal(objectPoint)

	return toWorldNormal(s, objectNormal)
}

func (s *SphereType) LocalNormal(objectPoint Tuple) Tuple {
	return objectPoint.Sub(Point(0, 0, 0))
}

func toObjectSpace(o Shape, point Tuple) Tuple {
	return o.GetTransform().Invert().MultiplyTuple(point)
}

func toWorldNormal(o Shape, vector Tuple) Tuple {
	worldNormal := o.GetTransform().Invert().Transpose().MultiplyTuple(vector)
	worldNormal.W = 0

	return worldNormal.Normalize()
}
