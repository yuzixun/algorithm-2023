package main

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var (
		left, right = 0, len(s) - 1
	)

	for left <= right {
		if !isValidByte(s[left]) {
			left++
			continue
		}
		if !isValidByte(s[right]) {
			right--
			continue
		}

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func isValidByte(a byte) bool {
	return (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9')
}
