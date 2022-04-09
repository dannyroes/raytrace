package main

import (
	"fmt"
	"math"

	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
	"github.com/dannyroes/raytrace/shape"
	"github.com/dannyroes/raytrace/world"
)

func hexCorner(m material.MaterialType) shape.Shape {
	c := shape.Sphere()
	c.SetTransform(data.IdentityMatrix().Scale(0.25, 0.25, 0.25).Translate(0, 0, -1))
	c.SetMaterial(m)
	return c
}

func hexEdge(m material.MaterialType) shape.Shape {
	c := shape.Cylinder()
	c.Minimum = 0
	c.Maximum = 1
	c.SetTransform(data.IdentityMatrix().Scale(0.25, 1, 0.25).RotateZ(-math.Pi/2).RotateY(-math.Pi/6).Translate(0, 0, -1))
	c.SetMaterial(m)
	return c
}

func hexSide(m material.MaterialType) shape.Shape {
	side := shape.Group()
	side.AddChild(hexCorner(m))
	side.AddChild(hexEdge(m))

	return side
}

func hex(m material.MaterialType) shape.Shape {
	hex := shape.Group()

	for x := 0; x < 6; x++ {
		side := hexSide(m)
		side.SetTransform(data.IdentityMatrix().RotateY(float64(x) * math.Pi / 3))
		hex.AddChild(side)
	}

	return hex
}

func hexScene() {
	w := world.World()

	w.Lights = append(w.Lights, world.PointLight(data.Point(2, 4, 2), material.Colour(1, 1, 1)))

	c := world.Camera(400, 300, math.Pi/3)

	c.Transform = data.ViewTransform(data.Point(2, 2, 2), data.Point(0, 0.5, 0), data.Vector(0, 1, 0))

	m := material.Material()
	m.Colour = material.Colour(0.9, 0.3, 0.1)
	m.Reflective = 0.7

	h1 := hex(m)

	m.Colour = material.Colour(0.1, 0.3, 0.8)
	h2 := hex(m)
	h2.SetTransform(h2.GetTransform().RotateZ(math.Pi/2).Translate(-2, 0.5, 0))

	m.Colour = material.Colour(0.1, 0.8, 0.1)
	h3 := hex(m)
	h3.SetTransform(h3.GetTransform().RotateX(math.Pi/2).Translate(0, 0.5, -2))

	w.Objects = []shape.Shape{
		h1, h2, h3,
	}
	image := c.Render(w)

	err := image.ToPNG("output/scene.png")
	if err != nil {
		fmt.Println(err)
	}
}
