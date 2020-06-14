package main

import (
	"math"
	"testing"
)

func TestRoots1(t *testing.T) {
	println(
		halfIntervalMethod(
			math.Sin,
			2.0,
			4.0,
		),
	)
	t.Error("")
}

func TestRoots2(t *testing.T) {
	println(
		halfIntervalMethod(
			func(x float64) float64 { return x*x*x - 2*x - 3 },
			1.0,
			2.0,
		),
	)
	t.Error("")
}
