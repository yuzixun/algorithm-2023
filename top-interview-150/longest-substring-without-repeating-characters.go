package main

import "github.com/algorithm-2023/tools"

func lengthOfLongestSubstring(s string) int {
	wordCount := make(map[byte]int)

	result := 0

	right := -1

	for i := 0; i < len(s); i++ {
		if i > 0 {
			delete(wordCount, s[i-1])
		}

		for (right+1) < len(s) && wordCount[s[right+1]] == 0 {
			wordCount[s[right+1]] += 1
			right++
		}

		result = tools.Max(result, right-i+1)
	}

	return result
}
