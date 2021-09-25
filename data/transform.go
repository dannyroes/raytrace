package data

import (
	"math"
)

func Translation(x, y, z float64) Matrix {
	m := IdentityMatrix()

	m[0][3] = x
	m[1][3] = y
	m[2][3] = z

	return m
}

func Scaling(x, y, z float64) Matrix {
	m := IdentityMatrix()

	m[0][0] = x
	m[1][1] = y
	m[2][2] = z

	return m
}

func RotateX(r float64) Matrix {
	m := IdentityMatrix()

	m[1][1] = math.Cos(r)
	m[1][2] = math.Sin(r) * -1
	m[2][1] = math.Sin(r)
	m[2][2] = math.Cos(r)

	return m
}

func RotateY(r float64) Matrix {
	m := IdentityMatrix()

	m[0][0] = math.Cos(r)
	m[0][2] = math.Sin(r)
	m[2][0] = math.Sin(r) * -1
	m[2][2] = math.Cos(r)

	return m
}

func RotateZ(r float64) Matrix {
	m := IdentityMatrix()

	m[0][0] = math.Cos(r)
	m[0][1] = math.Sin(r) * -1
	m[1][0] = math.Sin(r)
	m[1][1] = math.Cos(r)

	return m
}

func Shear(xy, xz, yx, yz, zx, zy float64) Matrix {
	m := IdentityMatrix()

	m[0][1] = xy
	m[0][2] = xz
	m[1][0] = yx
	m[1][2] = yz
	m[2][0] = zx
	m[2][1] = zy

	return m
}

func ViewTransform(from, to, up Tuple) Matrix {
	forward := to.Sub(from).Normalize()
	left := Cross(forward, up.Normalize())
	trueUp := Cross(left, forward)

	orientation := Matrix{
		{left.X, left.Y, left.Z, 0},
		{trueUp.X, trueUp.Y, trueUp.Z, 0},
		{-1 * forward.X, -1 * forward.Y, -1 * forward.Z, 0},
		{0, 0, 0, 1},
	}

	return orientation.Multiply(Translation(-1*from.X, -1*from.Y, -1*from.Z))
}
