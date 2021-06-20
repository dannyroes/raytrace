package main

import "testing"

func TestMaterial(t *testing.T) {
	m := Material()

	if !TupleEqual(m.Colour.Tuple, Colour(1, 1, 1).Tuple) {
		t.Errorf("bad colour expected: %v received %v", Colour(1, 1, 1), m.Colour)
	}

	if !FloatEqual(m.Ambient, 0.1) {
		t.Errorf("bad ambient expected: %f received %f", 0.1, m.Ambient)
	}

	if !FloatEqual(m.Diffuse, 0.9) {
		t.Errorf("bad diffuse expected: %f received %f", 0.1, m.Diffuse)
	}

	if !FloatEqual(m.Specular, 0.9) {
		t.Errorf("bad specular expected: %f received %f", 0.1, m.Specular)
	}

	if !FloatEqual(m.Shininess, 200.0) {
		t.Errorf("bad shininess expected: %f received %f", 0.1, m.Shininess)
	}
}
