package main

import "fmt"

type Environment struct {
	Gravity Tuple
	Wind    Tuple
}

type Projectile struct {
	Position Tuple
	Velocity Tuple
}

func main() {
	e := Environment{Gravity: Vector(0, -0.1, 0), Wind: Vector(-0.01, 0, 0)}

	p := Projectile{Position: Point(0, 1, 0), Velocity: Vector(1, 1, 0).Normalize().Mul(5)}
	ticks := 0
	for {
		fmt.Printf("%+v\n", p.Position)

		if p.Position.Y < 0.0 || FloatEqual(p.Position.Y, 0) {
			break
		}
		ticks++
		p = tick(e, p)
	}

	fmt.Printf("Completed after %d ticks", ticks)
}

func tick(e Environment, p Projectile) Projectile {
	pos := p.Position.Add(p.Velocity)
	vel := p.Velocity.Add(e.Gravity).Add(e.Wind)
	return Projectile{Position: pos, Velocity: vel}
}
