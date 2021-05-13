package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCreateCanvas(t *testing.T) {
	c := Canvas(10, 20)

	if c.Width != 10 {
		t.Errorf("Expected %+v, received %+v", 10, c.Width)
	}

	if c.Height != 20 {
		t.Errorf("Expected %+v, received %+v", 20, c.Height)
	}

	for x := 0; x < c.Width; x++ {
		for y := 0; y < c.Height; y++ {
			if !ColourEqual(c.Pixel(x, y), Colour(0, 0, 0)) {
				t.Errorf("Expected %+v, received %+v", Colour(0, 0, 0), c.Pixel(x, y))
			}
		}
	}
}

func TestWritePixel(t *testing.T) {
	c := Canvas(10, 20)
	red := Colour(1, 0, 0)

	c.WritePixel(2, 3, red)

	if !ColourEqual(c.Pixel(2, 3), red) {
		t.Errorf("Expected %+v, received %+v", red, c.Pixel(2, 3))
	}
}

func TestCanvasToPPMHeader(t *testing.T) {
	c := Canvas(5, 3)

	p := c.ToPPM()

	lines := strings.Split(p, "\n")
	if lines[1] != "5 3" {
		t.Errorf("Expected '5 3', received '%s'", lines[1])
	}
}

func TestCanvasToPPMPixels(t *testing.T) {
	c := Canvas(5, 3)

	c1 := Colour(1.5, 0, 0)
	c2 := Colour(0, 0.5, 0)
	c3 := Colour(-0.5, 0, 1)

	c.WritePixel(0, 0, c1)
	c.WritePixel(2, 1, c2)
	c.WritePixel(4, 2, c3)

	p := c.ToPPM()
	expected := `255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`

	lines := strings.Split(p, "\n")
	lines = lines[3:]
	pixels := strings.Join(lines, "\n")

	if pixels != expected {
		t.Errorf("Expected '%s', received '%s'", expected, pixels)
	}
}

func TestCanvasToPPMLineWrap(t *testing.T) {
	c := Canvas(10, 2)

	c1 := Colour(1, 0.8, 0.6)

	for x := 0; x < 10; x++ {
		for y := 0; y < 2; y++ {
			c.WritePixel(x, y, c1)
		}
	}

	p := c.ToPPM()
	expected := `255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
`
	lines := strings.Split(p, "\n")
	lines = lines[3:]
	pixels := strings.Join(lines, "\n")
	fmt.Println(expected, len(expected))
	fmt.Println(pixels, len(pixels))
	if pixels != expected {
		t.Errorf("Expected '%s', received '%s'", expected, pixels)
	}
}

func TestCanvasTrailingNewline(t *testing.T) {
	c := Canvas(5, 3)

	p := c.ToPPM()

	lines := strings.Split(p, "\n")
	if lines[len(lines)-1] != "" {
		t.Error("Expected trailing newline")
	}
}
