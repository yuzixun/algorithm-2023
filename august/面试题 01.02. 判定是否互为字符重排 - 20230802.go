package august

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[rune]int, 26)
	for _, e := range s1 {
		m[e]++
	}

	for _, e := range s2 {
		m[e]--
	}
	for _, e := range m {
		if e != 0 {
			return false
		}
	}

	return true
}
