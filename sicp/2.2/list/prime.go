package main

import (
	"flag"
	"strconv"
)

func prime(n int) bool {
	if smallestDivisor(n) == n {
		return true
	}
	return false
}

func smallestDivisor(n int) int {
	return findDivisor(n, 2)
}

func findDivisor(n, testDivisor int) int {
	if testDivisor*testDivisor > n {
		return n
	} else if n%testDivisor == 0 {
		return testDivisor
	} else {
		return findDivisor(n, testDivisor+1)
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		println("i want one int parameter")
		return
	}
	n, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		println(err, flag.Args()[0])
		return
	}
	println(prime(n))
}
