package august

func canPermutePalindrome(s string) bool {

	m := make(map[rune]int)
	for _, b := range s {
		m[b]++
	}

	odd := 0
	for _, count := range m {
		if count%2 == 1 {
			odd++
		}
	}

	return odd <= 1
}
