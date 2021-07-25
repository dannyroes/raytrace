package main

import "testing"

func TestStripPattern(t *testing.T) {
	pattern := StripePattern(White, Black)

	if !ColourEqual(pattern.A, White) {
		t.Errorf("Pattern A mismatch expected %+v received %+v", White, pattern.A)
	}
	if !ColourEqual(pattern.B, Black) {
		t.Errorf("Pattern B mismatch expected %+v received %+v", Black, pattern.B)
	}
}

func TestStripePatternPoints(t *testing.T) {
	pattern := StripePattern(White, Black)

	cases := []struct {
		p        Tuple
		expected ColourTuple
	}{
		{
			p:        Point(0, 0, 0),
			expected: White,
		},
		{
			p:        Point(0, 1, 0),
			expected: White,
		},
		{
			p:        Point(0, 2, 0),
			expected: White,
		},
		{
			p:        Point(0, 0, 1),
			expected: White,
		},
		{
			p:        Point(0, 0, 2),
			expected: White,
		},
		{
			p:        Point(0.9, 0, 0),
			expected: White,
		},
		{
			p:        Point(1, 0, 0),
			expected: Black,
		},
		{
			p:        Point(-0.1, 0, 0),
			expected: Black,
		},
		{
			p:        Point(-1, 0, 0),
			expected: Black,
		},
		{
			p:        Point(-1.1, 0, 0),
			expected: White,
		},
	}

	for _, tc := range cases {
		at := pattern.At(tc.p)
		if !ColourEqual(at, tc.expected) {
			t.Errorf("Pattern mismatch at %+v expected %+v received %+v", tc.p, tc.expected, at)
		}
	}
}

func TestPatternTransform(t *testing.T) {
	cases := []struct {
		o        Shape
		p        Pattern
		point    Tuple
		expected ColourTuple
	}{
		{
			o: func() *SphereType {
				s := Sphere()
				s.SetTransform(Scaling(2, 2, 2))

				return s
			}(),
			p:        StripePattern(White, Black),
			point:    Point(1.5, 0, 0),
			expected: White,
		},
		{
			o: Sphere(),
			p: func() Pattern {
				p := StripePattern(White, Black)
				p.SetTransform(Scaling(2, 2, 2))
				return p
			}(),
			point:    Point(1.5, 0, 0),
			expected: White,
		},
		{
			o: func() *SphereType {
				s := Sphere()
				s.SetTransform(Scaling(2, 2, 2))

				return s
			}(),
			p: func() Pattern {
				p := StripePattern(White, Black)
				p.SetTransform(Translation(0.5, 0, 0))
				return p
			}(),
			point:    Point(2.5, 0, 0),
			expected: White,
		},
	}

	for _, tc := range cases {
		at := PatternAtObject(tc.p, tc.o, tc.point)
		if !ColourEqual(at, tc.expected) {
			t.Errorf("Pattern mismatch at %+v expected %+v received %+v", tc.point, tc.expected, at)
		}
	}
}

func TestGradientPatternPoints(t *testing.T) {
	pattern := GradientPattern(White, Black)

	cases := []struct {
		p        Tuple
		expected ColourTuple
	}{
		{
			p:        Point(0, 0, 0),
			expected: White,
		},
		{
			p:        Point(0.25, 0, 0),
			expected: Colour(0.75, 0.75, 0.75),
		},
		{
			p:        Point(0.5, 0, 0),
			expected: Colour(0.5, 0.5, 0.5),
		},
		{
			p:        Point(0.75, 0, 0),
			expected: Colour(0.25, 0.25, 0.25),
		},
	}

	for _, tc := range cases {
		at := pattern.At(tc.p)
		if !ColourEqual(at, tc.expected) {
			t.Errorf("Pattern mismatch at %+v expected %+v received %+v", tc.p, tc.expected, at)
		}
	}
}

func TestRingPatternPoints(t *testing.T) {
	pattern := RingPattern(White, Black)

	cases := []struct {
		p        Tuple
		expected ColourTuple
	}{
		{
			p:        Point(0, 0, 0),
			expected: White,
		},
		{
			p:        Point(1, 0, 0),
			expected: Black,
		},
		{
			p:        Point(0, 0, 1),
			expected: Black,
		},
		{
			p:        Point(0.708, 0, 0.708),
			expected: Black,
		},
	}

	for _, tc := range cases {
		at := pattern.At(tc.p)
		if !ColourEqual(at, tc.expected) {
			t.Errorf("Pattern mismatch at %+v expected %+v received %+v", tc.p, tc.expected, at)
		}
	}
}

func TestCheckersPatternPoints(t *testing.T) {
	pattern := CheckersPattern(White, Black)

	cases := []struct {
		p        Tuple
		expected ColourTuple
	}{
		{
			p:        Point(0, 0, 0),
			expected: White,
		},
		{
			p:        Point(0.99, 0, 0),
			expected: White,
		},
		{
			p:        Point(1.01, 0, 0),
			expected: Black,
		},
		{
			p:        Point(55.01, 0, 0),
			expected: Black,
		},
		{
			p:        Point(0, 0.99, 0),
			expected: White,
		},
		{
			p:        Point(0, 1.01, 0),
			expected: Black,
		},
		{
			p:        Point(0, 55.01, 0),
			expected: Black,
		},
		{
			p:        Point(0, 0, 0.99),
			expected: White,
		},
		{
			p:        Point(0, 0, 1.01),
			expected: Black,
		},
		{
			p:        Point(0, 0, 55.01),
			expected: Black,
		},
		{
			p:        Point(5.01, 5.0, 5.01),
			expected: Black,
		},
	}

	for _, tc := range cases {
		at := pattern.At(tc.p)
		if !ColourEqual(at, tc.expected) {
			t.Errorf("Pattern mismatch at %+v expected %+v received %+v", tc.p, tc.expected, at)
		}
	}
}
