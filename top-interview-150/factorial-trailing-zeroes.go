package main

func trailingZeroes(n int) int {
	count := 0
	for n > 0 {
		count += n / 5
		n /= 5
	}
	return count
}
