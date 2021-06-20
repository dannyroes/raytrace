package main

import "math"

type Light struct {
	Position  Tuple
	Intensity ColourTuple
}

func PointLight(pos Tuple, intensity ColourTuple) Light {
	return Light{pos, intensity}
}

func Lighting(m MaterialType, l Light, pos Tuple, eyeV Tuple, normalV Tuple) ColourTuple {
	effective := MultiplyColours(m.Colour, l.Intensity)
	lightV := l.Position.Sub(pos).Normalize()
	ambient := effective.Mul(m.Ambient)
	diffuse := Black
	specular := Black

	lightDotNormal := Dot(lightV, normalV)
	if lightDotNormal >= 0 {
		diffuse = effective.Mul(m.Diffuse).Mul(lightDotNormal)

		reflectV := lightV.Neg().Reflect(normalV)
		reflectDotEye := Dot(reflectV, eyeV)
		if reflectDotEye > 0 {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = l.Intensity.Mul(m.Specular).Mul(factor)
		}
	}
	return ambient.Add(diffuse).Add(specular)
}
