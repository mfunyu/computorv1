package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseInput_Success(t *testing.T) {

	tests := []struct {
		name string
		arg  string
		want *polynomial
	}{
		{
			name: "Basic function with spaces only left hand side",
			arg:  "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 0",
			want: &polynomial{
				monomials: []monomial{
					{coefficient: 5, exponent: 0},
					{coefficient: 4, exponent: 1},
					{coefficient: -9.3, exponent: 2},
				},
			},
		},
		{
			name: "Basic function with spaces only right hand side",
			arg:  "0 = 5 * X^0 + 4 * X^1 - 9.3 * X^2",
			want: &polynomial{
				monomials: []monomial{
					{coefficient: -5, exponent: 0},
					{coefficient: -4, exponent: 1},
					{coefficient: 9.3, exponent: 2},
				},
			},
		},
		{
			name: "Most classic function with spaces and multipication sign",
			arg:  "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0",
			want: &polynomial{
				monomials: []monomial{
					{coefficient: 4, exponent: 0},
					{coefficient: 4, exponent: 1},
					{coefficient: -9.3, exponent: 2},
				},
			},
		},
		{
			name: "Inversed order with spaces and multipication sign",
			arg:  "X^0 * 5 + X^1 * 1 - X^2 * 9.3 = 1 * X^0",
			want: &polynomial{
				monomials: []monomial{
					{coefficient: 4, exponent: 0},
					{coefficient: 1, exponent: 1},
					{coefficient: -9.3, exponent: 2},
				},
			},
		},
		{
			name: "Basic",
			arg:  "5X^0 + 4X^1 - 9.3X^2 = 1X^0",
			want: &polynomial{
				monomials: []monomial{
					{coefficient: 4, exponent: 0},
					{coefficient: 4, exponent: 1},
					{coefficient: -9.3, exponent: 2},
				},
			},
		},
		{
			name: "No spaces",
			arg:  "5X^0+4X^1-9.3X^2=1X^0",
			want: &polynomial{
				monomials: []monomial{
					{coefficient: 4, exponent: 0},
					{coefficient: 4, exponent: 1},
					{coefficient: -9.3, exponent: 2},
				},
			},
		},
		{
			name: "Omitted X^0 and X^1",
			arg:  "5+4X-9.3X^2=1",
			want: &polynomial{
				monomials: []monomial{
					{coefficient: 4, exponent: 0},
					{coefficient: 4, exponent: 1},
					{coefficient: -9.3, exponent: 2},
				},
			},
		},
		{
			name: "Have operators for the first expressions",
			arg:  "-5X^0+4X^1-9.3X^2=+1X^0",
			want: &polynomial{
				monomials: []monomial{
					{coefficient: -6, exponent: 0},
					{coefficient: 4, exponent: 1},
					{coefficient: -9.3, exponent: 2},
				},
			},
		},
		{
			name: "Omitted coefficient when 1",
			arg:  "5+X-X^2=1",
			want: &polynomial{
				monomials: []monomial{
					{coefficient: 4, exponent: 0},
					{coefficient: 1, exponent: 1},
					{coefficient: -1, exponent: 2},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ParseInput(tt.arg)
			if err != nil {
				t.Errorf("ParseInput() expect nil error but got %v", err)
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ParseInput() response mismatch = (-want +got):\n%s", diff)
			}
		})
	}
}

func TestParseInput_Error(t *testing.T) {
	tests := []struct {
		name string
		arg  string
	}{
		{
			name: "Negative exponents",
			arg:  "5 * X^0 + 4 * X^-1 - 9.3 * X^-2 = 1 * X^0",
		},
		{
			name: "Negative exponents with no space, no multiplication operators",
			arg:  "5X^0+4X^-1-9.3X^-2=1X^0",
		},
		{
			name: "Negative exponents with omitted coefficient when 1",
			arg:  "5+X^+1-X^2=1",
		},
		{
			name: "Missing ^ for exponent expression",
			arg:  "5X0+ 4 * X1 - 9.3 * X^2 = 1 * X^0",
		},
		{
			name: "Duplicated signs",
			arg:  "--5++4X-9.3X^2=1",
		},
		{
			name: "Missing an operator",
			arg:  "5X0 + 4 * X1 9.3 * X^2 = 1 * X^0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ParseInput(tt.arg)
			if err == nil {
				t.Errorf("ParseInput() expect error but got %v", got)
			}
		})
	}
}
