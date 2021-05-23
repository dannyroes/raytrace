package main

import (
	"testing"
)

type matrixCheck struct {
	X     int
	Y     int
	Value float64
}

func TestMatrixBasics(t *testing.T) {
	cases := []struct {
		m      Matrix
		coords []matrixCheck
	}{
		{
			m: Matrix{
				{1, 2, 3, 4},
				{5.5, 6.5, 7.5, 8.5},
				{9, 10, 11, 12},
				{13.5, 14.5, 15.5, 16.5},
			},
			coords: []matrixCheck{
				{X: 0, Y: 0, Value: 1},
				{X: 0, Y: 3, Value: 4},
				{X: 1, Y: 0, Value: 5.5},
				{X: 1, Y: 2, Value: 7.5},
				{X: 2, Y: 2, Value: 11},
				{X: 3, Y: 0, Value: 13.5},
				{X: 3, Y: 2, Value: 15.5},
			},
		},
		{
			m: Matrix{
				{-3, 5},
				{1, -2},
			},
			coords: []matrixCheck{
				{X: 0, Y: 0, Value: -3},
				{X: 0, Y: 1, Value: 5},
				{X: 1, Y: 0, Value: 1},
				{X: 1, Y: 1, Value: -2},
			},
		},
		{
			m: Matrix{
				{-3, 5, 0},
				{1, -2, -7},
				{0, 1, 1},
			},
			coords: []matrixCheck{
				{X: 0, Y: 0, Value: -3},
				{X: 1, Y: 1, Value: -2},
				{X: 2, Y: 2, Value: 1},
			},
		},
	}

	for _, tc := range cases {
		for _, c := range tc.coords {
			if tc.m[c.X][c.Y] != c.Value {
				t.Errorf("Coordinate [%d, %d] expected: %f, reeived: %f", c.X, c.Y, c.Value, tc.m[c.X][c.Y])
			}
		}
	}
}

func TestCopy(t *testing.T) {
	a := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	b := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	if !a.Equals(b) {
		t.Error("a != b")
	}

	c := a.Copy()

	if !c.Equals(b) {
		t.Error("c != b")
	}

	c[0][2] = 9

	if !a.Equals(b) {
		t.Error("a != b after c mod")
	}
}

func TestCompareMatrix(t *testing.T) {
	cases := []struct {
		a     Matrix
		b     Matrix
		equal bool
	}{
		{
			a: Matrix{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 8, 7, 6},
				{5, 4, 3, 2},
			},
			b: Matrix{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 8, 7, 6},
				{5, 4, 3, 2},
			},
			equal: true,
		},
		{
			a: Matrix{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 8, 7, 6},
				{5, 4, 3, 2},
			},
			b: Matrix{
				{2, 3, 4, 5},
				{6, 7, 8, 9},
				{8, 7, 6, 5},
				{4, 3, 2, 1},
			},
			equal: false,
		},
	}

	for _, tc := range cases {
		if tc.a.Equals(tc.b) != tc.equal {
			t.Errorf("Expected: %+v, received: %+v %v", tc.a, tc.b, tc.equal)
		}
	}
}

func TestMultiplyMatrix(t *testing.T) {
	cases := []struct {
		a        Matrix
		b        Matrix
		expected Matrix
	}{
		{
			a: Matrix{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 8, 7, 6},
				{5, 4, 3, 2},
			},
			b: Matrix{
				{-2, 1, 2, 3},
				{3, 2, 1, -1},
				{4, 3, 6, 5},
				{1, 2, 7, 8},
			},
			expected: Matrix{
				{20, 22, 50, 48},
				{44, 54, 114, 108},
				{40, 58, 110, 102},
				{16, 26, 46, 42},
			},
		},
		{
			a: Matrix{
				{0, 1, 2, 4},
				{1, 2, 4, 8},
				{2, 4, 8, 16},
				{4, 8, 16, 32},
			},
			b: IdentityMatrix(),
			expected: Matrix{
				{0, 1, 2, 4},
				{1, 2, 4, 8},
				{2, 4, 8, 16},
				{4, 8, 16, 32},
			},
		},
	}

	for _, tc := range cases {
		result := tc.a.Multiply(tc.b)
		if !result.Equals(tc.expected) {
			t.Errorf("Expected: %+v, received: %+v", tc.expected, result)
		}
	}
}

func TestMultiplyMatrixTuple(t *testing.T) {
	cases := []struct {
		a        Matrix
		b        Tuple
		expected Tuple
	}{
		{
			a: Matrix{
				{1, 2, 3, 4},
				{2, 4, 4, 2},
				{8, 6, 4, 1},
				{0, 0, 0, 1},
			},
			b: Tuple{
				X: 1,
				Y: 2,
				Z: 3,
				W: 1,
			},
			expected: Tuple{
				X: 18,
				Y: 24,
				Z: 33,
				W: 1,
			},
		},
		{
			a: IdentityMatrix(),
			b: Tuple{
				X: 1,
				Y: 2,
				Z: 3,
				W: 4,
			},
			expected: Tuple{
				X: 1,
				Y: 2,
				Z: 3,
				W: 4,
			},
		},
	}

	for _, tc := range cases {
		result := tc.a.MultiplyTuple(tc.b)
		if !TupleEqual(tc.expected, result) {
			t.Errorf("Expected: %+v, received: %+v", tc.expected, result)
		}
	}
}

func TestTransposeMatrix(t *testing.T) {
	cases := []struct {
		a        Matrix
		expected Matrix
	}{
		{
			a: Matrix{
				{0, 9, 3, 0},
				{9, 8, 0, 8},
				{1, 8, 5, 3},
				{0, 0, 5, 8},
			},
			expected: Matrix{
				{0, 9, 1, 0},
				{9, 8, 8, 0},
				{3, 0, 5, 5},
				{0, 8, 3, 8},
			},
		},
		{
			a:        IdentityMatrix(),
			expected: IdentityMatrix(),
		},
	}

	for _, tc := range cases {
		result := tc.a.Transpose()
		if !result.Equals(tc.expected) {
			t.Errorf("Expected: %+v, received: %+v", tc.expected, result)
		}
	}
}

func TestFindDeterminant(t *testing.T) {
	cases := []struct {
		a        Matrix
		expected float64
	}{
		{
			a: Matrix{
				{1, 5},
				{-3, 2},
			},
			expected: 17,
		},
		{
			a: Matrix{
				{1, 2, 6},
				{-5, 8, -4},
				{2, 6, 4},
			},
			expected: -196,
		},
		{
			a: Matrix{
				{-2, -8, 3, 5},
				{-3, 1, 7, 3},
				{1, 2, -9, 6},
				{-6, 7, 7, -9},
			},
			expected: -4071,
		},
	}

	for _, tc := range cases {
		result := tc.a.Determinant()
		if !FloatEqual(result, tc.expected) {
			t.Errorf("Expected: %f, received: %f", tc.expected, result)
		}
	}
}

func TestSubmatrix(t *testing.T) {
	cases := []struct {
		a        Matrix
		row      int
		column   int
		expected Matrix
	}{
		{
			a: Matrix{
				{1, 5, 0},
				{-3, 2, 7},
				{0, 6, -3},
			},
			row:    0,
			column: 2,
			expected: Matrix{
				{-3, 2},
				{0, 6},
			},
		},
		{
			a: Matrix{
				{-6, 1, 1, 6},
				{-8, 5, 8, 6},
				{-1, 0, 8, 2},
				{-7, 1, -1, 1},
			},
			row:    2,
			column: 1,
			expected: Matrix{
				{-6, 1, 6},
				{-8, 8, 6},
				{-7, -1, 1},
			},
		},
	}

	for _, tc := range cases {
		result := tc.a.Submatrix(tc.row, tc.column)
		if !result.Equals(tc.expected) {
			t.Errorf("Expected: %f, received: %f", tc.expected, result)
		}
	}
}

func TestFindMinor(t *testing.T) {
	cases := []struct {
		a        Matrix
		row      int
		column   int
		expected float64
	}{
		{
			a: Matrix{
				{3, 5, 0},
				{2, -1, -7},
				{6, -1, 5},
			},
			row:      1,
			column:   0,
			expected: 25,
		},
		{
			a: Matrix{
				{3, 5, 0},
				{2, -1, -7},
				{6, -1, 5},
			},
			row:      0,
			column:   0,
			expected: -12,
		},
	}

	for _, tc := range cases {
		result := tc.a.Minor(tc.row, tc.column)
		if !FloatEqual(result, tc.expected) {
			t.Errorf("Expected: %f, received: %f", tc.expected, result)
		}
	}
}

func TestFindCofactor(t *testing.T) {
	cases := []struct {
		a        Matrix
		row      int
		column   int
		expected float64
	}{
		{
			a: Matrix{
				{3, 5, 0},
				{2, -1, -7},
				{6, -1, 5},
			},
			row:      0,
			column:   0,
			expected: -12,
		},
		{
			a: Matrix{
				{3, 5, 0},
				{2, -1, -7},
				{6, -1, 5},
			},
			row:      1,
			column:   0,
			expected: -25,
		},
		{
			a: Matrix{
				{1, 2, 6},
				{-5, 8, -4},
				{2, 6, 4},
			},
			row:      0,
			column:   0,
			expected: 56,
		},
		{
			a: Matrix{
				{1, 2, 6},
				{-5, 8, -4},
				{2, 6, 4},
			},
			row:      0,
			column:   1,
			expected: 12,
		},
		{
			a: Matrix{
				{1, 2, 6},
				{-5, 8, -4},
				{2, 6, 4},
			},
			row:      0,
			column:   2,
			expected: -46,
		},
	}

	for _, tc := range cases {
		result := tc.a.Cofactor(tc.row, tc.column)
		if !FloatEqual(result, tc.expected) {
			t.Errorf("Expected: %f, received: %f", tc.expected, result)
		}
	}
}

func TestIsInvertible(t *testing.T) {
	cases := []struct {
		a           Matrix
		determinant float64
		invertible  bool
	}{
		{
			a: Matrix{
				{6, 4, 4, 4},
				{5, 5, 7, 6},
				{4, -9, 3, -7},
				{9, 1, 7, -6},
			},
			determinant: -2120,
			invertible:  true,
		},
		{
			a: Matrix{
				{-4, 2, -2, -3},
				{9, 6, 2, 6},
				{0, -5, 1, -5},
				{0, 0, 0, 0},
			},
			determinant: 0,
			invertible:  false,
		},
	}

	for _, tc := range cases {
		det := tc.a.Determinant()
		if !FloatEqual(det, tc.determinant) {
			t.Errorf("Bad determinant - Expected: %f, received: %f", tc.determinant, det)
		}
		result := tc.a.Invertible()
		if result != tc.invertible {
			t.Errorf("Expected: %v, received: %v", tc.invertible, result)
		}
	}
}

func TestInvert(t *testing.T) {
	cases := []struct {
		a        Matrix
		expected Matrix
	}{
		{
			a: Matrix{
				{-5, 2, 6, -8},
				{1, -5, 1, 8},
				{7, 7, -6, -7},
				{1, -3, 7, 4},
			},
			expected: Matrix{
				{0.21805, 0.45113, 0.24060, -0.04511},
				{-0.80827, -1.45677, -0.44361, 0.52068},
				{-0.07895, -0.22368, -0.05263, 0.19737},
				{-0.52256, -0.81391, -0.30075, 0.30639},
			},
		},
		{
			a: Matrix{
				{8, -5, 9, 2},
				{7, 5, 6, 1},
				{-6, 0, 9, 6},
				{-3, 0, -9, -4},
			},
			expected: Matrix{
				{-0.15385, -0.15385, -0.28205, -0.53846},
				{-0.07692, 0.12308, 0.02564, 0.03077},
				{0.35897, 0.35897, 0.43590, 0.92308},
				{-0.69231, -0.69231, -0.76923, -1.92308},
			},
		},
		{
			a: Matrix{
				{9, 3, 0, 9},
				{-5, -2, -6, -3},
				{-4, 9, 6, 4},
				{-7, 6, 6, 2},
			},
			expected: Matrix{
				{-0.04074, -0.07778, 0.14444, -0.22222},
				{-0.07778, 0.03333, 0.36667, -0.33333},
				{-0.02901, -0.14630, -0.10926, 0.12963},
				{0.17778, 0.06667, -0.26667, 0.33333},
			},
		},
	}

	for _, tc := range cases {
		result, err := tc.a.Invert()

		if err != nil {
			t.Errorf("Encountered error %s", err)
		}

		if !tc.expected.Equals(result) {
			t.Errorf("Expected: %+v, received: %+v", tc.expected, result)
		}
	}
}

func TestMultiplyInverse(t *testing.T) {
	cases := []struct {
		a Matrix
		b Matrix
	}{
		{
			a: Matrix{
				{3, -9, 7, 3},
				{3, -8, 2, -9},
				{-4, 4, 4, 1},
				{-6, 5, -1, 1},
			},
			b: Matrix{
				{8, 2, 2, 2},
				{3, -1, 7, 0},
				{7, 0, 5, 4},
				{6, -2, 0, 5},
			},
		},
	}

	for _, tc := range cases {
		c := tc.a.Multiply(tc.b)

		i, _ := tc.b.Invert()
		d := c.Multiply(i)

		if !d.Equals(tc.a) {
			t.Errorf("Expected: %+v, received: %+v", tc.a, d)
		}
	}
}
