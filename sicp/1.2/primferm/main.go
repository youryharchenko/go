package main

import (
	"flag"
	"math/rand"
	"strconv"
	"time"
)

func primeFermat(n, t int) bool {
	if t == 0 {
		return true
	}
	if fermatTest(n) {
		return primeFermat(n, t-1)
	}
	return false
}

func fermatTest(n int) bool {
	a := 1 + rand.Intn(n-1)
	return a == expmod(a, n, n)
}

func expmod(base, exp, m int) int {
	if exp == 0 {
		return 1
	}
	if exp%2 == 0 {
		v := expmod(base, exp/2, m)
		return (v * v) % m
	}
	v := expmod(base, exp-1, m)
	return (base * v) % m
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 2 {
		println("i want two int parameters")
		return
	}
	n, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		println(err, flag.Args()[0])
		return
	}
	t, err := strconv.Atoi(flag.Args()[1])
	if err != nil {
		println(err, flag.Args()[0])
		return
	}
	rand.Seed(time.Now().UnixNano())
	println(primeFermat(n, t))
}
