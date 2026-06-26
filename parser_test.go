package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseInput_Success(t *testing.T) {

	tests := []struct {
		name string
		arg  string
		want polynomial
	}{
		{
			name: "Most classic function with spaces and multipication sign",
			arg:  "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0",
			want: polynomial{
				monomials: []monomial{
					{coefficient: 5, exponent: 0},
					{coefficient: 4, exponent: 1},
					{coefficient: -9.3, exponent: 2},
					{coefficient: 1, exponent: 0},
				},
			},
		},
		{
			name: "Inversed order with spaces and multipication sign",
			arg:  "X^0* 5 + X^1 * 1 - X^2 * 9.3 = 1 * X^0",
			want: polynomial{
				monomials: []monomial{
					{coefficient: 5, exponent: 0},
					{coefficient: 1, exponent: 1},
					{coefficient: 9.3, exponent: 2},
					{coefficient: 1, exponent: 0},
				},
			},
		},
		{
			name: "Basic",
			arg:  "5X^0 + 4X^1 - 9.3X^2 = 1X^0",
			want: polynomial{
				monomials: []monomial{
					{coefficient: 5, exponent: 0},
					{coefficient: 4, exponent: 1},
					{coefficient: 9.3, exponent: 2},
					{coefficient: 1, exponent: 0},
				},
			},
		},
		{
			name: "No spaces",
			arg:  "5X^0+4X^1-9.3X^2=1X^0",
			want: polynomial{
				monomials: []monomial{
					{coefficient: 5, exponent: 0},
					{coefficient: 4, exponent: 1},
					{coefficient: 9.3, exponent: 2},
					{coefficient: 1, exponent: 0},
				},
			},
		},
		{
			name: "Omitted X^0 and X^1",
			arg:  "5+4X-9.3X^2=1",
			want: polynomial{
				monomials: []monomial{
					{coefficient: 5, exponent: 0},
					{coefficient: 4, exponent: 1},
					{coefficient: 9.3, exponent: 2},
					{coefficient: 1, exponent: 0},
				},
			},
		},
		{
			name: "Have operators for the first expressions",
			arg:  "-5X^0+4X^1-9.3X^2=+1X^0",
			want: polynomial{
				monomials: []monomial{
					{coefficient: 5, exponent: 0},
					{coefficient: 4, exponent: 1},
					{coefficient: 9.3, exponent: 2},
					{coefficient: 1, exponent: 0},
				},
			},
		},
		{
			name: "Omitted coefficient when 1",
			arg:  "5+X-X^2=1",
			want: polynomial{
				monomials: []monomial{
					{coefficient: 5, exponent: 0},
					{coefficient: 1, exponent: 1},
					{coefficient: 1, exponent: 2},
					{coefficient: 1, exponent: 0},
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

// func TestParseInput_Error(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		arg args
// 		want
// 	}{
// 		{
// 			name: "Negative exponents",
// 			arg: "5 * X^0 + 4 * X^-1 - 9.3 * X^-2 = 1 * X^0",
// 			want:
// 		},
// 		{
// 			name: "Negative exponents with no space, no multiplication operators",
// 			arg: "5X^0+4X^-1-9.3X^-2=1X^0",
// 			want:
// 		},
// 		{
// 			name: "Negative exponents with omitted coefficient when 1",
// 			arg: "5+X^+1-X^2=1",
// 			want:
// 		},
// 		{
// 			name: "Missing ^ for exponent expression",
// 			arg: "5X0+ 4 * X1 - 9.3 * X^2 = 1 * X^0",
// 			want:
// 		},
// 		{
// 			name: "Duplicated signs",
// 			arg: "--5++4X-9.3X^2=1",
// 			want:
// 		},
// 		{
// 			name: "Missing an operator",
// 			arg: "5X0 + 4 * X1 9.3 * X^2 = 1 * X^0",
// 			want:
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()

// 			got, err := computorv1(t)
// 			if err != nil {
// 				t.Errorf("GetValidatedEquation() expect nil error but got %v", err)
// 			}

// 			if diff := pkgcmp.Diff(got, tt.want); diff != nil {
// 				t.Errorf("GetValidatedEquation() response mismatch = (-want +got):\n%s", diff)
// 			}
// 		}
// 	}
// }
