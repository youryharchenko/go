package main

import "math"

const tolerance = 0.00001
const dx = 0.00001

func sqrt1(x float64) float64 {
	return fixedPoint(
		averageDamp(
			func(y float64) float64 {
				return x / y
			},
		),
		1.0,
	)
}

func cubeRoot(x float64) float64 {
	return fixedPoint(
		averageDamp(
			func(y float64) float64 {
				return x / (y * y)
			},
		),
		1.0,
	)
}

func newtonMethod(g func(float64) float64, guess float64) float64 {
	return fixedPoint(newtonTransform(g), guess)
}

func sqrt2(x float64) float64 {
	return newtonMethod(
		func(y float64) float64 {
			return y*y - x
		},
		1.0,
	)
}

func fixedPoint(f func(float64) float64, firstGuess float64) float64 {
	closeEnough := func(v1, v2 float64) bool {
		return math.Abs(v1-v2) < tolerance
	}
	var try func(float64) float64
	try = func(guess float64) float64 {
		next := f(guess)
		if closeEnough(guess, next) {
			return next
		}
		return try(next)
	}
	return try(firstGuess)
}

func fixedPointOfTransform(g func(float64) float64, transform func(func(float64) float64) func(float64) float64, guess float64) float64 {
	return fixedPoint(transform(g), guess)
}

func sqrt3(x float64) float64 {
	return fixedPointOfTransform(
		func(y float64) float64 {
			return x / y
		},
		averageDamp,
		1.0,
	)
}

func sqrt4(x float64) float64 {
	return fixedPointOfTransform(
		func(y float64) float64 {
			return y*y - x
		},
		newtonTransform,
		1.0,
	)
}

func newtonTransform(g func(float64) float64) func(float64) float64 {
	return func(x float64) float64 { return x - g(x)/deriv(g)(x) }
}

func averageDamp(f func(float64) float64) func(float64) float64 {
	return func(x float64) float64 { return (x + f(x)) / 2.0 }
}

func deriv(g func(float64) float64) func(float64) float64 {
	return func(x float64) float64 { return (g(x+dx) - g(x)) / dx }
}
