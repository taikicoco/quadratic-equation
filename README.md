# quadratic-equation

The `quadratic` package is a Go library for solving quadratic equations. It provides a simple and efficient way to find the roots of a quadratic equation of the form `ax^2 + bx + c = 0`.

## Usage

The package provides the `Solve` function, which takes three float64 parameters `a`, `b`, and `c` (the coefficients of the equation) and returns two complex128 numbers representing the roots of the equation.

```go
a, b, c := 1.0, 3.0, 2.0
x1, x2 := quadratic.Solve(a, b, c)
```

## Installation

```bash
$ go get github.com/taikicoco/quadratic-equation
```
