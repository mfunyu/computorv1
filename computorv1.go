package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func getValidatedEquation(argv []string) (*polynomial, error) {
	switch len(argv) {
	case 2:
		// validate input
		polynomial, err := ParseInput(argv[1])
		if err != nil {
			return nil, err
		}
		return polynomial, nil
	case 1:
		// loop until getting a correct equation
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			// validate input
			if polynomial, err := ParseInput(line); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			} else {
				return polynomial, nil
			}
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		return nil, errors.New("missing a valid equation")
	default:
		fmt.Println("Usage: ./computorv1 <equation>")
		return nil, errors.New("invalid number of arguments")
	}
}

func computorv1() error {
	polynomial, err := getValidatedEquation(os.Args)
	if err != nil {
		return err
	}

	fmt.Printf("Reduced form: ")
	polynomial.print()
	fmt.Printf("Polynomial degree: %d\n", polynomial.degree())

	return nil
}

func main() {
	if err := computorv1(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
