package main

func isPalindrome(x int) bool {
	if x <= 0 {
		return false
	}

	var (
		y    int
		temp = x
	)

	for temp > 0 {
		t := temp % 10
		y = y*10 + t
		temp = temp / 10
	}

	return x == y
}
