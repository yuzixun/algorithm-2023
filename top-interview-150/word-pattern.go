package main

import "strings"

func wordPattern(pattern string, s string) bool {
	pattern2Words := make(map[byte]string)
	words2Patterns := make(map[string]byte)

	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		pItem, word := pattern[i], words[i]

		if (pattern2Words[pItem] != "" && pattern2Words[pItem] != word) || (words2Patterns[word] > 0 && words2Patterns[word] != pItem) {
			return false
		}

		pattern2Words[pItem] = word
		words2Patterns[word] = pItem
	}
	return true
}
