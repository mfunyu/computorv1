# Computor v1

A polynomial equation solver written from scratch — no math library functions used.

Computor v1 parses a polynomial equation, reduces it, determines its degree, and solves it for degrees 0, 1, and 2. It handles the discriminant case for quadratic equations and reports its sign.

## Features

- Reduces an equation to its canonical form `a * X^0 + b * X^1 + c * X^2 = 0`
- Reports the polynomial degree
- Solves equations of degree 0, 1, and 2
- For degree 2, reports the sign of the discriminant and the corresponding solution(s)
- No external math library — square roots and all arithmetic are implemented by hand

## Usage

```
$> ./computor "<equation>"
```

Each term must follow the form `a * X^p` (coefficient, then power).

## Example

```
$> ./computor "5 * X^0 + 4 * X^1 - 9.3 * X^2 = 1 * X^0"
Reduced form: 4 * X^0 + 4 * X^1 - 9.3 * X^2 = 0
Polynomial degree: 2
Discriminant is strictly positive, the two solutions are:
0.905239
-0.475131
```

## Behavior by degree

Depending on the degree and, for quadratics, the sign of the discriminant, the program prints one of the following messages:

| Degree | Case | Output message |
| ------ | ---- | -------------- |
| 0      | `0 = 0` | `Any real number is a solution.` |
| 0      | `c = 0`, c ≠ 0 | `No solution.` |
| 1      | — | `The solution is:` (followed by the single solution) |
| 2      | discriminant = 0 | `The solution is:` (followed by the single solution) |
| 2      | discriminant > 0 | `Discriminant is strictly positive, the two solutions are:` (followed by two real solutions) |
| 2      | discriminant < 0 | `Discriminant is strictly negative, the two complex solutions are:` (followed by two complex solutions) |
| > 2    | — | `The polynomial degree is strictly greater than 2, I can't solve.` |

## Notes

- Input is expected to be well-formed: every term respects the `a * X^p` format.
- No complex functions are used in the resolution.
