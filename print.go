package main

import (
	"fmt"
	"sort"
)

func (p polynomial) print() {
	for i, m := range p.monomials {
		switch {
		case i != 0 && m.coefficient < 0:
			fmt.Print("- ")
			m.coefficient *= -1
		case i != 0 && m.coefficient > 0:
			fmt.Print("+ ")
		default:
		}
		fmt.Printf("%g * X^%d ", m.coefficient, m.exponent)
	}
	fmt.Println("= 0")
}

func (p *polynomial) reduce() {
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

func (p *polynomial) add(other polynomial) {
	p.monomials = append(p.monomials, other.monomials...)
}

func (p *polynomial) reverse() {
	for i := range p.monomials {
		p.monomials[i].coefficient *= -1
	}
	p.print()
}


func (p polynomial) Equal(other polynomial) bool {
	if len(p.monomials) != len(other.monomials) {
		return false
	}
	for i := range p.monomials {
		if !p.monomials[i].Equal(other.monomials[i]) {
			return false
		}
	}
	return true
}

func (m *monomial) add(other monomial) {
	if m.exponent != other.exponent {
		fmt.Println("Cannot add monomials with different exponents")
		return
	}
	m.coefficient += other.coefficient
}

func (m monomial) Equal(other monomial) bool {
	return m.coefficient == other.coefficient && m.exponent == other.exponent
}
