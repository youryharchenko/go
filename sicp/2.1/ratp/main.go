package main

import "fmt"

type pairInt func(m int) int
type rat pairInt

func (r rat) String() string {
	return fmt.Sprintf("%d/%d", numer(r), denom(r))
}

func consInt(x, y int) pairInt {
	return func(m int) int {
		if m == 0 {
			return x
		}
		if m == 1 {
			return y
		}
		panic(fmt.Sprintf("Argument not 0 or 1 - CONS: %d", m))
	}
}

func carInt(p pairInt) int {
	return p(0)
}

func cdrInt(p pairInt) int {
	return p(1)
}

func makeRat(n, d int) rat {
	g := gcdrec(n, d)
	return rat(consInt(n/g, d/g))
}

func numer(r rat) int {
	return carInt(pairInt(r))
}

func denom(r rat) int {
	return cdrInt(pairInt(r))
}

func addRat(x, y rat) rat {
	return makeRat(numer(x)*denom(y)+numer(y)*denom(x), denom(x)*denom(y))
}

func subRat(x, y rat) rat {
	return makeRat(numer(x)*denom(y)-numer(y)*denom(x), denom(x)*denom(y))
}

func mulRat(x, y rat) rat {
	return makeRat(numer(x)*numer(y), denom(x)*denom(y))
}

func divRat(x, y rat) rat {
	return makeRat(numer(x)*denom(y), numer(y)*denom(x))
}

func gcdrec(a, b int) int {
	if b == 0 {
		return a
	}
	return gcdrec(b, a%b)
}
