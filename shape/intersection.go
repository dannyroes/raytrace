package shape

import (
	"math"
	"sort"

	"github.com/dannyroes/raytrace/data"
)

type IntersectionType struct {
	T      float64
	Object Shape
}

type IntersectionList []IntersectionType

func Intersection(t float64, o Shape) IntersectionType {
	return IntersectionType{T: t, Object: o}
}

func Intersections(ints ...IntersectionType) IntersectionList {
	return IntersectionList(ints).Sort()
}

func (l IntersectionList) Hit() IntersectionType {
	for _, i := range l {
		if i.T >= 0 {
			return i
		}
	}
	return IntersectionType{T: -1}
}

func (l IntersectionList) Sort() IntersectionList {
	sort.Slice(l, func(i, j int) bool { return l[i].T < l[j].T })
	return l
}

type Computations struct {
	T          float64
	Object     Shape
	Point      data.Tuple
	EyeV       data.Tuple
	NormalV    data.Tuple
	Inside     bool
	OverPoint  data.Tuple
	ReflectV   data.Tuple
	N1         float64
	N2         float64
	UnderPoint data.Tuple
}

func (i IntersectionType) PrepareComputations(r data.RayType, xs ...IntersectionType) Computations {
	if len(xs) == 0 {
		xs = IntersectionList{i}
	}

	comp := Computations{}
	comp.T = i.T
	comp.Object = i.Object
	comp.Point = r.Position(comp.T)
	comp.EyeV = r.Direction.Neg()
	comp.NormalV = NormalAt(comp.Object, comp.Point)

	if data.Dot(comp.NormalV, comp.EyeV) < 0 {
		comp.Inside = true
		comp.NormalV = comp.NormalV.Neg()
	}

	comp.ReflectV = r.Direction.Reflect(comp.NormalV)
	comp.OverPoint = comp.Point.Add(comp.NormalV.Mul(data.Epsilon))
	comp.UnderPoint = comp.Point.Sub(comp.NormalV.Mul(data.Epsilon))

	var containers []Shape

	for _, x := range xs {
		if x.Object == i.Object && x.T == i.T {
			if len(containers) == 0 {
				comp.N1 = 1.0
			} else {
				comp.N1 = containers[len(containers)-1].GetMaterial().RefractiveIndex
			}
		}

		if index := checkContainers(containers, x.Object); index >= 0 {
			containers = append(containers[:index], containers[index+1:]...)
		} else {
			containers = append(containers, x.Object)
		}

		if x.Object == i.Object && x.T == i.T {
			if len(containers) == 0 {
				comp.N2 = 1.0
			} else {
				comp.N2 = containers[len(containers)-1].GetMaterial().RefractiveIndex
			}
		}
	}

	return comp
}

func (c Computations) Schlick() float64 {
	cos := data.Dot(c.EyeV, c.NormalV)
	if c.N1 > c.N2 {
		n := c.N1 / c.N2
		sin2t := math.Pow(n, 2) * (1 - math.Pow(cos, 2))
		if sin2t > 1.0 {
			return 1.0
		}

		cost := math.Sqrt(1.0 - sin2t)
		cos = cost
	}

	r0 := math.Pow((c.N1-c.N2)/(c.N1+c.N2), 2)
	return r0 + (1-r0)*math.Pow(1-cos, 5)
}

func checkContainers(containers []Shape, o Shape) int {
	for i, s := range containers {
		if o == s {
			return i
		}
	}

	return -1
}
