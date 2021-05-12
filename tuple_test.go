package main

import "testing"

func TestFloatEqual(t *testing.T) {
	a := 0.1
	a += 0.2
	a += 0.2

	b := 0.5

	if !FloatEqual(a, b) {
		t.Error("Expected true")
	}

	a += 0.01

	if FloatEqual(a, b) {
		t.Error("Expected false")
	}
}

func TestPointTuple(t *testing.T) {
	a := Tuple{
		X: 4.3,
		Y: -4.2,
		Z: 3.1,
		W: 1.0,
	}

	if !FloatEqual(a.X, 4.3) {
		t.Error("X value incorrect")
	}

	if !FloatEqual(a.Y, -4.2) {
		t.Error("Y value incorrect")
	}

	if !FloatEqual(a.Z, 3.1) {
		t.Error("Z value incorrect")
	}

	if !FloatEqual(a.W, 1.0) {
		t.Error("W value incorrect")
	}

	if !a.IsPoint() || a.IsVector() {
		t.Error("Point reported as vector")
	}
}

func TestVectorTuple(t *testing.T) {
	a := Tuple{
		X: 4.3,
		Y: -4.2,
		Z: 3.1,
		W: 0.0,
	}

	if !FloatEqual(a.X, 4.3) {
		t.Error("X value incorrect")
	}

	if !FloatEqual(a.Y, -4.2) {
		t.Error("Y value incorrect")
	}

	if !FloatEqual(a.Z, 3.1) {
		t.Error("Z value incorrect")
	}

	if !FloatEqual(a.W, 0.0) {
		t.Error("W value incorrect")
	}

	if a.IsPoint() || !a.IsVector() {
		t.Error("Point reported as vector")
	}
}

func TestTupleEqual(t *testing.T) {
	a := Tuple{
		X: 4.3,
		Y: -4.2,
		Z: 3.1,
		W: 0.0,
	}

	b := a

	if !TupleEqual(a, b) {
		t.Error("Expected equal")
	}

	b.X += 0.2

	if TupleEqual(a, b) {
		t.Error("Expected unequal")
	}
}

func TestPointFactory(t *testing.T) {
	a := Tuple{
		X: 4,
		Y: -4,
		Z: 3,
		W: 1.0,
	}

	b := Point(4, -4, 3)

	if !TupleEqual(a, b) {
		t.Error("Expected equal")
	}
}

func TestVectorFactory(t *testing.T) {
	a := Tuple{
		X: 4,
		Y: -4,
		Z: 3,
		W: 0,
	}

	b := Vector(4, -4, 3)

	if !TupleEqual(a, b) {
		t.Error("Expected equal")
	}
}

func TestAddTuple(t *testing.T) {
	a := Point(3, -2, 5)
	b := Vector(-2, 3, 1)

	c := a.Add(b)
	expected := Point(1, 1, 6)

	if !TupleEqual(c, expected) {
		t.Errorf("Expected %+v, received %+v", expected, c)
	}

}
