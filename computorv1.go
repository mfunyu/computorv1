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
		ParseInput(argv[1])
		return "", nil
	case 1:
		// loop until getting a correct equation
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			// validate input
			if _, err := ParseInput(line); err != nil {
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
