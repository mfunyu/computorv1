package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func isValidChar(current, next byte) bool {
	validNextChars := map[byte]string{
		'*': "X0123456789.",
		'X': "^*",
		'-': "X0123456789.",
		'+': "X0123456789.",
		'^': "0123456789",
	}
	if current == 0 { // First char
		return strings.ContainsRune("-+X.0123456789", rune(next))
	} else if valids, ok := validNextChars[current]; ok {
		return strings.ContainsRune(valids, rune(next))
	} else if unicode.IsDigit(rune(current)) {
		return next == '*' || next == 'X'
	}
	fmt.Printf("Unexpected current char: %c\n, next char: %c\n", current, next)
	return false
}

func isEndToken(before, current byte) bool {
	if before == 'X' || unicode.IsDigit(rune(before)) {
		return current == '+' || current == '-'
	}
	return false
}

func parseIntPrefix(input string) (int, int, error) {
	var int_val int
	_, err := fmt.Sscanf(input, "%d", &int_val)
	if err != nil {
		return 0, 0, err
	}
	int_str := fmt.Sprintf("%d", int_val)
	return int_val, len(int_str), nil
}

func parseFloatPrefix(input string) (float64, int, error) {
	floatRegex := regexp.MustCompile(`^[-+]?(?:\d+(?:\.\d*)?|\.\d+)(?:[eE][-+]?\d+)?`)
	match := floatRegex.FindString(input)
	if match == "" {
		return 0, 0, fmt.Errorf("no valid float found in input: %s", input)
	}
	float, err := strconv.ParseFloat(match, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse float: %v", err)
	}

	return float, len(match), nil
}

func parseToMonomial(input string) (monomial, int) {
	m := monomial{
		coefficient: 1,
		exponent:    0,
	}

	isExponent := false
	seenX := false

	var before byte = 0
	i := 0
	for i < len(input) {
		if isEndToken(before, input[i]) {
			return m, i
		}
		if !isValidChar(before, input[i]) {
			//should return error
			break
		}
		before = input[i]
		switch c := input[i]; {
		case unicode.IsDigit(rune(c)):
			var len int
			var float float64
			var err error
			if isExponent {
				var int_val int
				int_val, len, err = parseIntPrefix(input[i:])
				if err != nil {
					fmt.Printf("Error parsing exponent: %v\n", err)
					// should return error
					break
				}
				m.exponent = int_val
				isExponent = false
			} else {
				float, len, err = parseFloatPrefix(input[i:])
				if err != nil {
					// should return error
					fmt.Printf("Error parsing float: %v\n", err)
					break
				}
				m.coefficient *= float
			}
			i += len
		case c == '+':
			fallthrough
		case c == '-':
			// '-' -> -1, '+' -> 1
			m.coefficient *= float64(44 - int(c))
			i++
		case c == 'X':
			if seenX {
				// should return error
				break
			}
			m.exponent = 1
			seenX = true
			i++
		case c == '^':
			isExponent = true
			i++
		default:
			i++
		}
	}
	return m, i
}

func parseToPolynomial(input string) (*polynomial, error) {
	p := &polynomial{}
	for i := 0; i < len(input); {
		monomial, len := parseToMonomial(input[i:])
		p.monomials = append(p.monomials, monomial)
		i += len
	}
	return p, nil
}

func ParseInput(input string) (*polynomial, error) {
	// separate right and left of equation
	if cnt := strings.Count(input, "="); cnt != 1 {
		return nil, fmt.Errorf("equation must contains exactly 1 '=', got: %d", cnt)
	}

	trimmed := strings.ReplaceAll(input, " ", "")
	toupper := strings.ToUpper(trimmed)
	split := strings.Split(toupper, "=")
	LHS := split[0]
	RHS := split[1]
	lhs, err := parseToPolynomial(LHS)
	if err != nil {
		return nil, err
	}
	rhs, err := parseToPolynomial(RHS)
	if err != nil {
		return nil, err
	}
	rhs.reverse()
	lhs.add(rhs)
	lhs.reduce()

	return lhs, nil
}
