package main

import (
	"flag"
	"strconv"
	"strings"
)

var coins = []int{1, 5, 10, 25, 50}

func countChange(amount int) int {
	return cc(amount, len(coins), 0)
}

func cc(amount, kindsOfCoins, deep int) int {
	ident := strings.Repeat(".", deep)
	println("cc:entry ", ident, amount, kindsOfCoins)
	if amount == 0 {
		println("cc:return", ident, 1)
		return 1
	} else if amount < 0 || kindsOfCoins == 0 {
		println("cc:return", ident, 0)
		return 0
	}
	r := cc(amount, kindsOfCoins-1, deep+1) + cc(amount-firstDenomi(kindsOfCoins), kindsOfCoins, deep+1)
	println("cc:return", ident, r)
	return r

}

func firstDenomi(kindsOfCoins int) int {
	return coins[kindsOfCoins-1]
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
	println(countChange(n))
}
