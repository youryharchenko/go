package main

import (
	"flag"
	"strconv"
)

func factr(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	return n * factr(n-1)
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
	println(factr(uint64(n)))
}
