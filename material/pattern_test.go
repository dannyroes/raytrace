package material

import (
	"testing"

	"github.com/dannyroes/raytrace/data"
)

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
		p        data.Tuple
		expected ColourTuple
	}{
		{
			p:        data.Point(0, 0, 0),
			expected: White,
		},
		{
			p:        data.Point(0, 1, 0),
			expected: White,
		},
		{
			p:        data.Point(0, 2, 0),
			expected: White,
		},
		{
			p:        data.Point(0, 0, 1),
			expected: White,
		},
		{
			p:        data.Point(0, 0, 2),
			expected: White,
		},
		{
			p:        data.Point(0.9, 0, 0),
			expected: White,
		},
		{
			p:        data.Point(1, 0, 0),
			expected: Black,
		},
		{
			p:        data.Point(-0.1, 0, 0),
			expected: Black,
		},
		{
			p:        data.Point(-1, 0, 0),
			expected: Black,
		},
		{
			p:        data.Point(-1.1, 0, 0),
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

func TestGradientPatternPoints(t *testing.T) {
	pattern := GradientPattern(White, Black)

	cases := []struct {
		p        data.Tuple
		expected ColourTuple
	}{
		{
			p:        data.Point(0, 0, 0),
			expected: White,
		},
		{
			p:        data.Point(0.25, 0, 0),
			expected: Colour(0.75, 0.75, 0.75),
		},
		{
			p:        data.Point(0.5, 0, 0),
			expected: Colour(0.5, 0.5, 0.5),
		},
		{
			p:        data.Point(0.75, 0, 0),
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
		p        data.Tuple
		expected ColourTuple
	}{
		{
			p:        data.Point(0, 0, 0),
			expected: White,
		},
		{
			p:        data.Point(1, 0, 0),
			expected: Black,
		},
		{
			p:        data.Point(0, 0, 1),
			expected: Black,
		},
		{
			p:        data.Point(0.708, 0, 0.708),
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
		p        data.Tuple
		expected ColourTuple
	}{
		{
			p:        data.Point(0, 0, 0),
			expected: White,
		},
		{
			p:        data.Point(0.99, 0, 0),
			expected: White,
		},
		{
			p:        data.Point(1.01, 0, 0),
			expected: Black,
		},
		{
			p:        data.Point(55.01, 0, 0),
			expected: Black,
		},
		{
			p:        data.Point(0, 0.99, 0),
			expected: White,
		},
		{
			p:        data.Point(0, 1.01, 0),
			expected: Black,
		},
		{
			p:        data.Point(0, 55.01, 0),
			expected: Black,
		},
		{
			p:        data.Point(0, 0, 0.99),
			expected: White,
		},
		{
			p:        data.Point(0, 0, 1.01),
			expected: Black,
		},
		{
			p:        data.Point(0, 0, 55.01),
			expected: Black,
		},
		{
			p:        data.Point(5.01, 5.0, 5.01),
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
