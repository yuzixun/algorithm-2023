package main

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}

	min, max := 0, x
	for max-min > 1 {
		m := (max + min) / 2
		if x/m < m {
			max = m
		} else {
			min = m
		}
	}
	return min
}
