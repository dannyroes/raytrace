package main

type MaterialType struct {
	Colour    ColourTuple
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
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
