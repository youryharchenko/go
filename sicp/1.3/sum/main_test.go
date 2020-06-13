package main

import "testing"

func TestCube(t *testing.T) {
	cube := func(n int) int {
		return n * n * n
	}
	inc := func(n int) int {
		return n + 1
	}
	sumCubes := func(a, b int) int {
		return sumInt(cube, a, inc, b)
	}
	println(sumCubes(1, 10))
	t.Error("")
}

func TestCube2(t *testing.T) {
	sumCubes := func(a, b int) int {
		return sumInt(
			func(n int) int {
				return n * n * n
			},
			a,
			func(n int) int {
				return n + 1
			},
			b,
		)
	}
	println(sumCubes(1, 10))
	t.Error("")
}

func TestIntegers(t *testing.T) {
	identity := func(n int) int {
		return n
	}
	inc := func(n int) int {
		return n + 1
	}
	sumIntegers := func(a, b int) int {
		return sumInt(identity, a, inc, b)
	}
	println(sumIntegers(1, 10))
	t.Error("")
}
func TestIntegers2(t *testing.T) {
	sumIntegers := func(a, b int) int {
		return sumInt(
			func(n int) int {
				return n
			},
			a,
			func(n int) int {
				return n + 1
			},
			b,
		)
	}
	println(sumIntegers(1, 10))
	t.Error("")
}

func TestPi(t *testing.T) {
	sumPi := func(a, b float64) float64 {
		return sumFloat(
			func(x float64) float64 {
				return 1.0 / (x * (x + 2))
			},
			a,
			func(x float64) float64 {
				return x + 4.0
			},
			b,
		)
	}
	println(8 * sumPi(1.0, 1000.0))
	t.Error("")
}

func TestIntegral(t *testing.T) {
	println(integral(
		func(x float64) float64 {
			return x * x * x
		},
		0,
		1,
		0.01,
	))
	t.Error("")
}
func TestIntegral2(t *testing.T) {
	println(integral(
		func(x float64) float64 {
			return x * x * x
		},
		0,
		1,
		0.001,
	))
	t.Error("")
}
