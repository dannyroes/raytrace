package main

import "sort"

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
	T         float64
	Object    Shape
	Point     Tuple
	EyeV      Tuple
	NormalV   Tuple
	Inside    bool
	OverPoint Tuple
	ReflectV  Tuple
}

func (i IntersectionType) PrepareComputations(r RayType) Computations {
	comp := Computations{}
	comp.T = i.T
	comp.Object = i.Object
	comp.Point = r.Position(comp.T)
	comp.EyeV = r.Direction.Neg()
	comp.NormalV = NormalAt(comp.Object, comp.Point)

	if Dot(comp.NormalV, comp.EyeV) < 0 {
		comp.Inside = true
		comp.NormalV = comp.NormalV.Neg()
	}

	comp.ReflectV = r.Direction.Reflect(comp.NormalV)
	comp.OverPoint = comp.Point.Add(comp.NormalV.Mul(Epsilon))

	return comp
}
