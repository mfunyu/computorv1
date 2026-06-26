package main

import (
	"fmt"
	"sort"
)

func (polynomial *Polynomial) print() {
	for i, nomial := range polynomial.monomials {
		switch {
		case i == 0 && nomial.operator == -1:
			fmt.Print("-")
		case i != 0 && nomial.operator == -1:
			fmt.Print("- ")
		case i != 0 && nomial.operator == 1:
			fmt.Print("+ ")
		default:
		}
		fmt.Printf("%g * X^%d ", nomial.coefficient, nomial.exponent)
	}
	fmt.Println("= 0")
}

func (polynomial *Polynomial) reduce() {
	sort.Slice(polynomial.monomials, func(i, j int) bool {
		return polynomial.monomials[i].exponent < polynomial.monomials[j].exponent
	})
	for i := 0; i+1 < len(polynomial.monomials); {
		if polynomial.monomials[i].exponent != polynomial.monomials[i+1].exponent {
			i++
		}
		i++
	}
}
