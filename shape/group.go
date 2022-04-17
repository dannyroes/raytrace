package shape

import (
	"fmt"
	"math"
	"strconv"

	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
)

type GroupType struct {
	ShapeType
	Children    []Shape
	GroupBounds *Bounds
}

type Bounds struct {
	Min data.Tuple
	Max data.Tuple
}

func Group() *GroupType {
	return &GroupType{ShapeType: ShapeType{Transform: data.IdentityMatrix()}}
}

func (g *GroupType) AddChild(s ...Shape) {
	g.Children = append(g.Children, s...)
	for _, shape := range s {
		shape.SetParent(g)
	}
	g.GroupBounds = nil
}

func (g *GroupType) LocalIntersect(r data.RayType) IntersectionList {
	xs := Intersections()

	if g.boundIntersect(r) {
		for _, s := range g.Children {
			xs = append(xs, Intersects(s, r)...)
		}
		xs.Sort()
	}

	return xs
}

func (g *GroupType) LocalNormalAt(objectPoint data.Tuple) data.Tuple {
	panic("called localNormalAt on group")
}

func (g *GroupType) GetMaterial() material.MaterialType {
	return g.Material
}

func (g *GroupType) SetMaterial(m material.MaterialType) {
	for _, o := range g.Children {
		o.SetMaterial(m)
	}
}

func (g *GroupType) GetTransform() data.Matrix {
	return g.Transform
}

func (g *GroupType) SetTransform(m data.Matrix) {
	g.Transform = m
}

func (g *GroupType) GetParent() *GroupType {
	return g.Parent
}

func (g *GroupType) SetParent(p *GroupType) {
	g.Parent = p
}

func (g *GroupType) Bounds() Bounds {
	if g.GroupBounds == nil {
		b := EmptyBounds()
		for _, c := range g.Children {
			points := boundsToPoints(c.Bounds().Min, c.Bounds().Max)
			for _, p := range points {
				t := c.GetTransform().MultiplyTuple(p)
				if t.X < b.Min.X {
					b.Min.X = t.X
				}
				if t.X > b.Max.X {
					b.Max.X = t.X
				}

				if t.Y < b.Min.Y {
					b.Min.Y = t.Y
				}
				if t.Y > b.Max.Y {
					b.Max.Y = t.Y
				}

				if t.Z < b.Min.Z {
					b.Min.Z = t.Z
				}
				if t.Z > b.Max.Z {
					b.Max.Z = t.Z
				}
			}
		}
		g.GroupBounds = &b
	}

	return *g.GroupBounds
}

func (g *GroupType) boundIntersect(r data.RayType) bool {
	b := g.Bounds()

	xtmin, xtmax := checkArbitraryAxis(b.Min.X, b.Max.X, r.Origin.X, r.Direction.X)
	ytmin, ytmax := checkArbitraryAxis(b.Min.Y, b.Max.Y, r.Origin.Y, r.Direction.Y)
	ztmin, ztmax := checkArbitraryAxis(b.Min.Z, b.Max.Z, r.Origin.Z, r.Direction.Z)

	tmin := data.FloatMax(xtmin, ytmin, ztmin)
	tmax := data.FloatMin(xtmax, ytmax, ztmax)

	return tmin <= tmax
}

func checkArbitraryAxis(min, max, origin, direction float64) (float64, float64) {
	tminNum := min - origin
	tmaxNum := max - origin

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

func EmptyBounds() Bounds {
	posInf := math.Inf(1)
	negInf := math.Inf(-1)
	return Bounds{
		Min: data.Point(posInf, posInf, posInf),
		Max: data.Point(negInf, negInf, negInf),
	}
}

func boundsToPoints(min, max data.Tuple) []data.Tuple {
	points := make([]data.Tuple, 8)
	var pointX, pointY, pointZ float64
	for x := 0; x < 2; x++ {
		if x == 0 {
			pointX = min.X
		} else {
			pointX = max.X
		}
		for y := 0; y < 2; y++ {
			if y == 0 {
				pointY = min.Y
			} else {
				pointY = max.Y
			}
			for z := 0; z < 2; z++ {
				if z == 0 {
					pointZ = min.Z
				} else {
					pointZ = max.Z
				}

				i, _ := strconv.ParseInt(fmt.Sprintf("%d%d%d", x, y, z), 2, 32)
				points[i] = data.Point(pointX, pointY, pointZ)
			}
		}
	}

	return points
}
