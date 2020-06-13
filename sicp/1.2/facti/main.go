package main

import (
	"flag"
	"strconv"
)

func iter(prod, count, n uint64) uint64 {
	if count > n {
		return prod
	}
	return iter(count*prod, count+1, n)
}

func facti(n uint64) uint64 {
	return iter(1, 1, n)
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
	println(facti(uint64(n)))
}
