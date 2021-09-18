package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
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
	width := 500
	height := 250
	supersample := 1

	fmt.Println(os.Args)

	if len(os.Args) == 3 {
		w, err := strconv.Atoi(os.Args[1])
		if err == nil {
			h, err := strconv.Atoi(os.Args[2])
			if err == nil {
				width = w
				height = h
			}
		}
	}

	drawScene(width, height, supersample)
	fmt.Println("Done!")
}

func drawScene(width, height, supersample int) {
	floor := Plane()
	m := Material()
	m.Pattern = CheckersPattern(Colour(1, 0.9, 0.9), Colour(1, 0.1, 0.1))
	m.Specular = 0
	m.Reflective = 0.6
	floor.SetMaterial(m)

	ceiling := Plane()
	ceiling.Transform = Translation(0, 10, 0)
	m = Material()
	m.Pattern = GradientPattern(Colour(0.9, 0.9, 0.9), Colour(0.1, 1, 0.1))
	m.Pattern.SetTransform(Scaling(20, 20, 20).Translate(-10, -10, -10))
	m.Specular = 0
	ceiling.SetMaterial(m)

	m = Material()
	m.Pattern = StripePattern(Colour(1, 0.9, 0.9), Colour(0.1, 0.1, 1))
	m.Specular = 0
	m.Pattern.SetTransform(RotateY(2))
	leftWall := Plane()
	leftWall.Transform = RotateX(math.Pi/2).RotateY(math.Pi/4*-1).Translate(0, 0, 5)
	leftWall.SetMaterial(m)

	rightWall := Plane()
	rightWall.Transform = RotateX(math.Pi/2).RotateY(math.Pi/4).Translate(0, 0, 5)
	rightWall.SetMaterial(m)

	rightBackWall := Plane()
	rightBackWall.Transform = RotateX(math.Pi/2).RotateY(math.Pi/4).Translate(0, 0, -15)
	rightBackWall.SetMaterial(m)

	leftBackWall := Plane()
	leftBackWall.Transform = RotateX(math.Pi/2).RotateY(math.Pi/4*-1).Translate(0, 0, -15)
	leftBackWall.SetMaterial(m)

	middle := Sphere()
	middle.Transform = Translation(-0.1, 1, -0.5)

	m = Material()
	m.Colour = Colour(0.5, 0.5, 0.5)
	m.Diffuse = 0.3
	m.Specular = 1
	m.Ambient = 0.05
	m.Shininess = 300
	m.Reflective = 0.9
	m.Transparency = 0.9
	m.RefractiveIndex = 1.52

	middle.SetMaterial(m)

	right := Sphere()
	right.Transform = IdentityMatrix().Scale(0.5, 0.5, 0.5).Translate(1.5, 0.5, -0.5)

	m = Material()
	m.Colour = Colour(0.5, 1, 0.1)
	m.Diffuse = 0.7
	m.Specular = 0.3
	m.Reflective = 0.1

	right.SetMaterial(m)

	left := Sphere()
	left.Transform = IdentityMatrix().Scale(0.33, 0.33, 0.33).Translate(-1.5, 0.33, -0.75)

	m = Material()
	m.Colour = Colour(1, 0.8, 0.1)
	m.Diffuse = 0.7
	m.Specular = 0.3
	m.Reflective = 0.1

	left.SetMaterial(m)

	w := World()
	w.Light = PointLight(Point(-4, 4, -4), Colour(1, 1, 1))

	c := Camera(width, height, math.Pi/3)
	c.Supersample = supersample
	c.Transform = ViewTransform(Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0))

	w.Objects = []Shape{
		floor,
		ceiling,
		leftWall,
		rightWall,
		leftBackWall,
		rightBackWall,
		middle,
		left,
		right,
	}
	image := c.Render(w)
	// ioutil.WriteFile("output/scene.ppm", []byte(image.ToPPM()), 0755)
	err := image.ToPNG("output/scene.png")
	if err != nil {
		fmt.Println(err)
	}
}

// func drawCircle() {
// 	canvas := Canvas(500, 500)
// 	colour := Colour(1, 0.5, 0.3)
// 	sphere := Sphere()
// 	sphere.SetTransform(Scaling(200, 100, 200).RotateZ(2.5).Translate(250, 250, 250))

// 	for x := 0; x < canvas.Width; x++ {
// 		for y := 0; y < canvas.Height; y++ {
// 			ray := Ray(Point(float64(x), float64(y), -500), Vector(0, 0, 1))
// 			h := Intersects(sphere, ray).Hit()

// 			if h.Object == sphere {
// 				canvas.WritePixel(x, y, colour)
// 			}
// 		}
// 	}

// 	ioutil.WriteFile("proj.ppm", []byte(canvas.ToPPM()), 0755)
// }

// func drawCircleSingleOrigin() {
// 	canvas := Canvas(500, 500)
// 	wallZ := 200.0
// 	rayOrigin := Point(0, 0, -20)
// 	sphere := Sphere()
// 	sphere.SetTransform(IdentityMatrix().Scale(10, 10, 10))
// 	sphere.Material.Colour = Colour(1, 0.2, 1)
// 	sphere.Material.Shininess = 50

// 	light := PointLight(Point(-10, 10, -10), Colour(1, 1, 1))

// 	start := time.Now()

// 	for x := -250; x < 250; x++ {
// 		for y := -250; y < 250; y++ {
// 			v := Point(float64(x), float64(y), wallZ).Sub(rayOrigin).Normalize()
// 			ray := Ray(rayOrigin, v)
// 			h := Intersects(sphere, ray).Hit()

// 			if h.T >= 0 {
// 				p := ray.Position(h.T)
// 				normal := NormalAt(h.Object, p)
// 				eye := ray.Direction.Neg()
// 				colour := Lighting(h.Object.GetMaterial(), light, p, eye, normal, false)
// 				canvas.WritePixel(x+250, y+250, colour)
// 			}
// 		}
// 	}

// 	elapsed := time.Since(start)
// 	fmt.Println(elapsed)

// 	ioutil.WriteFile("output/proj.ppm", []byte(canvas.ToPPM()), 0755)
// }

// func drawClock() {
// 	canvas := Canvas(500, 500)

// 	rad := math.Pi * 2

// 	start := Point(0, 225, 0)
// 	colour := Colour(1, 1, 1)

// 	for x := 0; x < 12; x++ {
// 		point := IdentityMatrix().RotateZ(rad/12*float64(x)).Translate(250, 250, 0).MultiplyTuple(start)

// 		fmt.Println(point)
// 		canvas.WritePixel(int(point.X), int(point.Y), colour)
// 	}

// 	ioutil.WriteFile("proj.ppm", []byte(canvas.ToPPM()), 0755)
// }

// func runSim() {
// 	var canvasHeight, canvasWidth int
// 	canvasHeight = 550
// 	canvasWidth = 900
// 	canvas := Canvas(canvasWidth, canvasHeight)

// 	projColour := Colour(1, 0, 0)
// 	e := Environment{Gravity: Vector(0, -0.08, 0), Wind: Vector(-0.05, 0, 0)}

// 	p := Projectile{Position: Point(0, 1, 0), Velocity: Vector(1.4, 1.8, 0).Normalize().Mul(11.25)}
// 	ticks := 0
// 	for {
// 		posX := int(p.Position.X)
// 		posY := canvasHeight - int(p.Position.Y)

// 		if posX > canvasWidth || posX < 0 || posY > canvasHeight || posY < 0 {
// 			break
// 		}
// 		canvas.WritePixel(posX, posY, projColour)

// 		if p.Position.Y < 0.0 || FloatEqual(p.Position.Y, 0) {
// 			break
// 		}
// 		ticks++
// 		p = tick(e, p)
// 	}

// 	ioutil.WriteFile("proj.ppm", []byte(canvas.ToPPM()), 0755)

// 	fmt.Printf("Completed after %d ticks", ticks)
// }

// func tick(e Environment, p Projectile) Projectile {
// 	pos := p.Position.Add(p.Velocity)
// 	vel := p.Velocity.Add(e.Gravity).Add(e.Wind)
// 	return Projectile{Position: pos, Velocity: vel}
// }
