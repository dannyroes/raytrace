package main

type ShapeType struct {
	Transform Matrix
	Material  MaterialType
}

type Shape interface {
	GetMaterial() MaterialType
	SetMaterial(MaterialType)
	GetTransform() Matrix
	SetTransform(Matrix)
	LocalIntersect(RayType) IntersectionList
	LocalNormalAt(Tuple) Tuple
}

func Intersects(s Shape, r RayType) IntersectionList {
	r = transformRay(s, r)
	return s.LocalIntersect(r)
}

func NormalAt(s Shape, p Tuple) Tuple {
	objectPoint := toObjectSpace(s, p)

	objectNormal := s.LocalNormalAt(objectPoint)

	return toWorldNormal(s, objectNormal)
}

func toObjectSpace(o Shape, point Tuple) Tuple {
	return o.GetTransform().Invert().MultiplyTuple(point)
}

func toWorldNormal(o Shape, vector Tuple) Tuple {
	worldNormal := o.GetTransform().Invert().Transpose().MultiplyTuple(vector)
	worldNormal.W = 0

	return worldNormal.Normalize()
}
