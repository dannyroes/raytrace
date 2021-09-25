package world

import (
	"math"

	"github.com/dannyroes/raytrace/data"
	"github.com/dannyroes/raytrace/material"
	"github.com/dannyroes/raytrace/shape"
)

type Light struct {
	Position  data.Tuple
	Intensity material.ColourTuple
}

func PointLight(pos data.Tuple, intensity material.ColourTuple) Light {
	return Light{pos, intensity}
}

func Lighting(m material.MaterialType, object shape.Shape, l Light, pos data.Tuple, eyeV data.Tuple, normalV data.Tuple, inShadow bool) material.ColourTuple {
	colour := m.Colour
	if m.Pattern != nil {
		colour = shape.PatternAtObject(m.Pattern, object, pos)
	}

	effective := material.MultiplyColours(colour, l.Intensity)
	lightV := l.Position.Sub(pos).Normalize()
	ambient := effective.Mul(m.Ambient)
	diffuse := material.Black
	specular := material.Black

	if !inShadow {
		lightDotNormal := data.Dot(lightV, normalV)
		if lightDotNormal >= 0 {
			diffuse = effective.Mul(m.Diffuse).Mul(lightDotNormal)

			reflectV := lightV.Neg().Reflect(normalV)
			reflectDotEye := data.Dot(reflectV, eyeV)
			if reflectDotEye > 0 {
				factor := math.Pow(reflectDotEye, m.Shininess)
				specular = l.Intensity.Mul(m.Specular).Mul(factor)
			}
		}
	}
	return ambient.Add(diffuse).Add(specular)
}
