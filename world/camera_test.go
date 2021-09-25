package world

import (
	"math"
	"testing"

	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
)

func TestCamera(t *testing.T) {
	hsize := 160
	vsize := 120
	fov := math.Pi / 2

	c := Camera(hsize, vsize, fov)

	if c.HSize != hsize {
		t.Errorf("hsize mismatch expected %d received %d", hsize, c.HSize)
	}

	if c.VSize != vsize {
		t.Errorf("vsize mismatch expected %d received %d", vsize, c.VSize)
	}

	if !data.FloatEqual(c.FieldOfView, fov) {
		t.Errorf("FieldOfView mismatch expected %f received %f", fov, c.FieldOfView)
	}

	if !c.Transform.Equals(data.IdentityMatrix()) {
		t.Errorf("transform mismatch expected %+v received %+v", data.IdentityMatrix(), c.Transform)
	}
}

func TestPixelSize(t *testing.T) {
	testCases := []struct {
		c        *CameraType
		expected float64
	}{
		{
			c:        Camera(200, 125, math.Pi/2),
			expected: 0.01,
		},
		{
			c:        Camera(125, 200, math.Pi/2),
			expected: 0.01,
		},
	}

	for _, tc := range testCases {
		if !data.FloatEqual(tc.c.PixelSize, tc.expected) {
			t.Errorf("FieldOfView mismatch expected %f received %f", tc.expected, tc.c.PixelSize)
		}
	}
}

func TestRayForPixel(t *testing.T) {
	testCases := []struct {
		c        *CameraType
		x        int
		y        int
		expected data.RayType
	}{
		{
			c:        Camera(201, 101, math.Pi/2),
			x:        100,
			y:        50,
			expected: data.Ray(data.Point(0, 0, 0), data.Vector(0, 0, -1)),
		},
		{
			c:        Camera(201, 101, math.Pi/2),
			x:        0,
			y:        0,
			expected: data.Ray(data.Point(0, 0, 0), data.Vector(0.66519, 0.33259, -0.66851)),
		},
		{
			c: func() *CameraType {
				c := Camera(201, 101, math.Pi/2)
				c.Transform = data.IdentityMatrix().Translate(0, -2, 5).RotateY(math.Pi / 4)
				return c
			}(),
			x:        100,
			y:        50,
			expected: data.Ray(data.Point(0, 2, -5), data.Vector(math.Sqrt(2)/2, 0, math.Sqrt(2)/2*-1)),
		},
	}

	for _, tc := range testCases {
		ray := tc.c.RayForPixel(tc.x, tc.y)
		if !data.TupleEqual(ray.Direction, tc.expected.Direction) {
			t.Errorf("Ray Direction mismatch expected %f received %f", tc.expected.Direction, ray.Direction)
		}

		if !data.TupleEqual(ray.Origin, tc.expected.Origin) {
			t.Errorf("Ray Origin mismatch expected %f received %f", tc.expected.Origin, ray.Origin)
		}
	}
}

func TestRender(t *testing.T) {

	w := DefaultWorld()
	c := Camera(11, 11, math.Pi/2)
	c.Transform = data.ViewTransform(
		data.Point(0, 0, -5),
		data.Point(0, 0, 0),
		data.Vector(0, 1, 0),
	)
	expected := material.Colour(0.38066, 0.47583, 0.2855)

	pixel := c.Render(w).Pixel(5, 5)
	if !material.ColourEqual(expected, pixel) {
		t.Errorf("Ray Direction mismatch expected %f received %f", expected, pixel)
	}
}
