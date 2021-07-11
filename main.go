package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"time"
)

type Environment struct {
	Gravity Tuple
	Wind    Tuple
}

type Projectile struct {
	Position Tuple
	Velocity Tuple
}

func main() {
	drawScene()
	fmt.Println("Done!")
}

func drawScene() {
	floor := Sphere(1)
	floor.Transform = Scaling(10, 0.01, 10)
	m := Material()
	m.Colour = Colour(1, 0.9, 0.9)
	m.Specular = 0
	floor.SetMaterial(m)

	leftWall := Sphere(2)
	leftWall.Transform = Scaling(10, 0.01, 10).RotateX(math.Pi/2).RotateY(math.Pi/4*-1).Translate(0, 0, 5)
	leftWall.SetMaterial(m)

	rightWall := Sphere(3)
	rightWall.Transform = Scaling(10, 0.01, 10).RotateX(math.Pi/2).RotateY(math.Pi/4).Translate(0, 0, 5)
	rightWall.SetMaterial(m)

	middle := Sphere(4)
	middle.Transform = Translation(-0.5, 1, 0.5)

	m = Material()
	m.Colour = Colour(0.1, 1, 0.5)
	m.Diffuse = 0.7
	m.Specular = 0.3

	middle.SetMaterial(m)

	right := Sphere(5)
	right.Transform = IdentityMatrix().Scale(0.5, 0.5, 0.5).Translate(1.5, 0.5, -0.5)

	m = Material()
	m.Colour = Colour(0.5, 1, 0.1)
	m.Diffuse = 0.7
	m.Specular = 0.3

	right.SetMaterial(m)

	left := Sphere(6)
	left.Transform = IdentityMatrix().Scale(0.33, 0.33, 0.33).Translate(-1.5, 0.33, -0.75)

	m = Material()
	m.Colour = Colour(1, 0.8, 0.1)
	m.Diffuse = 0.7
	m.Specular = 0.3

	left.SetMaterial(m)

	w := World()
	w.Light = PointLight(Point(-10, 10, -10), Colour(1, 1, 1))

	c := Camera(1000, 500, math.Pi/3)
	c.Transform = ViewTransform(Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0))

	w.Objects = []Object{
		floor,
		leftWall,
		rightWall,
		middle,
		left,
		right,
	}
	image := c.Render(w)
	ioutil.WriteFile("output/scene.ppm", []byte(image.ToPPM()), 0755)
	err := image.ToPNG("output/scene.png")
	if err != nil {
		fmt.Println(err)
	}
}

func drawCircle() {
	canvas := Canvas(500, 500)
	colour := Colour(1, 0.5, 0.3)
	sphere := Sphere(1)
	sphere.SetTransform(Scaling(200, 100, 200).RotateZ(2.5).Translate(250, 250, 250))

	for x := 0; x < canvas.Width; x++ {
		for y := 0; y < canvas.Height; y++ {
			ray := Ray(Point(float64(x), float64(y), -500), Vector(0, 0, 1))
			h := sphere.Intersects(ray).Hit()

			if h.Object.GetId() == 1 {
				canvas.WritePixel(x, y, colour)
			}
		}
	}

	ioutil.WriteFile("proj.ppm", []byte(canvas.ToPPM()), 0755)
}

func drawCircleSingleOrigin() {
	canvas := Canvas(500, 500)
	wallZ := 200.0
	rayOrigin := Point(0, 0, -20)
	sphere := Sphere(1)
	sphere.SetTransform(IdentityMatrix().Scale(10, 10, 10))
	sphere.Material.Colour = Colour(1, 0.2, 1)
	sphere.Material.Shininess = 50

	light := PointLight(Point(-10, 10, -10), Colour(1, 1, 1))

	start := time.Now()

	for x := -250; x < 250; x++ {
		for y := -250; y < 250; y++ {
			v := Point(float64(x), float64(y), wallZ).Sub(rayOrigin).Normalize()
			ray := Ray(rayOrigin, v)
			h := sphere.Intersects(ray).Hit()

			if h.Object.GetId() > 0 {
				p := ray.Position(h.T)
				normal := h.Object.NormalAt(p)
				eye := ray.Direction.Neg()
				colour := Lighting(h.Object.GetMaterial(), light, p, eye, normal, false)
				canvas.WritePixel(x+250, y+250, colour)
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed)

	ioutil.WriteFile("output/proj.ppm", []byte(canvas.ToPPM()), 0755)
}

func drawClock() {
	canvas := Canvas(500, 500)

	rad := math.Pi * 2

	start := Point(0, 225, 0)
	colour := Colour(1, 1, 1)

	for x := 0; x < 12; x++ {
		point := IdentityMatrix().RotateZ(rad/12*float64(x)).Translate(250, 250, 0).MultiplyTuple(start)

		fmt.Println(point)
		canvas.WritePixel(int(point.X), int(point.Y), colour)
	}

	ioutil.WriteFile("proj.ppm", []byte(canvas.ToPPM()), 0755)
}

func runSim() {
	var canvasHeight, canvasWidth int
	canvasHeight = 550
	canvasWidth = 900
	canvas := Canvas(canvasWidth, canvasHeight)

	projColour := Colour(1, 0, 0)
	e := Environment{Gravity: Vector(0, -0.08, 0), Wind: Vector(-0.05, 0, 0)}

	p := Projectile{Position: Point(0, 1, 0), Velocity: Vector(1.4, 1.8, 0).Normalize().Mul(11.25)}
	ticks := 0
	for {
		posX := int(p.Position.X)
		posY := canvasHeight - int(p.Position.Y)

		if posX > canvasWidth || posX < 0 || posY > canvasHeight || posY < 0 {
			break
		}
		canvas.WritePixel(posX, posY, projColour)

		if p.Position.Y < 0.0 || FloatEqual(p.Position.Y, 0) {
			break
		}
		ticks++
		p = tick(e, p)
	}

	ioutil.WriteFile("proj.ppm", []byte(canvas.ToPPM()), 0755)

	fmt.Printf("Completed after %d ticks", ticks)
}

func tick(e Environment, p Projectile) Projectile {
	pos := p.Position.Add(p.Velocity)
	vel := p.Velocity.Add(e.Gravity).Add(e.Wind)
	return Projectile{Position: pos, Velocity: vel}
}
