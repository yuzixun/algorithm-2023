package main

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isLetterOrNum(s[left]) {
			left++
		}
		for left < right && !isLetterOrNum(s[right]) {
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}

	return true
}

func isLetterOrNum(a byte) bool {
	return a >= 'a' && a <= 'z' || a >= '0' && a <= '9'
}
