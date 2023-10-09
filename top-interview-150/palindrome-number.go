package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	keep := x
	reverse := 0
	for x > 0 {
		reverse = reverse*10 + x%10
		x /= 10
	}
	return reverse == keep
}
