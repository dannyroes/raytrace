package main

import (
	"errors"
)

type Matrix [][]float64

func IdentityMatrix() Matrix {
	return Matrix{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func (m Matrix) Copy() Matrix {
	c := make(Matrix, len(m))

	for x, row := range m {
		newRow := make([]float64, len(row))
		copy(newRow, row)
		c[x] = newRow
	}

	return c
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
	r := m.Copy()

	for x := 0; x < 4; x++ {
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
	result := m.Copy()

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
	result := m.Copy()
	result = append(result[:row], result[row+1:]...)

	for x, row := range result {
		result[x] = append(row[:column], row[column+1:]...)
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

	result := m.Copy()

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

func (m Matrix) Translate(x, y, z float64) Matrix {
	return Translation(x, y, z).Multiply(m)
}

func (m Matrix) Scale(x, y, z float64) Matrix {
	return Scaling(x, y, z).Multiply(m)
}

func (m Matrix) RotateX(r float64) Matrix {
	return RotateX(r).Multiply(m)
}

func (m Matrix) RotateY(r float64) Matrix {
	return RotateY(r).Multiply(m)
}

func (m Matrix) RotateZ(r float64) Matrix {
	return RotateZ(r).Multiply(m)
}

func (m Matrix) Shear(xy, xz, yx, yz, zx, zy float64) Matrix {
	return Shear(xy, xz, yx, yz, zx, zy).Multiply(m)
}
