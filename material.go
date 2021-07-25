package main

type MaterialType struct {
	Colour    ColourTuple
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
	Pattern   Pattern
}

func Material() MaterialType {
	return MaterialType{
		Colour:    Colour(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}

func (m MaterialType) Equals(m2 MaterialType) bool {
	return FloatEqual(m.Ambient, m2.Ambient) &&
		FloatEqual(m.Diffuse, m2.Diffuse) &&
		FloatEqual(m.Specular, m2.Specular) &&
		FloatEqual(m.Shininess, m2.Shininess) &&
		TupleEqual(m.Colour.Tuple, m2.Colour.Tuple)
}
