package main

import (
	"math"
	"testing"
)

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

func TestTupleBasics(t *testing.T) {
	cases := []struct {
		a        Tuple
		x        float64
		y        float64
		z        float64
		w        float64
		isPoint  bool
		isVector bool
	}{
		{
			a: Tuple{
				X: 4.3,
				Y: -4.2,
				Z: 3.1,
				W: 1.0,
			},
			x:        4.3,
			y:        -4.2,
			z:        3.1,
			w:        1.0,
			isPoint:  true,
			isVector: false,
		},
		{
			a: Tuple{
				X: 4.3,
				Y: -4.2,
				Z: 3.1,
				W: 0.0,
			},
			x:        4.3,
			y:        -4.2,
			z:        3.1,
			w:        0.0,
			isPoint:  false,
			isVector: true,
		},
	}

	for _, c := range cases {
		tupleCheck(t, c.a, c.x, c.y, c.z, c.w, c.isPoint, c.isVector)
	}
}

func tupleCheck(t *testing.T, a Tuple, x, y, z, w float64, isPoint, isVector bool) {
	if !FloatEqual(a.X, x) {
		t.Errorf("X value incorrect. expected: %f; actual: %f", x, a.X)
	}

	if !FloatEqual(a.Y, y) {
		t.Errorf("Y value incorrect. expected: %f; actual: %f", y, a.Y)
	}

	if !FloatEqual(a.Z, z) {
		t.Errorf("Z value incorrect. expected: %f; actual: %f", z, a.Z)
	}

	if !FloatEqual(a.W, w) {
		t.Errorf("W value incorrect. expected: %f; actual: %f", w, a.W)
	}

	if a.IsPoint() != isPoint {
		t.Errorf("IsPoint incorrect. expected: %v; actual: %v", isPoint, a.IsPoint())
	}

	if a.IsVector() != isVector {
		t.Errorf("IsVector incorrect. expected: %v; actual: %v", isVector, a.IsVector())
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

func TestSubtractTuple(t *testing.T) {
	cases := []struct {
		a        Tuple
		b        Tuple
		expected Tuple
	}{
		{
			a:        Point(3, 2, 1),
			b:        Point(5, 6, 7),
			expected: Vector(-2, -4, -6),
		},
		{
			a:        Point(3, 2, 1),
			b:        Vector(5, 6, 7),
			expected: Point(-2, -4, -6),
		},
		{
			a:        Vector(3, 2, 1),
			b:        Vector(5, 6, 7),
			expected: Vector(-2, -4, -6),
		},
		{
			a:        Vector(0, 0, 0),
			b:        Vector(1, -2, 3),
			expected: Vector(-1, 2, -3),
		},
	}

	for _, c := range cases {
		result := c.a.Sub(c.b)
		if !TupleEqual(result, c.expected) {
			t.Errorf("Expected %+v, received %+v", c.expected, result)
		}
	}
}

func TestNegateTuple(t *testing.T) {
	a := Tuple{X: 1, Y: -2, Z: 3, W: -4}
	expected := Tuple{X: -1, Y: 2, Z: -3, W: 4}

	result := a.Neg()
	if !TupleEqual(result, expected) {
		t.Errorf("Expected %+v, received %+v", expected, result)
	}
}

func TestMultiplyTuple(t *testing.T) {
	cases := []struct {
		a        Tuple
		mul      float64
		expected Tuple
	}{
		{
			a:        Tuple{X: 1, Y: -2, Z: 3, W: -4},
			mul:      3.5,
			expected: Tuple{X: 3.5, Y: -7, Z: 10.5, W: -14},
		},
		{
			a:        Tuple{X: 1, Y: -2, Z: 3, W: -4},
			mul:      0.5,
			expected: Tuple{X: 0.5, Y: -1, Z: 1.5, W: -2},
		},
	}

	for _, c := range cases {
		result := c.a.Mul(c.mul)
		if !TupleEqual(result, c.expected) {
			t.Errorf("Expected %+v, received %+v", c.expected, result)
		}
	}
}

func TestDivideTuple(t *testing.T) {
	cases := []struct {
		a        Tuple
		div      float64
		expected Tuple
	}{
		{
			a:        Tuple{X: 1, Y: -2, Z: 3, W: -4},
			div:      2,
			expected: Tuple{X: 0.5, Y: -1, Z: 1.5, W: -2},
		},
	}

	for _, c := range cases {
		result := c.a.Div(c.div)
		if !TupleEqual(result, c.expected) {
			t.Errorf("Expected %+v, received %+v", c.expected, result)
		}
	}
}

func TestMagnitudeVector(t *testing.T) {
	cases := []struct {
		a        Tuple
		expected float64
	}{
		{
			a:        Vector(1, 0, 0),
			expected: 1,
		},
		{
			a:        Vector(0, 1, 0),
			expected: 1,
		},
		{
			a:        Vector(0, 0, 1),
			expected: 1,
		},
		{
			a:        Vector(1, 2, 3),
			expected: math.Sqrt(14),
		},
		{
			a:        Vector(-1, -2, -3),
			expected: math.Sqrt(14),
		},
	}

	for _, c := range cases {
		result := c.a.Magnitude()
		if !FloatEqual(result, c.expected) {
			t.Errorf("Expected %+f, received %+f", c.expected, result)
		}
	}
}

func TestNormalizeVector(t *testing.T) {
	cases := []struct {
		a        Tuple
		expected Tuple
	}{
		{
			a:        Vector(4, 0, 0),
			expected: Vector(1, 0, 0),
		},
		{
			a:        Vector(1, 2, 3),
			expected: Vector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14)),
		},
	}

	for _, c := range cases {
		result := c.a.Normalize()
		if !TupleEqual(result, c.expected) {
			t.Errorf("Expected %+v, received %+v", c.expected, result)
		}

		if !FloatEqual(result.Magnitude(), 1) {
			t.Errorf("Normalized magnitude incorrect. Expected 1.0, received %+v", result.Magnitude())
		}
	}
}

func TestDotVector(t *testing.T) {
	cases := []struct {
		a        Tuple
		b        Tuple
		expected float64
	}{
		{
			a:        Vector(1, 2, 3),
			b:        Vector(2, 3, 4),
			expected: 20,
		},
	}

	for _, c := range cases {
		result := Dot(c.a, c.b)
		if !FloatEqual(result, c.expected) {
			t.Errorf("Expected %+v, received %+v", c.expected, result)
		}
	}
}

func TestCrossVector(t *testing.T) {
	cases := []struct {
		a        Tuple
		b        Tuple
		expected Tuple
	}{
		{
			a:        Vector(1, 2, 3),
			b:        Vector(2, 3, 4),
			expected: Vector(-1, 2, -1),
		},
	}

	for _, c := range cases {
		result := Cross(c.a, c.b)
		if !TupleEqual(result, c.expected) {
			t.Errorf("Expected %+v, received %+v", c.expected, result)
		}

		result = Cross(c.b, c.a)
		if !TupleEqual(result, c.expected.Neg()) {
			t.Errorf("Expected %+v, received %+v", c.expected.Neg(), result)
		}
	}
}
