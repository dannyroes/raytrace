package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"strconv"
	"strings"
)

type CanvasType struct {
	Width  int
	Height int
	Pixels [][]ColourTuple
}

func Canvas(width, height int) CanvasType {
	c := CanvasType{Width: width, Height: height}

	pixels := make([][]ColourTuple, width)
	for x := 0; x < width; x++ {
		column := make([]ColourTuple, height)
		pixels[x] = column
	}

	c.Pixels = pixels
	return c
}

func (c CanvasType) Pixel(x, y int) ColourTuple {
	return c.Pixels[x][y]
}

func (c CanvasType) WritePixel(x, y int, colour ColourTuple) {
	c.Pixels[x][y] = colour
}

func (c CanvasType) ToPPM() string {
	var pixels strings.Builder
	var line strings.Builder

	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			writePixelValue(&pixels, &line, strconv.Itoa(getCappedColour(c.Pixel(x, y).Red(), 255)))
			writePixelValue(&pixels, &line, strconv.Itoa(getCappedColour(c.Pixel(x, y).Green(), 255)))
			writePixelValue(&pixels, &line, strconv.Itoa(getCappedColour(c.Pixel(x, y).Blue(), 255)))
		}
		line.WriteString("\n")
		pixels.WriteString(line.String())
		line.Reset()
	}

	return fmt.Sprintf("P3\n%d %d\n255\n%s", c.Width, c.Height, pixels.String())
}

func (c CanvasType) ToImage() image.Image {
	im := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))

	for x := 0; x < c.Width; x++ {
		for y := 0; y < c.Height; y++ {
			im.Set(x, y, c.Pixel(x, y))
		}
	}

	return im
}

func (c CanvasType) ToPNG(filename string) error {
	im := c.ToImage()
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	png.Encode(f, im)
	return nil
}

func getCappedColour(ratio float64, cap int) int {
	colour := int(math.Ceil(ratio * float64(cap)))

	if colour < 0 {
		return 0
	}

	if colour > cap {
		return cap
	}

	return colour
}

func writePixelValue(pixels, line *strings.Builder, value string) {
	spacer := " "
	if line.Len() == 0 {
		spacer = ""
	} else if (line.Len() + len(value)) >= 70 {
		spacer = ""
		line.WriteString("\n")
		pixels.WriteString(line.String())
		line.Reset()
	}
	fmt.Fprintf(line, "%s%s", spacer, value)
}
