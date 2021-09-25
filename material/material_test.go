package material

import (
	"testing"

	"github.com/dannyroes/raytrace/data"
)

func TestMaterial(t *testing.T) {
	m := Material()

	if !data.TupleEqual(m.Colour.Tuple, Colour(1, 1, 1).Tuple) {
		t.Errorf("bad colour expected: %v received %v", Colour(1, 1, 1), m.Colour)
	}

	if !data.FloatEqual(m.Ambient, 0.1) {
		t.Errorf("bad ambient expected: %f received %f", 0.1, m.Ambient)
	}

	if !data.FloatEqual(m.Diffuse, 0.9) {
		t.Errorf("bad diffuse expected: %f received %f", 0.1, m.Diffuse)
	}

	if !data.FloatEqual(m.Specular, 0.9) {
		t.Errorf("bad specular expected: %f received %f", 0.1, m.Specular)
	}

	if !data.FloatEqual(m.Shininess, 200.0) {
		t.Errorf("bad shininess expected: %f received %f", 0.1, m.Shininess)
	}

	if !data.FloatEqual(m.Transparency, 0.0) {
		t.Errorf("bad transparency expected: %f received %f", 0.1, m.Transparency)
	}

	if !data.FloatEqual(m.RefractiveIndex, 1.0) {
		t.Errorf("bad refractive index expected: %f received %f", 0.1, m.RefractiveIndex)
	}
}
