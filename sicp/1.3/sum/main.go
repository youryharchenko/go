package main

func sumInt(term func(int) int, a int, next func(int) int, b int) int {
	if a > b {
		return 0
	}
	return term(a) + sumInt(term, next(a), next, b)
}

func sumFloat(term func(float64) float64, a float64, next func(float64) float64, b float64) float64 {
	if a > b {
		return 0.0
	}
	return term(a) + sumFloat(term, next(a), next, b)
}

func integral(f func(float64) float64, a, b, dx float64) float64 {
	return sumFloat(
		f,
		a+dx/2.0,
		func(x float64) float64 {
			return x + dx
		},
		b,
	) * dx
}
