package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Polynomial struct {
	monomials []monomial
}

type monomial struct {
	operator    int
	coefficient float64
	exponent    int
}

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
	_, err := fmt.Sscanf(input, "%f", &int_val)
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
		return 0, 0,fmt.Errorf("no valid float found in input: %s", input)
	}
	float, err := strconv.ParseFloat(match, 64)
	if err != nil {
		return 0, 0,fmt.Errorf("failed to parse float: %v", err)
	}

	return float, len(match), nil
}

func parseToMonomial(input string) (monomial, int) {
	nomial := monomial{
		operator:    1,
		coefficient: 1,
		exponent:    0,
	}

	isExponent := false
	seenX := false

	var before byte = 0
	i := 0
	for i < len(input) {
		if isEndToken(before, input[i]) {
			return nomial, i
		}
		if !isValidChar(before, input[i]) {
			//should return error
			break
		}
		before = input[i]
		switch c := input[i]; {
		case unicode.IsDigit(rune(c)):
			var len int
			var err error
			if isExponent {
				var int_val int
				int_val, len, err = parseIntPrefix(input[i:])
				if err != nil {
					// should return error
					break
				}
				nomial.exponent = int_val
				isExponent = false
			} else {
				float, len, err := parseFloatPrefix(input[i:])
				if err != nil {
					// should return error
					break
				}
				nomial.coefficient *= float
				fmt.Printf("coefficient: %f\n, len: %d\n", float, len)
			}
			i += len
		case c == '+':
			fallthrough
		case c == '-':
			// '-' -> -1, '+' -> 1
			nomial.operator *= int(byte(54) - c)
			i++
		case c == 'X':
			if seenX {
				// should return error
				break
			}
			seenX = true
		case c == '^':
			isExponent = true
			i++
		default:
			// should return error
			break
		}
		fmt.Printf("%c", rune(input[i]))
	}
	return nomial, i
}

func parseToPolynomial(input string) Polynomial {
	var polynomial Polynomial

	for i := 0; i < len(input); {
		monomial, len := parseToMonomial(input[i:])
		i += len
		i++
		fmt.Printf("operator: %d, coefficient: %f, exponent: %d\n", monomial.operator, monomial.coefficient, monomial.exponent)
		polynomial.monomials = append(polynomial.monomials, monomial)
	}
	return polynomial
}

func parseInput(input string) (string, error) {
	// separate right and left of equation
	if cnt := strings.Count(input, "="); cnt != 1 {
		return "", fmt.Errorf("equation must contains exactly 1 '=', got: %d", cnt)
	}

	trimmed := strings.ReplaceAll(input, " ", "")
	toupper := strings.ToUpper(trimmed)

	split := strings.Split(toupper, "=")
	LHS := split[0]
	// RHS := split[1]

	parseToPolynomial(LHS)
	// rule 1: ^ needs to be present, cannot have
	// possible chars: number, '.', 'X', '-/+', '*', '^'
	// for number -> *, X, end
	// for X(x) -> ^, *, end
	// for - / + -> number, X
	// for ^ -> number
	// start -> -/+, number, X

	// 5X
	// -5X
	// X^(number) / X *+-

	// for number -> check next: if X = coefficience, +
	return input, nil
}

func getValidatedEquation(argv []string) (string, error) {
	switch len(argv) {
	case 2:
		// validate input
		return parseInput(argv[1])
	case 1:
		// loop until getting a correct equation
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			// validate input
			if _, err := parseInput(line); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			} else {
				return line, nil
			}
		}
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", errors.New("missing a valid equation")
	default:
		fmt.Println("Usage: ./computorv1 <equation>")
		return "", nil
	}
	return "", nil
}

func computorv1() error {
	equation, err := getValidatedEquation(os.Args)
	if err != nil {
		return err
	}

	fmt.Printf("Reduced form: %s", equation)

	return nil
}

func main() {
	if err := computorv1(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
