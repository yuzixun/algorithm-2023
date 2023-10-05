package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		m[t[i]]--
	}

	for _, count := range m {
		if count != 0 {
			return false
		}
	}
	return true
}
