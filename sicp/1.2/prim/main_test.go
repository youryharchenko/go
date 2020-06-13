package main

import "testing"

func Test(t *testing.T) {
	println(smallestDivisor(199))
	println(smallestDivisor(1999))
	println(smallestDivisor(19999))
	t.Error("")
}
