package main

import (
	"fmt"
	"io/ioutil"
	"math"
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
	drawCircleSingleOrigin()
}

func drawCircle() {
	canvas := Canvas(500, 500)
	colour := Colour(1, 0.5, 0.3)
	sphere := Sphere(1).SetTransform(Scaling(200, 100, 200).RotateZ(2.5).Translate(250, 250, 250))

	for x := 0; x < canvas.Width; x++ {
		for y := 0; y < canvas.Height; y++ {
			ray := Ray(Point(float64(x), float64(y), -500), Vector(0, 0, 1))
			h := sphere.Intersects(ray).Hit()

			if h.Object.Id == 1 {
				canvas.WritePixel(x, y, colour)
			}
		}
	}

	ioutil.WriteFile("proj.ppm", []byte(canvas.ToPPM()), 0755)
}

func drawCircleSingleOrigin() {
	canvas := Canvas(500, 500)
	wallZ := 200.0
	rayOrigin := Point(0, 0, -500)
	colour := Colour(1, 0.5, 0.3)
	sphere := Sphere(1).SetTransform(Scaling(100, 75, 100).RotateZ(2.5))

	for x := -250; x < 250; x++ {
		for y := -250; y < 250; y++ {
			v := Point(float64(x), float64(y), wallZ).Sub(rayOrigin).Normalize()
			ray := Ray(rayOrigin, v)
			h := sphere.Intersects(ray).Hit()

			if h.Object.Id == 1 {
				canvas.WritePixel(x+250, y+250, colour)
			}
		}
	}

	ioutil.WriteFile("proj.ppm", []byte(canvas.ToPPM()), 0755)
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
