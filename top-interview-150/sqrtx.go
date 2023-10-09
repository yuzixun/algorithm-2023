package main

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	left, right := 1, x/2
	for left < right {
		mid := left + (right-left+1)/2
		switch {
		case mid*mid == x:
			return mid
		case mid*mid < x:
			left = mid
		case mid*mid > x:
			right = mid - 1
		}
	}

	return left
}
