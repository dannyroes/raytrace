package shape

import (
	"math"

	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
)

type CubeType struct {
	ShapeType
}

func (c *CubeType) GetMaterial() material.MaterialType {
	return c.Material
}

func (c *CubeType) SetMaterial(m material.MaterialType) {
	c.Material = m
}

func (c *CubeType) GetTransform() data.Matrix {
	return c.Transform
}

func (c *CubeType) SetTransform(m data.Matrix) {
	c.Transform = m
}

func (c *CubeType) LocalIntersect(r data.RayType) IntersectionList {
	xtmin, xtmax := checkAxis(r.Origin.X, r.Direction.X)
	ytmin, ytmax := checkAxis(r.Origin.Y, r.Direction.Y)
	ztmin, ztmax := checkAxis(r.Origin.Z, r.Direction.Z)

	tmin := data.FloatMax(xtmin, ytmin, ztmin)
	tmax := data.FloatMin(xtmax, ytmax, ztmax)

	if tmin > tmax {
		return Intersections()
	}

	return Intersections(Intersection(tmin, c), Intersection(tmax, c))
}

func (c *CubeType) LocalNormalAt(objectPoint data.Tuple) data.Tuple {
	maxc := data.FloatMax(math.Abs(objectPoint.X), math.Abs(objectPoint.Y), math.Abs(objectPoint.Z))

	if data.FloatEqual(maxc, math.Abs(objectPoint.X)) {
		return data.Vector(objectPoint.X, 0, 0)
	} else if data.FloatEqual(maxc, math.Abs(objectPoint.Y)) {
		return data.Vector(0, objectPoint.Y, 0)
	}

	return data.Vector(0, 0, objectPoint.Z)
}

func (c *CubeType) GetParent() *GroupType {
	return c.Parent
}

func (c *CubeType) SetParent(p *GroupType) {
	c.Parent = p
}

func Cube() *CubeType {
	return &CubeType{}
}

func (c *CubeType) Bounds() Bounds {
	return Bounds{
		Min: data.Point(-1, -1, -1),
		Max: data.Point(1, 1, 1),
	}
}

func checkAxis(origin, direction float64) (float64, float64) {
	tminNum := -1 - origin
	tmaxNum := 1 - origin

	tmin := 0.0
	tmax := 0.0

	if math.Abs(direction) >= data.Epsilon {
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
