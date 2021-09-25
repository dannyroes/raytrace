package data

type RayType struct {
	Origin    Tuple
	Direction Tuple
}

func Ray(origin, direction Tuple) RayType {
	return RayType{origin, direction}
}

func (r RayType) Position(t float64) Tuple {
	return r.Origin.Add(r.Direction.Mul(t))
}

func (r RayType) Transform(m Matrix) RayType {
	return RayType{
		Origin:    m.MultiplyTuple(r.Origin),
		Direction: m.MultiplyTuple(r.Direction),
	}
}
