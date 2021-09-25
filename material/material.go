package material

import "github.com/dannyroes/raytrace/data"

type MaterialType struct {
	Colour          ColourTuple
	Ambient         float64
	Diffuse         float64
	Specular        float64
	Shininess       float64
	Reflective      float64
	Transparency    float64
	RefractiveIndex float64
	Pattern         Pattern
}

func Material() MaterialType {
	return MaterialType{
		Colour:          Colour(1, 1, 1),
		Ambient:         0.1,
		Diffuse:         0.9,
		Specular:        0.9,
		Shininess:       200.0,
		Reflective:      0.0,
		Transparency:    0.0,
		RefractiveIndex: 1.0,
	}
}

func (m MaterialType) Equals(m2 MaterialType) bool {
	return data.FloatEqual(m.Ambient, m2.Ambient) &&
		data.FloatEqual(m.Diffuse, m2.Diffuse) &&
		data.FloatEqual(m.Specular, m2.Specular) &&
		data.FloatEqual(m.Shininess, m2.Shininess) &&
		data.TupleEqual(m.Colour.Tuple, m2.Colour.Tuple)
}
