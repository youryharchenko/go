package main

import (
	"flag"
	"strconv"
)

func gcdrec(a, b int) int {
	if b == 0 {
		return a
	}
	return gcdrec(b, a%b)
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 2 {
		println("i want two int parameters")
		return
	}
	a, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		println(err, flag.Args()[0])
		return
	}
	b, err := strconv.Atoi(flag.Args()[1])
	if err != nil {
		println(err, flag.Args()[1])
		return
	}
	println(gcdrec(a, b))
}
