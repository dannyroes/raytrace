package main

import "math"

type CameraType struct {
	HSize       int
	VSize       int
	FieldOfView float64
	Transform   Matrix
	PixelSize   float64
	halfWidth   float64
	halfHeight  float64
}

func Camera(hsize, vsize int, fieldOfView float64) *CameraType {
	c := &CameraType{HSize: hsize, VSize: vsize, FieldOfView: fieldOfView, Transform: IdentityMatrix()}
	c.CalcPixelSize()
	return c
}

func (c *CameraType) CalcPixelSize() {
	halfView := math.Tan(c.FieldOfView / 2)
	aspect := float64(c.HSize) / float64(c.VSize)

	if aspect >= 1 {
		c.halfWidth = halfView
		c.halfHeight = halfView / aspect
	} else {
		c.halfHeight = halfView
		c.halfWidth = halfView * aspect
	}

	c.PixelSize = (c.halfWidth * 2) / float64(c.HSize)
}

func (c *CameraType) RayForPixel(x, y int) RayType {
	xOffset := (float64(x) + 0.5) * c.PixelSize
	yOffset := (float64(y) + 0.5) * c.PixelSize

	worldX := c.halfWidth - xOffset
	worldY := c.halfHeight - yOffset

	pixel := c.Transform.Invert().MultiplyTuple(Point(worldX, worldY, -1))
	origin := c.Transform.Invert().MultiplyTuple(Point(0, 0, 0))

	dir := pixel.Sub(origin).Normalize()

	return Ray(origin, dir)
}

func (c *CameraType) Render(w WorldType) CanvasType {
	image := Canvas(c.HSize, c.VSize)

	for y := 0; y < c.VSize; y++ {
		for x := 0; x < c.HSize; x++ {
			ray := c.RayForPixel(x, y)
			colour := w.ColourAt(ray)
			image.WritePixel(x, y, colour)
		}
	}
	return image
}
