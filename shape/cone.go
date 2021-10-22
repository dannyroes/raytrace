package shape

import (
	"fmt"
	"math"

	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
)

type ConeType struct {
	ShapeType
	Minimum float64
	Maximum float64
	Closed  bool
}

func Cone() *ConeType {
	return &ConeType{Minimum: math.Inf(-1), Maximum: math.Inf(1)}
}

func (cone *ConeType) GetMaterial() material.MaterialType {
	return cone.Material
}

func (cone *ConeType) SetMaterial(m material.MaterialType) {
	cone.Material = m
}

func (cone *ConeType) LocalIntersect(r data.RayType) IntersectionList {
	a := math.Pow(r.Direction.X, 2) - math.Pow(r.Direction.Y, 2) + math.Pow(r.Direction.Z, 2)

	b := 2*r.Origin.X*r.Direction.X - 2*r.Origin.Y*r.Direction.Y + 2*r.Origin.Z*r.Direction.Z
	c := math.Pow(r.Origin.X, 2) - math.Pow(r.Origin.Y, 2) + math.Pow(r.Origin.Z, 2)

	xs := Intersections()

	if data.FloatEqual(a, 0) && !data.FloatEqual(b, 0) {
		xs = Intersections(Intersection(-c/(b*2), cone))
	} else {
		disc := math.Pow(b, 2) - 4*a*c

		if disc < 0 {
			return Intersections()
		}

		t0 := (-b - math.Sqrt(disc)) / (2 * a)
		t1 := (-b + math.Sqrt(disc)) / (2 * a)
		if t0 > t1 {
			t0, t1 = t1, t0
		}

		y0 := r.Origin.Y + t0*r.Direction.Y
		if cone.Minimum < y0 && y0 < cone.Maximum {
			xs = append(xs, Intersection(t0, cone))
		}

		y1 := r.Origin.Y + t1*r.Direction.Y
		if cone.Minimum < y1 && y1 < cone.Maximum {
			xs = append(xs, Intersection(t1, cone))
		}
	}
	xs = cone.intersectCaps(r, xs)

	return xs
}

func (cone *ConeType) SetTransform(m data.Matrix) {
	cone.Transform = m
}

func (cone *ConeType) GetTransform() data.Matrix {
	return cone.Transform
}

func (cone *ConeType) LocalNormalAt(objectPoint data.Tuple) data.Tuple {
	dist := math.Pow(objectPoint.X, 2) + math.Pow(objectPoint.Z, 2)
	if dist < 1 && objectPoint.Y >= cone.Maximum-data.Epsilon {
		return data.Vector(0, 1, 0)
	}

	if dist < 1 && objectPoint.Y <= cone.Minimum+data.Epsilon {
		return data.Vector(0, -1, 0)
	}

	y := math.Sqrt(math.Pow(objectPoint.X, 2) + math.Pow(objectPoint.Z, 2))
	if objectPoint.Y > 0 {
		y = -y
	}

	return data.Vector(objectPoint.X, y, objectPoint.Z)
}

// func checkCap(r data.RayType, t float64) bool {
// 	x := r.Origin.X + t*r.Direction.X
// 	z := r.Origin.Z + t*r.Direction.Z

// 	return math.Pow(x, 2)+math.Pow(z, 2) <= 1
// }

func (cone *ConeType) intersectCaps(r data.RayType, xs IntersectionList) IntersectionList {
	fmt.Println(xs)
	if !cone.Closed || data.FloatEqual(r.Direction.Y, 0) {
		return xs
	}

	t := (cone.Minimum - r.Origin.Y) / r.Direction.Y
	if checkCap(r, t, cone.Minimum) {
		xs = append(xs, Intersection(t, cone))
	}

	t = (cone.Maximum - r.Origin.Y) / r.Direction.Y
	if checkCap(r, t, cone.Maximum) {
		xs = append(xs, Intersection(t, cone))
	}

	fmt.Println(xs)

	return xs
}
