package main

import (
	"math"
	"testing"
)

func TestFixedPoint1(t *testing.T) {
	println(
		fixedPoint(
			math.Cos,
			1.0,
		),
	)
	t.Error("")
}
func TestFixedPoint2(t *testing.T) {
	println(
		fixedPoint(
			func(x float64) float64 { return math.Sin(x) + math.Cos(x) },
			1.0,
		),
	)
	t.Error("")
}

func TestAverageDamp(t *testing.T) {
	println(
		averageDamp(
			func(x float64) float64 { return x * x },
		)(10.0),
	)
	t.Error("")
}
func TestSqrt1(t *testing.T) {
	println(sqrt1(64.0))
	t.Error("")
}
func TestCubeRoot(t *testing.T) {
	println(cubeRoot(64.0))
	t.Error("")
}

func TestDeriv(t *testing.T) {
	println(deriv(
		func(x float64) float64 { return x * x * x },
	)(5.0),
	)
	t.Error("")
}

func TestSqrt2(t *testing.T) {
	println(sqrt2(64.0))
	t.Error("")
}
func TestSqrt3(t *testing.T) {
	println(sqrt3(64.0))
	t.Error("")
}
func TestSqrt4(t *testing.T) {
	println(sqrt4(64.0))
	t.Error("")
}
