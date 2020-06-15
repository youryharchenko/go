package main

import (
	"fmt"
	"testing"
)

var oneHalf = makeRat(1, 2)
var oneThird = makeRat(1, 3)

func TestOneHalf(t *testing.T) {
	fmt.Println(oneHalf)
	t.Error("")
}

func TestOneThird(t *testing.T) {
	fmt.Println(oneThird)
	t.Error("")
}

func TestAdd(t *testing.T) {
	fmt.Println(addRat(oneHalf, oneThird))
	t.Error("")
}

func TestMul1(t *testing.T) {
	fmt.Println(mulRat(oneHalf, oneThird))
	t.Error("")
}
func TestAdd2(t *testing.T) {
	fmt.Println(addRat(oneThird, oneThird))
	t.Error("")
}

func TestMul2(t *testing.T) {
	fmt.Println(mulRat(makeRat(-1, 2), makeRat(1, 2)))
	t.Error("")
}

func TestMul3(t *testing.T) {
	fmt.Println(mulRat(makeRat(1, 2), makeRat(-1, -2)))
	t.Error("")
}
