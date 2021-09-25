package material

import (
	"math"

	"github.com/dannyroes/raytrace/data"
)

type Pattern interface {
	At(data.Tuple) ColourTuple
	SetTransform(data.Matrix)
	GetTransform() data.Matrix
}

type StripePatternType struct {
	A         ColourTuple
	B         ColourTuple
	Transform data.Matrix
}

func StripePattern(a, b ColourTuple) *StripePatternType {
	return &StripePatternType{A: a, B: b, Transform: data.IdentityMatrix()}
}

func (p *StripePatternType) At(point data.Tuple) ColourTuple {
	if int64(math.Floor(point.X))%2 == 0 {
		return p.A
	}
	return p.B
}

func (p *StripePatternType) GetTransform() data.Matrix {
	return p.Transform
}

func (p *StripePatternType) SetTransform(m data.Matrix) {
	p.Transform = m
}

type GradientPatternType struct {
	A         ColourTuple
	B         ColourTuple
	Transform data.Matrix
}

func GradientPattern(a, b ColourTuple) *GradientPatternType {
	return &GradientPatternType{A: a, B: b, Transform: data.IdentityMatrix()}
}

func (p *GradientPatternType) At(point data.Tuple) ColourTuple {
	dist := p.B.Sub(p.A)
	fract := point.X - math.Floor(point.X)

	return p.A.Add(dist.Mul(fract))
}

func (p *GradientPatternType) GetTransform() data.Matrix {
	return p.Transform
}

func (p *GradientPatternType) SetTransform(m data.Matrix) {
	p.Transform = m
}

type RingPatternType struct {
	A         ColourTuple
	B         ColourTuple
	Transform data.Matrix
}

func RingPattern(a, b ColourTuple) *RingPatternType {
	return &RingPatternType{A: a, B: b, Transform: data.IdentityMatrix()}
}

func (p *RingPatternType) At(point data.Tuple) ColourTuple {
	squares := point.X*point.X + point.Z*point.Z

	if int64(math.Sqrt(math.Floor(squares)))%2 == 0 {
		return p.A
	}
	return p.B
}

func (p *RingPatternType) GetTransform() data.Matrix {
	return p.Transform
}

func (p *RingPatternType) SetTransform(m data.Matrix) {
	p.Transform = m
}

type CheckersPatternType struct {
	A         ColourTuple
	B         ColourTuple
	Transform data.Matrix
}

func CheckersPattern(a, b ColourTuple) *CheckersPatternType {
	return &CheckersPatternType{A: a, B: b, Transform: data.IdentityMatrix()}
}

func (p *CheckersPatternType) At(point data.Tuple) ColourTuple {
	sum := int64(math.Floor(point.X)) + int64(math.Floor(point.Y)) + int64(math.Floor(point.Z))

	if sum%2 == 0 {
		return p.A
	}
	return p.B
}

func (p *CheckersPatternType) GetTransform() data.Matrix {
	return p.Transform
}

func (p *CheckersPatternType) SetTransform(m data.Matrix) {
	p.Transform = m
}

type TestPatternType struct {
	Transform data.Matrix
}

func TestPattern() *TestPatternType {
	return &TestPatternType{Transform: data.IdentityMatrix()}
}

func (p *TestPatternType) At(point data.Tuple) ColourTuple {
	return Colour(point.X, point.Y, point.Z)
}

func (p *TestPatternType) GetTransform() data.Matrix {
	return p.Transform
}

func (p *TestPatternType) SetTransform(m data.Matrix) {
	p.Transform = m
}
