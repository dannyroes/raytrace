package main

import "sort"

type IntersectionType struct {
	T      float64
	Object Object
}

type IntersectionList []IntersectionType

func Intersection(t float64, o Object) IntersectionType {
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
	return IntersectionType{}
}

func (l IntersectionList) Sort() IntersectionList {
	sort.Slice(l, func(i, j int) bool { return l[i].T < l[j].T })
	return l
}

type Computations struct {
	T       float64
	Object  Object
	Point   Tuple
	EyeV    Tuple
	NormalV Tuple
	Inside  bool
}

func (i IntersectionType) PrepareComputations(r RayType) Computations {
	comp := Computations{}
	comp.T = i.T
	comp.Object = i.Object
	comp.Point = r.Position(comp.T)
	comp.EyeV = r.Direction.Neg()
	comp.NormalV = comp.Object.NormalAt(comp.Point)

	if Dot(comp.NormalV, comp.EyeV) < 0 {
		comp.Inside = true
		comp.NormalV = comp.NormalV.Neg()
	}

	return comp
}
