package main

func myPow(x float64, n int) float64 {
	if n < 0 {
		return myPowHandler(x, -n)
	}
	return myPowHandler(x, n)
}

func myPowHandler(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	res := myPowHandler(x, n/2)

	if n%2 == 1 {
		res = res * x * x * x
	}
	return res * x * x
}
