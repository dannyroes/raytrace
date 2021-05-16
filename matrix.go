package main

import (
	"errors"
)

type Matrix [][]float64

var IdentityMatrix = Matrix{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 1, 0},
	{0, 0, 0, 1},
}

func (m Matrix) Equals(b Matrix) bool {
	if len(m) != len(b) {
		return false
	}
	for x, row := range m {
		if len(row) != len(b[x]) {
			return false
		}
		for y, val := range row {
			if !FloatEqual(val, b[x][y]) {
				return false
			}
		}
	}

	return true
}

func (m Matrix) Multiply(b Matrix) Matrix {
	r := make(Matrix, 4)

	for x := 0; x < 4; x++ {
		r[x] = make([]float64, 4)
		for y := 0; y < 4; y++ {
			r[x][y] = m[x][0]*b[0][y] +
				m[x][1]*b[1][y] +
				m[x][2]*b[2][y] +
				m[x][3]*b[3][y]
		}
	}
	return r
}

func (m Matrix) MultiplyTuple(t Tuple) Tuple {
	r := Tuple{
		X: Dot(m.rowToTuple(0), t),
		Y: Dot(m.rowToTuple(1), t),
		Z: Dot(m.rowToTuple(2), t),
		W: Dot(m.rowToTuple(3), t),
	}
	return r
}

func (m Matrix) Transpose() Matrix {
	result := make(Matrix, len(m))

	for x := 0; x < len(m); x++ {
		result[x] = make([]float64, len(m))
	}

	for x, row := range m {
		for y, val := range row {
			result[y][x] = val
		}
	}

	return result
}

func (m Matrix) Determinant() float64 {
	if len(m) == 2 {
		return m[0][0]*m[1][1] - m[1][0]*m[0][1]
	}

	var d float64 = 0.0

	for y := 0; y < len(m); y++ {
		d += m[0][y] * m.Cofactor(0, y)
	}

	return d
}

func (m Matrix) Submatrix(row, column int) Matrix {
	result := make(Matrix, len(m))
	copy(result, m)
	result = append(result[:row], result[row+1:]...)

	for x, row := range result {
		newRow := make([]float64, len(row))
		copy(newRow, row)
		result[x] = append(newRow[:column], newRow[column+1:]...)
	}

	return result
}

func (m Matrix) Minor(row, column int) float64 {
	sub := m.Submatrix(row, column)

	return sub.Determinant()
}

func (m Matrix) Cofactor(row, column int) float64 {
	minor := m.Minor(row, column)
	factor := 1.0

	if (row+column)%2 == 1 {
		factor = -1.0
	}

	return factor * minor
}

func (m Matrix) Invertible() bool {
	return !FloatEqual(0.0, m.Determinant())
}

func (m Matrix) Invert() (Matrix, error) {
	if !m.Invertible() {
		return nil, errors.New("matrix is not invertible")
	}

	result := make(Matrix, len(m))
	for x := range m {
		result[x] = make([]float64, len(m))
	}

	det := m.Determinant()
	for x, row := range m {
		for y := range row {
			// Reverse x and y to invert
			result[y][x] = m.Cofactor(x, y) / det
		}
	}

	return result, nil
}

func (m Matrix) rowToTuple(x int) Tuple {
	return Tuple{
		X: m[x][0],
		Y: m[x][1],
		Z: m[x][2],
		W: m[x][3],
	}
}
