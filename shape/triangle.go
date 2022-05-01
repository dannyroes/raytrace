package shape

import (
	"math"

	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
)

type TriangleType struct {
	ShapeType
	p1     data.Tuple
	p2     data.Tuple
	p3     data.Tuple
	e1     data.Tuple
	e2     data.Tuple
	normal data.Tuple
	smooth bool
	n1     data.Tuple
	n2     data.Tuple
	n3     data.Tuple
}

func Triangle(p1, p2, p3 data.Tuple) *TriangleType {
	t := &TriangleType{p1: p1, p2: p2, p3: p3,
		ShapeType: ShapeType{
			Material:  material.Material(),
			Transform: data.IdentityMatrix(),
		},
	}
	t.compute()

	return t
}

func SmoothTriangle(p1, p2, p3, n1, n2, n3 data.Tuple) *TriangleType {
	t := &TriangleType{p1: p1, p2: p2, p3: p3, n1: n1, n2: n2, n3: n3,
		smooth: true,
		ShapeType: ShapeType{
			Material:  material.Material(),
			Transform: data.IdentityMatrix(),
		},
	}
	t.compute()

	return t
}

func (t *TriangleType) compute() {
	t.e1 = t.p2.Sub(t.p1)
	t.e2 = t.p3.Sub(t.p1)

	t.normal = data.Cross(t.e2, t.e1).Normalize()
}

func (t *TriangleType) LocalNormalAt(objectPoint data.Tuple, i IntersectionType) data.Tuple {
	if t.smooth {
		return t.n2.Mul(i.U).Add(t.n3.Mul(i.V)).Add(t.n1.Mul(1 - i.U - i.V))
	}
	return t.normal
}

func (t *TriangleType) LocalIntersect(r data.RayType) IntersectionList {
	dirCrossE2 := data.Cross(r.Direction, t.e2)
	det := data.Dot(t.e1, dirCrossE2)

	if math.Abs(det) <= data.Epsilon {
		return Intersections()
	}

	f := 1.0 / det

	p1ToOrigin := r.Origin.Sub(t.p1)
	u := f * data.Dot(p1ToOrigin, dirCrossE2)

	if u < 0 || u > 1 {
		return Intersections()
	}

	originCrossE1 := data.Cross(p1ToOrigin, t.e1)
	v := f * data.Dot(r.Direction, originCrossE1)

	if v < 0 || u+v > 1 {
		return Intersections()
	}

	time := f * data.Dot(t.e2, originCrossE1)
	if t.smooth {
		return Intersections(IntersectionWithUv(time, t, u, v))
	} else {
		return Intersections(Intersection(time, t))
	}

}

func (t *TriangleType) GetParent() Shape {
	return t.Parent
}

func (t *TriangleType) SetParent(p Shape) {
	t.Parent = p
}

func (t *TriangleType) Bounds() Bounds {
	b := Bounds{}
	b.Min.X = data.FloatMin(t.p1.X, t.p2.X, t.p3.X)
	b.Min.Y = data.FloatMin(t.p1.Y, t.p2.Y, t.p3.Y)
	b.Min.Z = data.FloatMin(t.p1.Z, t.p2.Z, t.p3.Z)

	b.Max.X = data.FloatMax(t.p1.X, t.p2.X, t.p3.X)
	b.Max.Y = data.FloatMax(t.p1.Y, t.p2.Y, t.p3.Y)
	b.Max.Z = data.FloatMax(t.p1.Z, t.p2.Z, t.p3.Z)

	return b
}

func (t *TriangleType) GetMaterial() material.MaterialType {
	return t.Material
}

func (t *TriangleType) SetMaterial(m material.MaterialType) {
	t.Material = m
}

func (t *TriangleType) GetTransform() data.Matrix {
	return t.Transform
}

func (t *TriangleType) SetTransform(m data.Matrix) {
	t.Transform = m
}

func (t *TriangleType) CastsShadow() bool {
	return !t.DisableShadow
}
