package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Polynomial struct {
	monomials []monomial
}

type monomial struct {
	operator    operators
	coefficient float64
	exponent    int
}

type operators int

const (
	negative operators = -1
	positive operators = 1
)

func parseInput(input string) (string, error) {
	// separate right and left of equation
	if cnt := strings.Count(input, "="); cnt != 1 {
		return "", fmt.Errorf("equation must contains exactly 1 '=', got: %d", cnt)
	}
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
