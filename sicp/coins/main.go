package main

import (
	"flag"
	"strconv"
)

var coins = []int{1, 5, 10, 25, 50}

func countChange(amount int) int {
	return cc(amount, len(coins))
}

func cc(amount, kindsOfCoins int) int {
	//println("cc:entry", amount, kindsOfCoins)
	if amount == 0 {
		//println("cc:return", 1)
		return 1
	} else if amount < 0 || kindsOfCoins == 0 {
		//println("cc:return", 0)
		return 0
	} else {
		r := cc(amount, kindsOfCoins-1) + cc(amount-firstDenomi(kindsOfCoins), kindsOfCoins)
		//println("cc:return", r)
		return r
	}
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
