package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func isValidChar(prev, current byte) bool {
	validNextChars := map[byte]string{
		'*': "X0123456789.",
		'X': "^*",
		'-': "X0123456789.",
		'+': "X0123456789.",
		'^': "0123456789",
	}
	if prev == 0 { // First char
		return strings.ContainsRune("-+X.0123456789", rune(current))
	} else if valids, ok := validNextChars[prev]; ok {
		return strings.ContainsRune(valids, rune(current))
	} else if unicode.IsDigit(rune(prev)) {
		return current == '*' || current == 'X'
	}
	return false
}

func isEndToken(prev, current byte) bool {
	if prev == 'X' || unicode.IsDigit(rune(prev)) {
		return current == '+' || current == '-'
	}
	return false
}

func parseIntPrefix(input string) (int, int, error) {
	var val int
	_, err := fmt.Sscanf(input, "%d", &val)
	if err != nil {
		return 0, 0, err
	}
	str := fmt.Sprintf("%d", val)
	return val, len(str), nil
}

func parseFloatPrefix(input string) (float64, int, error) {
	floatRegex := regexp.MustCompile(`^[-+]?(?:\d+(?:\.\d*)?|\.\d+)(?:[eE][-+]?\d+)?`)
	match := floatRegex.FindString(input)
	if match == "" {
		return 0, 0, fmt.Errorf("no valid float found: %s", input)
	}
	float, err := strconv.ParseFloat(match, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("parsing float: %v", err)
	}

	return float, len(match), nil
}

func parseToMonomial(input string) (*monomial, int, error) {
	m := &monomial{
		coefficient: 1,
		exponent:    0,
	}

	var isExponent bool
	var seenX bool
	var prev byte
	i := 0
	for i < len(input) {
		if isEndToken(prev, input[i]) {
			return m, i, nil
		}
		if !isValidChar(prev, input[i]) {
			return nil, 0, fmt.Errorf("invalid char: %c", input[i])
		}
		prev = input[i]
		switch c := input[i]; {
		case unicode.IsDigit(rune(c)) && isExponent:
			val, len, err := parseIntPrefix(input[i:])
			if err != nil {
				return nil, 0, fmt.Errorf("parsing exponent: %v\n", err)
			}
			m.exponent = val
			isExponent = false
			i += len - 1
		case unicode.IsDigit(rune(c)):
			val, len, err := parseFloatPrefix(input[i:])
			if err != nil {
				return nil, 0, fmt.Errorf("parsing coefficient: %v\n", err)
			}
			m.coefficient *= val
			i += len - 1
		case c == '-':
			m.coefficient *= -1
		case c == 'X':
			if seenX {
				return nil, 0, errors.New("one clause can only have one X")
			}
			m.exponent = 1
			seenX = true
		case c == '^':
			isExponent = true
		default: // include '+'
		}
		i++
	}
	return m, i, nil
}

func parseToPolynomial(input string) (*polynomial, error) {
	p := &polynomial{}
	for i := 0; i < len(input); {
		monomial, len, err := parseToMonomial(input[i:])
		if err != nil {
			return nil, err
		}
		p.monomials = append(p.monomials, *monomial)
		i += len
	}
	return p, nil
}

func ParseInput(input string) (*polynomial, error) {
	// separate right and left of equation
	if cnt := strings.Count(input, "="); cnt != 1 {
		return nil, fmt.Errorf("equation must contains exactly 1 '=', got: %d", cnt)
	}

	// remove spaces and make all upper case
	trimmed := strings.ReplaceAll(input, " ", "")
	toupper := strings.ToUpper(trimmed)

	// validate each side
	split := strings.Split(toupper, "=")
	LHS := split[0]
	RHS := split[1]
	lhs, err := parseToPolynomial(LHS)
	if err != nil {
		return nil, err
	}
	if len(lhs.monomials) == 0 {
		return nil, errors.New("incomplete equation: missing left hand side")
	}
	rhs, err := parseToPolynomial(RHS)
	if err != nil {
		return nil, err
	}
	if len(rhs.monomials) == 0 {
		return nil, errors.New("incomplete equation: missing right hand side")
	}

	// move all terms to the left side
	rhs.reverse()
	lhs.add(rhs)
	lhs.reduce()
	return lhs, nil
}
