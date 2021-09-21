package main

import "math"

type CubeType struct {
	ShapeType
}

func (c *CubeType) GetMaterial() MaterialType {
	return c.Material
}

func (c *CubeType) SetMaterial(m MaterialType) {
	c.Material = m
}

func (c *CubeType) GetTransform() Matrix {
	return c.Transform
}

func (c *CubeType) SetTransform(m Matrix) {
	c.Transform = m
}

func (c *CubeType) LocalIntersect(r RayType) IntersectionList {
	xtmin, xtmax := checkAxis(r.Origin.X, r.Direction.X)
	ytmin, ytmax := checkAxis(r.Origin.Y, r.Direction.Y)
	ztmin, ztmax := checkAxis(r.Origin.Z, r.Direction.Z)

	tmin := FloatMax(xtmin, ytmin, ztmin)
	tmax := FloatMin(xtmax, ytmax, ztmax)

	if tmin > tmax {
		return Intersections()
	}

	return Intersections(Intersection(tmin, c), Intersection(tmax, c))
}

func (c *CubeType) LocalNormalAt(objectPoint Tuple) Tuple {
	maxc := FloatMax(math.Abs(objectPoint.X), math.Abs(objectPoint.Y), math.Abs(objectPoint.Z))

	if FloatEqual(maxc, math.Abs(objectPoint.X)) {
		return Vector(objectPoint.X, 0, 0)
	} else if FloatEqual(maxc, math.Abs(objectPoint.Y)) {
		return Vector(0, objectPoint.Y, 0)
	}

	return Vector(0, 0, objectPoint.Z)
}

func Cube() *CubeType {
	return &CubeType{}
}

func checkAxis(origin, direction float64) (float64, float64) {
	tminNum := -1 - origin
	tmaxNum := 1 - origin

	tmin := 0.0
	tmax := 0.0

	if math.Abs(direction) >= Epsilon {
		tmin = tminNum / direction
		tmax = tmaxNum / direction
	} else {
		tmin = math.Inf(int(tminNum))
		tmax = math.Inf(int(tmaxNum))
	}

	if tmin > tmax {
		tmin, tmax = tmax, tmin
	}

	return tmin, tmax
}
