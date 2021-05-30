package main

import "sort"

type IntersectionType struct {
	T      float64
	Object SphereType
}

type IntersectionList []IntersectionType

func Intersection(t float64, s SphereType) IntersectionType {
	return IntersectionType{T: t, Object: s}
}

func Intersections(ints ...IntersectionType) IntersectionList {
	sort.Slice(ints, func(i, j int) bool { return ints[i].T < ints[j].T })
	return ints
}

func (l IntersectionList) Hit() IntersectionType {
	for _, i := range l {
		if i.T >= 0 {
			return i
		}
	}
	return IntersectionType{}
}
