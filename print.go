package main

import (
	"fmt"
	"sort"
)

func (polynomial Polynomial) print() {
	for i, nomial := range polynomial.monomials {
		switch {
		case i != 0 && nomial.coefficient < 0:
			fmt.Print("- ")
			nomial.coefficient *= -1
		case i != 0 && nomial.coefficient > 0:
			fmt.Print("+ ")
		default:
		}
		fmt.Printf("%g * X^%d ", nomial.coefficient, nomial.exponent)
	}
	fmt.Println("= 0")
}

func (p *Polynomial) reduce() {
	sort.Slice(p.monomials, func(i, j int) bool {
		return p.monomials[i].exponent < p.monomials[j].exponent
	})

	var reducedMonomials []monomial
	i := 0
	for i+1 < len(p.monomials) {
		current := p.monomials[i]
		if current.exponent == p.monomials[i+1].exponent {
			current.add(p.monomials[i+1])
			i++
		}
		reducedMonomials = append(reducedMonomials, current)
		i++
	}
	if i < len(p.monomials) {
		reducedMonomials = append(reducedMonomials, p.monomials[i])
	}

	p.monomials = reducedMonomials
}

func (m *monomial) add(other monomial) {
	if m.exponent != other.exponent {
		fmt.Println("Cannot add monomials with different exponents")
		return
	}
	m.coefficient += other.coefficient
}
