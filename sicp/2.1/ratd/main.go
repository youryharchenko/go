package main

import "fmt"

type pairInt [2]int
type rat pairInt

func (r rat) String() string {
	return fmt.Sprintf("%d/%d", numer(r), denom(r))
}

func consInt(x, y int) pairInt {
	return [2]int{x, y}
}

func carInt(p pairInt) int {
	return p[0]
}

func cdrInt(p pairInt) int {
	return p[1]
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
