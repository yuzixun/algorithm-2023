package august

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	origin := x
	reversed := 0
	for origin != 0 {
		tmp := origin % 10
		origin = origin / 10
		reversed = reversed*10 + tmp
	}

	return x == reversed
}
