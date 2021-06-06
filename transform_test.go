package main

import (
	"math"
	"testing"
)

func TestTranslate(t *testing.T) {
	cases := []struct {
		p        Tuple
		x        float64
		y        float64
		z        float64
		invert   bool
		expected Tuple
	}{
		{
			p:        Point(-3, 4, 5),
			x:        5,
			y:        -3,
			z:        2,
			expected: Point(2, 1, 7),
		},
		{
			p:        Point(-3, 4, 5),
			x:        5,
			y:        -3,
			z:        2,
			invert:   true,
			expected: Point(-8, 7, 3),
		},
		{
			p:        Vector(-3, 4, 5),
			x:        5,
			y:        -3,
			z:        2,
			expected: Vector(-3, 4, 5),
		},
	}

	for _, tc := range cases {
		translate := Translation(tc.x, tc.y, tc.z)

		if tc.invert {
			translate = translate.Invert()
		}

		result := translate.MultiplyTuple(tc.p)
		if !TupleEqual(result, tc.expected) {
			t.Errorf("expected: %v, received: %v", tc.expected, result)
		}
	}
}

func TestScale(t *testing.T) {
	cases := []struct {
		p        Tuple
		x        float64
		y        float64
		z        float64
		invert   bool
		expected Tuple
	}{
		{
			p:        Point(-4, 6, 8),
			x:        2,
			y:        3,
			z:        4,
			expected: Point(-8, 18, 32),
		},
		{
			p:        Vector(-4, 6, 8),
			x:        2,
			y:        3,
			z:        4,
			expected: Vector(-8, 18, 32),
		},
		{
			p:        Vector(-4, 6, 8),
			x:        2,
			y:        3,
			z:        4,
			invert:   true,
			expected: Vector(-2, 2, 2),
		},
		{
			p:        Point(2, 3, 4),
			x:        -1,
			y:        1,
			z:        1,
			expected: Point(-2, 3, 4),
		},
	}

	for _, tc := range cases {
		scale := Scaling(tc.x, tc.y, tc.z)

		if tc.invert {
			scale = scale.Invert()
		}

		result := scale.MultiplyTuple(tc.p)
		if !TupleEqual(result, tc.expected) {
			t.Errorf("expected: %v, received: %v", tc.expected, result)
		}
	}
}

func TestRotateX(t *testing.T) {
	cases := []struct {
		p        Tuple
		r        Matrix
		invert   bool
		expected Tuple
	}{
		{
			p:        Point(0, 1, 0),
			r:        RotateX(math.Pi / 4),
			expected: Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2),
		},
		{
			p:        Point(0, 1, 0),
			r:        RotateX(math.Pi / 2),
			expected: Point(0, 0, 1),
		},
		{
			p:        Point(0, 1, 0),
			r:        RotateX(math.Pi / 4),
			invert:   true,
			expected: Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2*-1),
		},
	}

	for _, tc := range cases {
		if tc.invert {
			tc.r = tc.r.Invert()
		}
		result := tc.r.MultiplyTuple(tc.p)
		if !TupleEqual(result, tc.expected) {
			t.Errorf("expected: %v, received: %v", tc.expected, result)
		}
	}
}

func TestRotateY(t *testing.T) {
	cases := []struct {
		p        Tuple
		r        Matrix
		invert   bool
		expected Tuple
	}{
		{
			p:        Point(0, 0, 1),
			r:        RotateY(math.Pi / 4),
			expected: Point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2),
		},
		{
			p:        Point(0, 0, 1),
			r:        RotateY(math.Pi / 2),
			expected: Point(1, 0, 0),
		},
	}

	for _, tc := range cases {
		if tc.invert {
			tc.r = tc.r.Invert()
		}
		result := tc.r.MultiplyTuple(tc.p)
		if !TupleEqual(result, tc.expected) {
			t.Errorf("expected: %v, received: %v", tc.expected, result)
		}
	}
}

func TestRotateZ(t *testing.T) {
	cases := []struct {
		p        Tuple
		r        Matrix
		invert   bool
		expected Tuple
	}{
		{
			p:        Point(0, 1, 0),
			r:        RotateZ(math.Pi / 4),
			expected: Point(math.Sqrt(2)/2*-1, math.Sqrt(2)/2, 0),
		},
		{
			p:        Point(0, 1, 0),
			r:        RotateZ(math.Pi / 2),
			expected: Point(-1, 0, 0),
		},
	}

	for _, tc := range cases {
		if tc.invert {
			tc.r = tc.r.Invert()
		}
		result := tc.r.MultiplyTuple(tc.p)
		if !TupleEqual(result, tc.expected) {
			t.Errorf("expected: %v, received: %v", tc.expected, result)
		}
	}
}

func TestShear(t *testing.T) {
	cases := []struct {
		p        Tuple
		r        Matrix
		expected Tuple
	}{
		{
			p:        Point(2, 3, 4),
			r:        Shear(1, 0, 0, 0, 0, 0),
			expected: Point(5, 3, 4),
		},
		{
			p:        Point(2, 3, 4),
			r:        Shear(0, 1, 0, 0, 0, 0),
			expected: Point(6, 3, 4),
		},
		{
			p:        Point(2, 3, 4),
			r:        Shear(0, 0, 1, 0, 0, 0),
			expected: Point(2, 5, 4),
		},
		{
			p:        Point(2, 3, 4),
			r:        Shear(0, 0, 0, 1, 0, 0),
			expected: Point(2, 7, 4),
		},
		{
			p:        Point(2, 3, 4),
			r:        Shear(0, 0, 0, 0, 1, 0),
			expected: Point(2, 3, 6),
		},
		{
			p:        Point(2, 3, 4),
			r:        Shear(0, 0, 0, 0, 0, 1),
			expected: Point(2, 3, 7),
		},
	}

	for _, tc := range cases {
		result := tc.r.MultiplyTuple(tc.p)
		if !TupleEqual(result, tc.expected) {
			t.Errorf("expected: %v, received: %v", tc.expected, result)
		}
	}
}

func TestMatrixHelpers(t *testing.T) {
	cases := []struct {
		result   Tuple
		expected Tuple
	}{
		{
			result:   IdentityMatrix().Translate(5, -3, 2).MultiplyTuple(Point(-3, 4, 5)),
			expected: Point(2, 1, 7),
		},
		{
			result:   IdentityMatrix().Scale(2, 3, 4).MultiplyTuple(Point(-4, 6, 8)),
			expected: Point(-8, 18, 32),
		},
		{
			result:   IdentityMatrix().RotateX(math.Pi / 4).MultiplyTuple(Point(0, 1, 0)),
			expected: Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2),
		},
		{
			result:   IdentityMatrix().RotateY(math.Pi / 4).MultiplyTuple(Point(0, 0, 1)),
			expected: Point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2),
		},
		{
			result:   IdentityMatrix().RotateZ(math.Pi / 4).MultiplyTuple(Point(0, 1, 0)),
			expected: Point(math.Sqrt(2)/2*-1, math.Sqrt(2)/2, 0),
		},
		{
			result:   IdentityMatrix().Shear(1, 0, 0, 0, 0, 0).MultiplyTuple(Point(2, 3, 4)),
			expected: Point(5, 3, 4),
		},
	}

	for _, tc := range cases {
		if !TupleEqual(tc.result, tc.expected) {
			t.Errorf("expected %+v, received %+v", tc.expected, tc.result)
		}
	}
}

func TestMultiple(t *testing.T) {
	point := Point(1, 0, 1)

	a := RotateX(math.Pi / 2)
	b := Scaling(5, 5, 5)
	c := Translation(10, 5, 7)

	p2 := a.MultiplyTuple(point)
	if !TupleEqual(p2, Point(1, -1, 0)) {
		t.Errorf("Individual rotate expected %v received %v", Point(1, -1, 0), p2)
	}

	p3 := b.MultiplyTuple(p2)
	if !TupleEqual(p3, Point(5, -5, 0)) {
		t.Errorf("Individual scale expected %v received %v", Point(5, -5, 0), p3)
	}

	p4 := c.MultiplyTuple(p3)
	if !TupleEqual(p4, Point(15, 0, 7)) {
		t.Errorf("Individual translate expected %v received %v", Point(15, 0, 7), p4)
	}

	result := IdentityMatrix().RotateX(math.Pi/2).Scale(5, 5, 5).Translate(10, 5, 7).MultiplyTuple(point)
	if !TupleEqual(result, Point(15, 0, 7)) {
		t.Errorf("Chained transform expected %v received %v", Point(15, 0, 7), result)
	}
}
