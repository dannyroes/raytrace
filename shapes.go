package main

import "math"

type SphereType struct {
	Id        int
	Transform Matrix
	Material  MaterialType
}

func Sphere(id int) SphereType {
	return SphereType{Id: id, Transform: IdentityMatrix(), Material: Material()}
}

func (s SphereType) Intersects(r RayType) IntersectionList {
	inverse := s.Transform.Invert()

	r = r.Transform(inverse)
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

func (s SphereType) SetTransform(m Matrix) SphereType {
	s.Transform = m
	return s
}

func (s SphereType) NormalAt(p Tuple) Tuple {
	objectPoint := s.Transform.Invert().MultiplyTuple(p)
	objectNormal := objectPoint.Sub(Point(0, 0, 0))
	worldNormal := s.Transform.Invert().Transpose().MultiplyTuple(objectNormal)
	worldNormal.W = 0

	return worldNormal.Normalize()
}
