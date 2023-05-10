package leetcode_75

var (
	bMap = map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
)

func maxVowels(s string, k int) int {
	var (
		cur      int
		maxValue int
	)
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			cur++
		}
	}
	maxValue = cur

	for i := k; i < len(s); i++ {
		if isVowel(s[k-i]) {
			cur--
		}
		if isVowel(s[i]) {
			cur++
		}

		maxValue = max(maxValue, cur)
	}

	return maxValue
}

func isVowel(b byte) bool {
	return bMap[b]
}
