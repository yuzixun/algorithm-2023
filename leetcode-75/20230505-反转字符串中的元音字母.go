package leetcode_75

import "strings"

func reverseVowels(s string) string {
	var (
		bs          = []byte(s)
		size        = len(s)
		left, right = 0, len(s) - 1
	)

	for left < right {
		for left < size && !strings.Contains("aeiouAEOIU", string(s[left])) {
			left++
		}
		for 0 < right && !strings.Contains("aeiouAEOIU", string(s[right])) {
			right--
		}
		if left < right {
			bs[left], bs[right] = bs[right], bs[left]
			left++
			right--
		}

	}

	return string(bs)
}
