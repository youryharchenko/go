package main

import (
	"flag"
	"strconv"
)

func iter(a, b, count uint64) uint64 {
	if count == 0 {
		return b
	}
	return iter(a+b, a, count-1)
}

func fibi(n uint64) uint64 {
	return iter(1, 0, n)
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
	println(fibi(uint64(n)))
}
