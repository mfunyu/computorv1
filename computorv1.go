package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func getValidatedEquation(argv []string) (string, error) {
	switch len(argv) {
	case 2:
		// validate input
		return argv[1], nil
	case 1:
		// loop until getting a correct equation
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			// validate input
		}
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", errors.New("Each term must follow the form `a * X^p` (coefficient, then power)")
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
