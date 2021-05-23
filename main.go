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
