package main

import "math"

func halfIntervalMethod(f func(float64) float64, a, b float64) float64 {
	aValue := f(a)
	bValue := f(b)
	if aValue < 0.0 && bValue > 0.0 {
		return search(f, a, b)
	}
	if aValue > 0.0 && bValue < 0.0 {
		return search(f, b, a)
	}
	println("Values are not of opposite sign", aValue, bValue)
	return 0
}

func search(f func(float64) float64, negPoint, posPoint float64) float64 {
	midPoint := (negPoint + posPoint) / 2
	if closeEnough(negPoint, posPoint) {
		return midPoint
	}
	testValue := f(midPoint)
	if testValue > 0.0 {
		return search(f, negPoint, midPoint)
	}
	if testValue < 0.0 {
		return search(f, midPoint, posPoint)
	}
	return midPoint
}

func closeEnough(x, y float64) bool {
	return math.Abs(x-y) < 0.001
}
