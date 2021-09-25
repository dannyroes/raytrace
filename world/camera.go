package world

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"

	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
)

const MaxReflect int = 5

type CameraType struct {
	HSize       int
	VSize       int
	Supersample int
	FieldOfView float64
	Transform   data.Matrix
	PixelSize   float64
	halfWidth   float64
	halfHeight  float64
}

func Camera(hsize, vsize int, fieldOfView float64) *CameraType {
	c := &CameraType{HSize: hsize, VSize: vsize, FieldOfView: fieldOfView, Transform: data.IdentityMatrix()}
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

func (c *CameraType) RayForPixel(x, y int) data.RayType {
	xOffset := (float64(x) + 0.5) * c.PixelSize
	yOffset := (float64(y) + 0.5) * c.PixelSize

	worldX := c.halfWidth - xOffset
	worldY := c.halfHeight - yOffset

	pixel := c.Transform.Invert().MultiplyTuple(data.Point(worldX, worldY, -1))
	origin := c.Transform.Invert().MultiplyTuple(data.Point(0, 0, 0))

	dir := pixel.Sub(origin).Normalize()

	return data.Ray(origin, dir)
}

func (c *CameraType) Render(w WorldType) CanvasType {
	var image CanvasType

	if c.Supersample > 1 {
		c.HSize = c.HSize * c.Supersample
		c.VSize = c.VSize * c.Supersample
		c.CalcPixelSize()
	}
	fmt.Printf("Rendering width %d; height %d\n", c.HSize, c.VSize)
	image = Canvas(c.HSize, c.VSize)

	in := make(chan PixelJob)
	out := make(chan PixelColour)

	wg := &sync.WaitGroup{}

	for x := 0; x < runtime.NumCPU()-1; x++ {
		wg.Add(1)
		go renderPixel(in, out, wg)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	fmt.Println("Beginning render")

	p := 0
	t := time.Now()

	go func() {
		for y := 0; y < c.VSize; y++ {
			for x := 0; x < c.HSize; x++ {
				in <- PixelJob{x, y, c, w}
			}
		}
		close(in)
	}()

	for pixel := range out {
		image.WritePixel(pixel.x, pixel.y, pixel.c)
		p++
	}

	duration := time.Since(t)
	fmt.Printf("Rendered %d pixels in %v\n", p, duration)

	if c.Supersample > 1 {
		c.HSize = c.HSize / c.Supersample
		c.VSize = c.VSize / c.Supersample
		fmt.Printf("Downsampling to %dx%d\n", c.HSize, c.VSize)
		image = downsample(image, c.HSize, c.VSize)
	}
	return image
}

func downsample(image CanvasType, width, height int) CanvasType {
	factor := image.Width / width

	newImage := Canvas(width, height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := material.Colour(0, 0, 0)
			for iy := 0; iy < factor; iy++ {
				for ix := 0; ix < factor; ix++ {
					c = c.Add(image.Pixel(x*factor+ix, y*factor+iy))
				}
			}

			c = c.Div(float64(factor * factor))
			newImage.WritePixel(x, y, c)
		}
	}

	return newImage
}

type PixelJob struct {
	x int
	y int
	c *CameraType
	w WorldType
}

type PixelColour struct {
	x int
	y int
	c material.ColourTuple
}

func renderPixel(c <-chan PixelJob, out chan<- PixelColour, wg *sync.WaitGroup) {
	for p := range c {
		ray := p.c.RayForPixel(p.x, p.y)
		colour := p.w.ColourAt(ray, MaxReflect)
		out <- PixelColour{p.x, p.y, colour}
	}

	wg.Done()
}
