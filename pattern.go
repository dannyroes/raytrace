package main

import (
	"math"
)

type Pattern interface {
	At(Tuple) ColourTuple
	SetTransform(Matrix)
	GetTransform() Matrix
}

type StripePatternType struct {
	A         ColourTuple
	B         ColourTuple
	Transform Matrix
}

func PatternAtObject(p Pattern, o Shape, point Tuple) ColourTuple {
	objectPoint := o.GetTransform().Invert().MultiplyTuple(point)
	patternPoint := p.GetTransform().Invert().MultiplyTuple(objectPoint)

	return p.At(patternPoint)
}

func StripePattern(a, b ColourTuple) *StripePatternType {
	return &StripePatternType{A: a, B: b, Transform: IdentityMatrix()}
}

func (p *StripePatternType) At(point Tuple) ColourTuple {
	if int64(math.Floor(point.X))%2 == 0 {
		return p.A
	}
	return p.B
}

func (p *StripePatternType) GetTransform() Matrix {
	return p.Transform
}

func (p *StripePatternType) SetTransform(m Matrix) {
	p.Transform = m
}

type GradientPatternType struct {
	A         ColourTuple
	B         ColourTuple
	Transform Matrix
}

func GradientPattern(a, b ColourTuple) *GradientPatternType {
	return &GradientPatternType{A: a, B: b, Transform: IdentityMatrix()}
}

func (p *GradientPatternType) At(point Tuple) ColourTuple {
	dist := p.B.Sub(p.A)
	fract := point.X - math.Floor(point.X)

	return p.A.Add(dist.Mul(fract))
}

func (p *GradientPatternType) GetTransform() Matrix {
	return p.Transform
}

func (p *GradientPatternType) SetTransform(m Matrix) {
	p.Transform = m
}

type RingPatternType struct {
	A         ColourTuple
	B         ColourTuple
	Transform Matrix
}

func RingPattern(a, b ColourTuple) *RingPatternType {
	return &RingPatternType{A: a, B: b, Transform: IdentityMatrix()}
}

func (p *RingPatternType) At(point Tuple) ColourTuple {
	squares := point.X*point.X + point.Z*point.Z

	if int64(math.Sqrt(math.Floor(squares)))%2 == 0 {
		return p.A
	}
	return p.B
}

func (p *RingPatternType) GetTransform() Matrix {
	return p.Transform
}

func (p *RingPatternType) SetTransform(m Matrix) {
	p.Transform = m
}

type CheckersPatternType struct {
	A         ColourTuple
	B         ColourTuple
	Transform Matrix
}

func CheckersPattern(a, b ColourTuple) *CheckersPatternType {
	return &CheckersPatternType{A: a, B: b, Transform: IdentityMatrix()}
}

func (p *CheckersPatternType) At(point Tuple) ColourTuple {
	sum := int64(math.Floor(point.X)) + int64(math.Floor(point.Y)) + int64(math.Floor(point.Z))

	if sum%2 == 0 {
		return p.A
	}
	return p.B
}

func (p *CheckersPatternType) GetTransform() Matrix {
	return p.Transform
}

func (p *CheckersPatternType) SetTransform(m Matrix) {
	p.Transform = m
}
