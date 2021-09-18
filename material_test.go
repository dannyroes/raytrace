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

	if !FloatEqual(m.Transparency, 0.0) {
		t.Errorf("bad transparency expected: %f received %f", 0.1, m.Transparency)
	}

	if !FloatEqual(m.RefractiveIndex, 1.0) {
		t.Errorf("bad refractive index expected: %f received %f", 0.1, m.RefractiveIndex)
	}
}

func TestMaterialPattern(t *testing.T) {
	m := Material()
	m.Ambient = 1
	m.Diffuse = 0
	m.Specular = 0
	m.Pattern = StripePattern(White, Black)

	eyeV := Vector(0, 0, -1)
	normalV := Vector(0, 0, -1)
	light := PointLight(Point(0, 0, -1), White)

	c1 := Lighting(m, Sphere(), light, Point(0.9, 0, 0), eyeV, normalV, false)
	if !ColourEqual(c1, White) {
		t.Errorf("C1 mismatch expected %+v received %+v", White, c1)
	}

	c2 := Lighting(m, Sphere(), light, Point(1.1, 0, 0), eyeV, normalV, false)
	if !ColourEqual(c2, Black) {
		t.Errorf("C2 mismatch expected %+v received %+v", Black, c2)
	}
}
